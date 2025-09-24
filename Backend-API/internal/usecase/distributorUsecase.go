package usecase

import (
	"context"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/internal/repository"
)

type DistributorUsecase interface {
	CreateDistribution(ctx context.Context, input *dto.CreateDistributionRequest) error
	UpdateDistribution(ctx context.Context, input *dto.UpdateDistributionRequest) error
	DeleteDistribution(ctx context.Context, distrbutionId, distributorId uint) error
	UpdateStatusOfDistribution(ctx context.Context, input *dto.UpdateStatusDistributionRequest) error
	ApprovedRetailerCartForRetailer(ctx context.Context, input *dto.ApprovedRetailerCart) error
	GetDistributionFYP(ctx context.Context) ([]dto.GetDistribution, error)
	//get
	SearchDistributions(ctx context.Context, search string) ([]dto.GetDistribution, error)
	GetDistributionsByDistributorId(ctx context.Context, id uint) ([]dto.GetDistribution, error)
	GetDistributionByid(ctx context.Context, id uint) (*dto.GetDistributionById, error)
}

type distributorUsecase struct {
	distributorRepo repository.DistributorRepository
}

func NewDistributorUsecase(distributorRepo repository.DistributorRepository) DistributorUsecase {
	return &distributorUsecase{distributorRepo}
}

func (u *distributorUsecase) CreateDistribution(ctx context.Context, input *dto.CreateDistributionRequest) error {
	return u.distributorRepo.CreateDistribution(ctx, input)
}

func (u *distributorUsecase) UpdateDistribution(ctx context.Context, input *dto.UpdateDistributionRequest) error {
	return u.distributorRepo.UpdateDistribution(ctx, input)
}

func (u *distributorUsecase) DeleteDistribution(ctx context.Context, distrbutionId, distributorId uint) error {
	return u.distributorRepo.DeleteDistribution(ctx, distrbutionId, distributorId)
}

func (u *distributorUsecase) SearchDistributions(ctx context.Context, search string) ([]dto.GetDistribution, error) {
	return u.distributorRepo.SearchDistributions(ctx, search)
}

func (u *distributorUsecase) GetDistributionsByDistributorId(ctx context.Context, id uint) ([]dto.GetDistribution, error) {
	return u.distributorRepo.GetDistributionsByDistributorId(ctx, id)
}

func (u *distributorUsecase) GetDistributionByid(ctx context.Context, id uint) (*dto.GetDistributionById, error) {
	return u.distributorRepo.GetDistributionByid(ctx, id)
}

func (u *distributorUsecase) UpdateStatusOfDistribution(ctx context.Context, input *dto.UpdateStatusDistributionRequest) error {
	return u.distributorRepo.UpdateStatusOfDistribution(ctx, input)
}

func (u *distributorUsecase) ApprovedRetailerCartForRetailer(ctx context.Context, input *dto.ApprovedRetailerCart) error {
	return u.distributorRepo.ApprovedRetailerCartForRetailer(ctx, input)
}

func (u *distributorUsecase) GetDistributionFYP(ctx context.Context) ([]dto.GetDistribution, error) {
	return u.distributorRepo.GetDistributionFYP(ctx)
}
