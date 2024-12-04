package repositories

import (
	"time"
	"trinity/internal/model"

	"gorm.io/gorm"
)

type VoucherUserRepository interface {
	CreateVoucherUser(voucherId, userId uint, discount int, expiredTime time.Time, isUsed bool) error
	GetVoucherByVoucherName(voucherName string) (*model.Voucher, error)
}

type voucherUserRepository struct {
	db *gorm.DB
}

func NewVoucherUserRepository(db *gorm.DB) VoucherUserRepository {
	return &voucherUserRepository{db}
}

func (repo *voucherUserRepository) CreateVoucherUser(voucherId, userId uint, discount int, expiredTime time.Time, isUsed bool) error {
	voucherUser := model.VoucherUser{
		VoucherID:   voucherId,
		UserID:      userId,
		Discount:    discount,
		ExpiredTime: expiredTime,
		IsUsed:      false,
	}

	if err := repo.db.Create(&voucherUser).Error; err != nil {
		return err
	}

	return nil
}

func (repo *voucherUserRepository) GetVoucherByVoucherName(voucherName string) (*model.Voucher, error) {
	voucher := &model.Voucher{}
	if err := repo.db.First(voucher, "voucher_name = ?", voucherName).Error; err != nil {
		return nil, err
	}
	return voucher, nil
}

func (repo *voucherUserRepository) GetAllVouchers() ([]model.Voucher, error) {
	var vouchers []model.Voucher
	if err := repo.db.Find(&vouchers).Error; err != nil {
		return nil, err
	}
	return vouchers, nil
}
