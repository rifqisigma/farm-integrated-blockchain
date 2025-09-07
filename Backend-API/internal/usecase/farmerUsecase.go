package usecase

import (
	"farm-integrated-web3/dto"
	"farm-integrated-web3/internal/repository"
)

type FarmerUsecase interface {
	CreateHarvest(input *dto.HarvestRequest) error
	UpdateHarvest(input *dto.HarvestUpdate) error
	DeleteHarvest(farmerProfileId, harvestId uint) error
	AcceptedFarmerForDistributor(input *dto.AcceptFarmerForDistributor) error

	//get
	ListHarvestByFarmerId(farmerId uint) ([]dto.GetListHarvest, error)
	HarvestById(harvestId uint) (*dto.GetHarvestById, error)
	SearchHarvest(search string) ([]dto.GetListHarvest, error)
}

type farmerUsecase struct {
	farmerRepo repository.FarmerRepository
}

func NewFarmerUsecase(farmerRepo repository.FarmerRepository) FarmerUsecase {
	return &farmerUsecase{farmerRepo}
}

func (u *farmerUsecase) CreateHarvest(input *dto.HarvestRequest) error {
	return u.farmerRepo.CreateHarvest(input)
}

func (u *farmerUsecase) UpdateHarvest(input *dto.HarvestUpdate) error {
	return u.farmerRepo.UpdateHarvest(input)
}

func (u *farmerUsecase) DeleteHarvest(farmerProfileId, harvestId uint) error {
	return u.farmerRepo.DeleteHarvest(farmerProfileId, harvestId)
}

func (u *farmerUsecase) AcceptedFarmerForDistributor(input *dto.AcceptFarmerForDistributor) error {
	return u.farmerRepo.AcceptedFarmerForDistributor(input)
}

func (u *farmerUsecase) ListHarvestByFarmerId(farmerId uint) ([]dto.GetListHarvest, error) {
	return u.farmerRepo.ListHarvestByFarmerId(farmerId)
}

func (u *farmerUsecase) HarvestById(harvestId uint) (*dto.GetHarvestById, error) {
	return u.farmerRepo.HarvestById(harvestId)
}

func (u *farmerUsecase) SearchHarvest(search string) ([]dto.GetListHarvest, error) {
	return u.farmerRepo.SearchHarvest(search)

}
