package repositories

import (
	"trinity/internal/model"

	"gorm.io/gorm"
)

type VoucherRepository interface {
	CreateVoucher(voucherName, describe string, discount, expiredTimeAfterCreate int) error
	UpdateVoucher(voucherName, describe string, discount, expiredTimeAfterCreate int) error
	GetVoucherByVoucherName(voucherName string) (*model.Voucher, error)
	GetAllVouchers() ([]model.Voucher, error)
}

type voucherRepository struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) VoucherRepository {
	return &voucherRepository{db}
}

func (repo *voucherRepository) CreateVoucher(voucherName, describe string, discount, expiredTimeAfterCreate int) error {
	voucher := model.Voucher{
		VoucherName:            voucherName,
		Describe:               describe,
		Discount:               discount,
		ExpiredTimeAfterCreate: expiredTimeAfterCreate,
	}

	if err := repo.db.Create(&voucher).Error; err != nil {
		return err
	}

	return nil
}

func (repo *voucherRepository) UpdateVoucher(voucherName, describe string, discount, expiredTimeAfterCreate int) error {
	voucher := model.Voucher{}
	if err := repo.db.First(&voucher, voucherName).Error; err != nil {
		return err
	}

	voucher.VoucherName = voucherName
	voucher.Describe = describe
	voucher.Discount = discount
	voucher.ExpiredTimeAfterCreate = expiredTimeAfterCreate

	if err := repo.db.Save(&voucher).Error; err != nil {
		return err
	}

	return nil
}

func (repo *voucherRepository) GetVoucherByVoucherName(voucherName string) (*model.Voucher, error) {
	voucher := &model.Voucher{}
	if err := repo.db.First(voucher, "voucher_name = ?", voucherName).Error; err != nil {
		return nil, err
	}
	return voucher, nil
}

func (repo *voucherRepository) GetAllVouchers() ([]model.Voucher, error) {
	var vouchers []model.Voucher
	if err := repo.db.Find(&vouchers).Error; err != nil {
		return nil, err
	}
	return vouchers, nil
}
