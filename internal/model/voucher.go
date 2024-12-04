package model

import (
	"gorm.io/gorm"
)

type Voucher struct {
	gorm.Model
	VoucherName            string `json:"voucher_name" gorm:"column:voucher_name"`
	Describe               string `json:"describe" gorm:"column:describe;"`
	Discount               int    `json:"discount" gorm:"column:discount;"`
	ExpiredTimeAfterCreate int    `json:"expired_time_after_create" gorm:"column:expired_time_after_create;"`
}

func (Voucher) TableName() string {
	return "voucher"
}
