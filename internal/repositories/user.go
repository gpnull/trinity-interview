package repositories

import (
	"trinity/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(email, name, userType, password string) error
	GetUser(email string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (repo *userRepository) CreateUser(email, name, userType, password string) error {
	user := model.User{
		Email:    email,
		Name:     name,
		UserType: userType,
		Password: password,
	}

	if err := repo.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) GetUser(email string) (*model.User, error) {
	var user model.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
