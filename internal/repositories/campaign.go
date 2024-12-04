package repositories

import (
	"time"
	"trinity/internal/model"

	"gorm.io/gorm"
)

type CampaignRepository interface {
	CreateCampaign(campaignName, link, describe string, limit int, expiredTime time.Time, voucherName string) error
	UpdateCampaign(link, describe string, limit int, expiredTime time.Time, voucherName string) error
	GetCampaignByLink(link string) (*model.Campaign, error)
	GetCampaignByCampaignName(campaignName string) (*model.Campaign, error)
	UpdateLimitCampaign(campaignName string, newLimit int) error
}

type campaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) CampaignRepository {
	return &campaignRepository{db}
}

func (repo *campaignRepository) CreateCampaign(campaignName, link, describe string, limit int, expiredTime time.Time, voucherName string) error {
	campaign := model.Campaign{
		CampaignName: campaignName,
		Link:         link,
		Describe:     describe,
		Limit:        limit,
		ExpiredTime:  expiredTime,

		VoucherName: voucherName,
	}

	if err := repo.db.Create(&campaign).Error; err != nil {
		return err
	}

	return nil
}

func (repo *campaignRepository) UpdateCampaign(link, describe string, limit int, expiredTime time.Time, voucherName string) error {
	campaign := model.Campaign{}
	if err := repo.db.First(&campaign, link).Error; err != nil {
		return err
	}

	campaign.Link = link
	campaign.Describe = describe
	campaign.Limit = limit
	campaign.ExpiredTime = expiredTime
	campaign.VoucherName = voucherName

	if err := repo.db.Save(&campaign).Error; err != nil {
		return err
	}

	return nil
}

func (repo *campaignRepository) GetCampaignByLink(link string) (*model.Campaign, error) {
	campaign := &model.Campaign{}
	if err := repo.db.First(campaign, "link = ?", link).Error; err != nil {
		return nil, err
	}
	return campaign, nil
}

func (repo *campaignRepository) GetCampaignByCampaignName(campaignName string) (*model.Campaign, error) {
	campaign := &model.Campaign{}
	if err := repo.db.First(campaign, "campaign_name = ?", campaignName).Error; err != nil {
		return nil, err
	}
	return campaign, nil
}

func (repo *campaignRepository) UpdateLimitCampaign(campaignName string, newLimit int) error {
	campaign := model.Campaign{}
	if err := repo.db.First(&campaign, "campaign_name = ?", campaignName).Error; err != nil {
		return err
	}

	campaign.Limit = newLimit

	if err := repo.db.Save(&campaign).Error; err != nil {
		return err
	}

	return nil
}
