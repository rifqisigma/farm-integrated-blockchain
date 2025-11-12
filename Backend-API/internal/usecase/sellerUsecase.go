package usecase

import (
	"context"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/internal/repository"

	"github.com/redis/go-redis/v9"
)

type SellerUsecase interface {
	AddSellerBox(ctx context.Context, input *dto.CreateSellerBox) error
	UpdateSellerBox(ctx context.Context, input *dto.UpdateSellerBox) error
	DeleteSellerBox(ctx context.Context, SellerBoxId, retailerId uint) error

	//get
	SearchSellerBox(ctx context.Context, search string) ([]dto.GetSellerBox, error)
	ListGetSellerBoxsbySellerId(ctx context.Context, SellerProfileId uint) ([]dto.GetSellerBox, error)
	GetSellerBoxById(ctx context.Context, SellerBoxId uint) (*dto.GetSellerBoxById, error)
	ListGetSellerBoxsbySellerIdFYP(ctx context.Context) ([]dto.GetSellerBox, error)
}

type sellerUsecase struct {
	retailerRepo repository.SellerRepository
	redis        *redis.Client
}

func NewSellerUsecase(retailerRepo repository.SellerRepository, redis *redis.Client) SellerUsecase {
	return &sellerUsecase{retailerRepo, redis}
}

func (u *sellerUsecase) AddSellerBox(ctx context.Context, input *dto.CreateSellerBox) error {
	return u.retailerRepo.AddSellerBox(ctx, input)
}

func (u *sellerUsecase) UpdateSellerBox(ctx context.Context, input *dto.UpdateSellerBox) error {
	return u.retailerRepo.UpdateSellerBox(ctx, input)
}

func (u *sellerUsecase) DeleteSellerBox(ctx context.Context, SellerBoxId, retailerId uint) error {
	return u.retailerRepo.DeleteSellerBox(ctx, SellerBoxId, retailerId)
}
func (u *sellerUsecase) SearchSellerBox(ctx context.Context, search string) ([]dto.GetSellerBox, error) {
	return u.retailerRepo.SearchSellerBox(ctx, search)
}

func (u *sellerUsecase) ListGetSellerBoxsbySellerId(ctx context.Context, SellerProfileId uint) ([]dto.GetSellerBox, error) {
	return u.retailerRepo.ListGetSellerBoxsbySellerId(ctx, SellerProfileId)
}
func (u *sellerUsecase) GetSellerBoxById(ctx context.Context, SellerBoxId uint) (*dto.GetSellerBoxById, error) {
	return u.retailerRepo.GetSellerBoxById(ctx, SellerBoxId)
}

func (u *sellerUsecase) ListGetSellerBoxsbySellerIdFYP(ctx context.Context) ([]dto.GetSellerBox, error) {
	return u.retailerRepo.ListGetSellerBoxsbySellerIdFYP(ctx)
}
