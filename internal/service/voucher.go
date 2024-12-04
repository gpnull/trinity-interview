package service

import (
	"trinity/internal/repositories"
)

type CreateVoucherRequest struct {
	VoucherName            string `json:"voucher_name" gorm:"column:voucher_name"`
	Describe               string `json:"describe" gorm:"column:describe;"`
	Discount               int    `json:"discount" gorm:"column:discount;"`
	ExpiredTimeAfterCreate int    `json:"expired_time_after_create" gorm:"column:expired_time_after_create;"`
}

type GetVoucherRequest struct {
	VoucherName string `json:"voucher_name" gorm:"column:voucher_name"`
}

type GetVoucherResponse struct {
	VoucherName            string `json:"voucher_name" gorm:"column:voucher_name"`
	Describe               string `json:"describe" gorm:"column:describe;"`
	Discount               int    `json:"discount" gorm:"column:discount;"`
	ExpiredTimeAfterCreate int    `json:"expired_time_after_create" gorm:"column:expired_time_after_create;"`
}

type VoucherService interface {
	CreateVoucher(request CreateVoucherRequest) (GetVoucherResponse, error)
	GetVoucherByVoucherName(request GetVoucherRequest) (GetVoucherResponse, error)
}

type voucherService struct {
	voucherRepo repositories.VoucherRepository
}

func NewVoucherService(voucherRepo repositories.VoucherRepository) VoucherService {
	return &voucherService{voucherRepo}
}

func (svc *voucherService) CreateVoucher(request CreateVoucherRequest) (GetVoucherResponse, error) {

	err := svc.voucherRepo.CreateVoucher(request.VoucherName, request.Describe, request.Discount, request.ExpiredTimeAfterCreate)
	if err != nil {
		return GetVoucherResponse{}, err
	}

	voucher, err := svc.GetVoucherByVoucherName(GetVoucherRequest{VoucherName: request.VoucherName})
	if err != nil {
		return GetVoucherResponse{}, err
	}

	return voucher, nil
}

func (svc *voucherService) GetVoucherByVoucherName(request GetVoucherRequest) (GetVoucherResponse, error) {
	voucher, err := svc.voucherRepo.GetVoucherByVoucherName(request.VoucherName)
	if err != nil {
		return GetVoucherResponse{}, err
	}

	return GetVoucherResponse{
		VoucherName:            voucher.VoucherName,
		Describe:               voucher.Describe,
		Discount:               voucher.Discount,
		ExpiredTimeAfterCreate: voucher.ExpiredTimeAfterCreate}, nil
}
