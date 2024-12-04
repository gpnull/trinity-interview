package service

import (
	"strings"
	"time"
	"trinity/internal/repositories"
)

type CreateCampaignRequest struct {
	Link        string    `json:"link" gorm:"column:link;unique"`
	Describe    string    `json:"describe" gorm:"column:describe;"`
	Limit       int       `json:"limit" gorm:"column:limit;"`
	ExpiredTime time.Time `json:"expired_time" gorm:"column:expired_time;"`

	VoucherName string `json:"voucher_name" gorm:"column:voucher_name;"`
}

type AA struct {
	Link     string `json:"link" gorm:"column:link;unique"`
	Describe string `json:"describe" gorm:"column:describe;"`
	Limit    int    `json:"limit" gorm:"column:limit;"`

	VoucherName string `json:"voucher_name" gorm:"column:voucher_name;"`
}

type GetCampaignRequest struct {
	Link string `json:"link" gorm:"column:link"`
}

type UpdateCampaignRequest struct {
	Link     string ` gorm:"column:link"`
	NewLimit int
}

type GetCampaignResponse struct {
	Link        string    `json:"link" gorm:"column:link"`
	Describe    string    `json:"describe" gorm:"column:describe;"`
	Limit       int       `json:"limit" gorm:"column:limit;"`
	ExpiredTime time.Time `json:"expired_time" gorm:"column:expiredTime;"`
}

type CampaignService interface {
	CreateCampaign(request CreateCampaignRequest) (GetCampaignResponse, error)
	GetCampaignByLink(request GetCampaignRequest) (GetCampaignResponse, error)
	UpdateLimitCampaign(request UpdateCampaignRequest) error
}

type campaignService struct {
	campaignRepo repositories.CampaignRepository
	voucherRepo  repositories.VoucherRepository
}

func NewCampaignService(campaignRepo repositories.CampaignRepository, voucherRepo repositories.VoucherRepository) CampaignService {
	return &campaignService{campaignRepo, voucherRepo}
}

func (svc *campaignService) CreateCampaign(request CreateCampaignRequest) (GetCampaignResponse, error) {
	voucher, err := svc.voucherRepo.GetVoucherByVoucherName(request.VoucherName)
	if err != nil {
		return GetCampaignResponse{}, err
	}

	parts := strings.Split(request.Link, "/")
	campaignName := parts[len(parts)-1]

	err = svc.campaignRepo.CreateCampaign(campaignName, request.Link, request.Describe, request.Limit, request.ExpiredTime, voucher.VoucherName)
	if err != nil {
		return GetCampaignResponse{}, err
	}

	campaign, err := svc.GetCampaignByLink(GetCampaignRequest{Link: request.Link})
	if err != nil {
		return GetCampaignResponse{}, err
	}

	return campaign, nil
}

func (svc *campaignService) GetCampaignByLink(request GetCampaignRequest) (GetCampaignResponse, error) {
	campaign, err := svc.campaignRepo.GetCampaignByLink(request.Link)
	if err != nil {
		return GetCampaignResponse{}, err
	}

	return GetCampaignResponse{
		Link:        campaign.Link,
		Describe:    campaign.Describe,
		Limit:       campaign.Limit,
		ExpiredTime: campaign.ExpiredTime}, nil
}

func (svc *campaignService) UpdateLimitCampaign(request UpdateCampaignRequest) error {
	err := svc.campaignRepo.UpdateLimitCampaign(request.Link, request.NewLimit)
	if err != nil {
		return err
	}

	return nil
}
