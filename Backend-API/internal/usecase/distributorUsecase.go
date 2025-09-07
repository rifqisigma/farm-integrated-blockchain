package usecase

import (
	"farm-integrated-web3/dto"
	"farm-integrated-web3/internal/repository"
)

type DistributorUsecase interface {
	CreateDistribution(input *dto.CreateDistributionRequest) error
	UpdateDistribution(input *dto.UpdateDistributionRequest) error
	DeleteDistribution(distrbutionId, distributorId uint) error
	UpdateStatusOfDistribution(input *dto.UpdateStatusDistributionRequest) error
	ApprovedRetailerCartForRetailer(input *dto.ApprovedRetailerCart) error
	//get
	SearchDistributions(search string) ([]dto.GetDistribution, error)
	GetDistributionsByDistributorId(id uint) ([]dto.GetDistribution, error)
	GetDistributionByid(id uint) (*dto.GetDistributionById, error)
}

type distributorUsecase struct {
	distributorRepo repository.DistributorRepository
}

func NewDistributorUsecase(distributorRepo repository.DistributorRepository) DistributorUsecase {
	return &distributorUsecase{distributorRepo}
}

func (u *distributorUsecase) CreateDistribution(input *dto.CreateDistributionRequest) error {
	return u.distributorRepo.CreateDistribution(input)
}

func (u *distributorUsecase) UpdateDistribution(input *dto.UpdateDistributionRequest) error {
	return u.distributorRepo.UpdateDistribution(input)
}

func (u *distributorUsecase) DeleteDistribution(distrbutionId, distributorId uint) error {
	return u.distributorRepo.DeleteDistribution(distrbutionId, distributorId)
}

func (u *distributorUsecase) SearchDistributions(search string) ([]dto.GetDistribution, error) {
	return u.distributorRepo.SearchDistributions(search)
}

func (u *distributorUsecase) GetDistributionsByDistributorId(id uint) ([]dto.GetDistribution, error) {
	return u.distributorRepo.GetDistributionsByDistributorId(id)
}

func (u *distributorUsecase) GetDistributionByid(id uint) (*dto.GetDistributionById, error) {
	return u.distributorRepo.GetDistributionByid(id)
}

func (u *distributorUsecase) UpdateStatusOfDistribution(input *dto.UpdateStatusDistributionRequest) error {
	return u.distributorRepo.UpdateStatusOfDistribution(input)
}

func (u *distributorUsecase) ApprovedRetailerCartForRetailer(input *dto.ApprovedRetailerCart) error {
	return u.distributorRepo.ApprovedRetailerCartForRetailer(input)
}
