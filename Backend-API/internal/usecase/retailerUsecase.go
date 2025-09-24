package usecase

import (
	"context"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/internal/repository"
)

type RetailerUsecase interface {
	AddRetailerCart(ctx context.Context, input *dto.CreateRetailerCartRequest) error
	UpdateRetailerCart(ctx context.Context, input *dto.UpdateRetailerCartRequest) error
	DeleteRetailerCart(ctx context.Context, retailerCartId, retailerId uint) error
	SearchRetailerCart(ctx context.Context, search string) ([]dto.GetRetailerCart, error)
	GetRetailerCarts(ctx context.Context, retailerProfileId uint) ([]dto.GetRetailerCart, error)
	GetRetailerCartById(ctx context.Context, retailerCartId uint) (*dto.GetRetailerCartById, error)
	GetRetailerCartsFYP(ctx context.Context) ([]dto.GetRetailerCart, error)
}

type retailerUsecase struct {
	retailerRepo repository.RetailerRepository
}

func NewRetailerUsecase(retailerRepo repository.RetailerRepository) RetailerUsecase {
	return &retailerUsecase{retailerRepo}
}

func (u *retailerUsecase) AddRetailerCart(ctx context.Context, input *dto.CreateRetailerCartRequest) error {
	return u.retailerRepo.AddRetailerCart(ctx, input)
}

func (u *retailerUsecase) UpdateRetailerCart(ctx context.Context, input *dto.UpdateRetailerCartRequest) error {
	return u.retailerRepo.UpdateRetailerCart(ctx, input)
}

func (u *retailerUsecase) DeleteRetailerCart(ctx context.Context, retailerCartId, retailerId uint) error {
	return u.retailerRepo.DeleteRetailerCart(ctx, retailerCartId, retailerId)
}
func (u *retailerUsecase) SearchRetailerCart(ctx context.Context, search string) ([]dto.GetRetailerCart, error) {
	return u.retailerRepo.SearchRetailerCart(ctx, search)
}

func (u *retailerUsecase) GetRetailerCarts(ctx context.Context, retailerProfileId uint) ([]dto.GetRetailerCart, error) {
	return u.retailerRepo.GetRetailerCarts(ctx, retailerProfileId)
}
func (u *retailerUsecase) GetRetailerCartById(ctx context.Context, retailerCartId uint) (*dto.GetRetailerCartById, error) {
	return u.retailerRepo.GetRetailerCartById(ctx, retailerCartId)
}

func (u *retailerUsecase) GetRetailerCartsFYP(ctx context.Context) ([]dto.GetRetailerCart, error) {
	return u.retailerRepo.GetRetailerCartsFYP(ctx)
}
