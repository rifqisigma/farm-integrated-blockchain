package usecase

import (
	"farm-integrated-web3/dto"
	"farm-integrated-web3/internal/repository"
)

type RetailerUsecase interface {
	AddRetailerCart(input *dto.CreateRetailerCartRequest) error
	UpdateRetailerCart(input *dto.UpdateRetailerCartRequest) error
	DeleteRetailerCart(retailerCartId, retailerId uint) error
	SearchRetailerCart(search string) ([]dto.GetRetailerCart, error)
	GetRetailerCarts(retailerProfileId uint) ([]dto.GetRetailerCart, error)
	GetRetailerCartById(retailerCartId uint) (*dto.GetRetailerCartById, error)
}

type retailerUsecase struct {
	retailerRepo repository.RetailerRepository
}

func NewRetailerUsecase(retailerRepo repository.RetailerRepository) RetailerUsecase {
	return &retailerUsecase{retailerRepo}
}

func (u *retailerUsecase) AddRetailerCart(input *dto.CreateRetailerCartRequest) error {
	return u.retailerRepo.AddRetailerCart(input)
}

func (u *retailerUsecase) UpdateRetailerCart(input *dto.UpdateRetailerCartRequest) error {
	return u.retailerRepo.UpdateRetailerCart(input)
}

func (u *retailerUsecase) DeleteRetailerCart(retailerCartId, retailerId uint) error {
	return u.retailerRepo.DeleteRetailerCart(retailerCartId, retailerId)
}
func (u *retailerUsecase) SearchRetailerCart(search string) ([]dto.GetRetailerCart, error) {
	return u.retailerRepo.SearchRetailerCart(search)
}

func (u *retailerUsecase) GetRetailerCarts(retailerProfileId uint) ([]dto.GetRetailerCart, error) {
	return u.retailerRepo.GetRetailerCarts(retailerProfileId)
}
func (u *retailerUsecase) GetRetailerCartById(retailerCartId uint) (*dto.GetRetailerCartById, error) {
	return u.retailerRepo.GetRetailerCartById(retailerCartId)
}
