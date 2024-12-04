package service

import (
	"time"
	errorType "trinity/errors"
	"trinity/internal/repositories"
	"trinity/utils"
)

type CreateUserRequest struct {
	Email    string `json:"email" gorm:"column:email;"`
	Name     string `json:"name" gorm:"column:name;"`
	Password string `json:"password" gorm:"column:password;"`
	UserType string `json:"user_type" gorm:"column:user_type;"`

	CampaignName string `json:"campaign_name"`
}

type GetUserRequest struct {
	Email string `json:"email" gorm:"column:email"`
}

type GetUserResponse struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	UserType string `json:"user_type"`
}

type UserService interface {
	CreateUser(request CreateUserRequest) (GetUserResponse, error)
	GetUser(request GetUserRequest) (GetUserResponse, error)
}

type userService struct {
	userRepo        repositories.UserRepository
	campaignRepo    repositories.CampaignRepository
	voucherRepo     repositories.VoucherRepository
	voucherUserRepo repositories.VoucherUserRepository
}

func NewUserService(userRepo repositories.UserRepository, campaignRepo repositories.CampaignRepository,
	voucherRepo repositories.VoucherRepository, voucherUserRepo repositories.VoucherUserRepository) UserService {
	return &userService{userRepo, campaignRepo, voucherRepo, voucherUserRepo}
}

func (svc *userService) CreateUser(request CreateUserRequest) (GetUserResponse, error) {
	passwordHashed, err := utils.HashPassword(request.Password)
	if err != nil {
		return GetUserResponse{}, err
	}

	err = svc.userRepo.CreateUser(request.Email, request.Name, request.UserType, passwordHashed)
	if err != nil {
		return GetUserResponse{}, err
	}

	user, err := svc.GetUser(GetUserRequest{Email: request.Email})
	if err != nil {
		return GetUserResponse{}, err
	}

	if request.CampaignName != "" {
		campaign, err := svc.campaignRepo.GetCampaignByCampaignName(request.CampaignName)
		if err != nil {
			return user, errorType.ErrInvalidCampaign
		}

		if campaign.ExpiredTime.Before(time.Now()) {
			return user, errorType.ErrInvalidCampaign
		}

		voucher, err := svc.voucherRepo.GetVoucherByVoucherName(campaign.VoucherName)
		if err != nil {
			return user, errorType.ErrInvalidCampaign
		}

		expiredTime := time.Now().AddDate(0, 0, voucher.ExpiredTimeAfterCreate)
		err = svc.voucherUserRepo.CreateVoucherUser(voucher.ID, user.Id, voucher.Discount, expiredTime, false)
		if err != nil {
			return user, errorType.ErrInvalidCampaign
		}
	}

	return user, nil
}

func (svc *userService) GetUser(request GetUserRequest) (GetUserResponse, error) {
	user, err := svc.userRepo.GetUser(request.Email)
	if err != nil {
		return GetUserResponse{}, err
	}

	return GetUserResponse{
		Id:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		UserType: user.UserType}, nil
}
