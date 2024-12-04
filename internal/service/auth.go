package service

import (
	"trinity/internal/repositories"
	"trinity/utils"
)

type AuthService interface {
	Login(request LoginRequest) (string, error)
}
type authService struct {
	userRepo repositories.UserRepository
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo}
}

func (svc *authService) Login(request LoginRequest) (string, error) {
	user, err := svc.userRepo.GetUser(request.Email)
	if err != nil {
		return "email or password is wrong", err
	}

	isCorrectPassword := utils.CheckPasswordHash(request.Password, user.Password)
	if !isCorrectPassword {
		return "email or password is wrong", err
	}

	accessToken, err := utils.GenerateJWTToken(user.Email)
	if err != nil {
		return "error access token", err
	}

	return accessToken, nil
}
