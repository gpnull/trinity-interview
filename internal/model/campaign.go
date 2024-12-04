package model

import (
	"time"

	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model
	CampaignName string    `json:"campaign_name" gorm:"column:campaign_name"`
	Link         string    `json:"link" gorm:"column:link;unique"`
	Describe     string    `json:"describe" gorm:"column:describe;"`
	Limit        int       `json:"limit" gorm:"column:limit;"`
	ExpiredTime  time.Time `json:"expired_time" gorm:"column:expiredTime;"`

	VoucherName string `json:"voucher_name" gorm:"column:voucher_name"`
}

func (Campaign) TableName() string {
	return "campaigns"
}
