package model

import (
	"time"

	"gorm.io/gorm"
)

type VoucherUser struct {
	gorm.Model
	VoucherID   uint      `json:"voucher_id" gorm:"column:voucher_id"`
	UserID      uint      `json:"user_id" gorm:"column:user_id"`
	Discount    int       `json:"discount" gorm:"column:discount;"`
	ExpiredTime time.Time `json:"expired_time" gorm:"column:expired_time;"`
	IsUsed      bool      `json:"is_used" gorm:"column:is_used;"`
}

func (VoucherUser) TableName() string {
	return "voucher_user"
}
