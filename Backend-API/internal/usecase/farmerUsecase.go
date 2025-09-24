package usecase

import (
	"context"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/internal/repository"
)

type FarmerUsecase interface {
	CreateHarvest(ctx context.Context, input *dto.HarvestRequest) error
	UpdateHarvest(ctx context.Context, input *dto.HarvestUpdate) error
	DeleteHarvest(ctx context.Context, farmerProfileId, harvestId uint) error
	AcceptedFarmerForDistributor(ctx context.Context, input *dto.AcceptFarmerForDistributor) error
	ListHarvestFYP(ctx context.Context) ([]dto.GetListHarvest, error)

	//get
	ListHarvestByFarmerId(ctx context.Context, farmerId uint) ([]dto.GetListHarvest, error)
	HarvestById(ctx context.Context, harvestId uint) (*dto.GetHarvestById, error)
	SearchHarvest(ctx context.Context, search string) ([]dto.GetListHarvest, error)
}

type farmerUsecase struct {
	farmerRepo repository.FarmerRepository
}

func NewFarmerUsecase(farmerRepo repository.FarmerRepository) FarmerUsecase {
	return &farmerUsecase{farmerRepo}
}

func (u *farmerUsecase) CreateHarvest(ctx context.Context, input *dto.HarvestRequest) error {
	return u.farmerRepo.CreateHarvest(ctx, input)
}

func (u *farmerUsecase) UpdateHarvest(ctx context.Context, input *dto.HarvestUpdate) error {
	return u.farmerRepo.UpdateHarvest(ctx, input)
}

func (u *farmerUsecase) DeleteHarvest(ctx context.Context, farmerProfileId, harvestId uint) error {
	return u.farmerRepo.DeleteHarvest(ctx, farmerProfileId, harvestId)
}

func (u *farmerUsecase) AcceptedFarmerForDistributor(ctx context.Context, input *dto.AcceptFarmerForDistributor) error {
	return u.farmerRepo.AcceptedFarmerForDistributor(ctx, input)
}

func (u *farmerUsecase) ListHarvestByFarmerId(ctx context.Context, farmerId uint) ([]dto.GetListHarvest, error) {
	return u.farmerRepo.ListHarvestByFarmerId(ctx, farmerId)
}

func (u *farmerUsecase) HarvestById(ctx context.Context, harvestId uint) (*dto.GetHarvestById, error) {
	return u.farmerRepo.HarvestById(ctx, harvestId)
}

func (u *farmerUsecase) SearchHarvest(ctx context.Context, search string) ([]dto.GetListHarvest, error) {
	return u.farmerRepo.SearchHarvest(ctx, search)

}

func (u *farmerUsecase) ListHarvestFYP(ctx context.Context) ([]dto.GetListHarvest, error) {
	return u.farmerRepo.ListHarvestFYP(ctx)
}
