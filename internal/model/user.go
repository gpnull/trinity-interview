package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"column:email;unique"`
	Name     string `json:"name" gorm:"column:name;"`
	UserType string `json:"user_type" gorm:"column:user_type;"`
	Password string `json:"password" gorm:"column:password;"`
}

func (User) TableName() string {
	return "users"
}
