package usecase

import (
	"farm-integrated-web3/dto"
	"farm-integrated-web3/utils/helper"
	"fmt"
	"os"
)

type BlockchainUsecase interface {
	GetAllHarvest() ([]dto.GetHarvestBc, error)
	GetAllHarvestCollector() ([]dto.GetCollectorHarvestBc, error)
	GetAllHarvestProcessor() ([]dto.GetProcessorHarvestBc, error)
	GetAllDistribution() ([]dto.GetDistributionBc, error)
	GetAllSellerBox() ([]dto.GetSellerBoxBc, error)

	GetHarvestById(id int64) (*dto.BcHarvest, error)
	GetHarvestCollectorById(id int64) (*dto.GetCollectorHarvestBc, error)
	GetHarvestProcessorById(id int64) (*dto.GetProcessorHarvestBc, error)
	GetDistributionById(id int64) (dto.GetDistributionBc, error)
	GetSellerBoxById(id int64) (dto.GetSellerBoxBc, error)
}

type blockchainUsecase struct {
}

func NewBlockchainRepository() BlockchainUsecase {
	return &blockchainUsecase{}
}

func (b *blockchainUsecase) GetAllHarvest() ([]dto.GetHarvestBc, error) {
	var results []dto.GetHarvestBc
	url := os.Getenv("BLOCKCHAIN_API") + "/blockchain/harvest"
	if err := helper.FetchJSON(url, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (b *blockchainUsecase) GetAllHarvestCollector() ([]dto.GetCollectorHarvestBc, error) {
	var results []dto.GetCollectorHarvestBc
	url := os.Getenv("BLOCKCHAIN_API") + "/blockchain/harvest-collector"
	if err := helper.FetchJSON(url, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (b *blockchainUsecase) GetAllHarvestProcessor() ([]dto.GetProcessorHarvestBc, error) {
	var results []dto.GetProcessorHarvestBc
	url := os.Getenv("BLOCKCHAIN_API") + "/blockchain/harvest-processor"
	if err := helper.FetchJSON(url, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (b *blockchainUsecase) GetAllDistribution() ([]dto.GetDistributionBc, error) {
	var results []dto.GetDistributionBc
	url := os.Getenv("BLOCKCHAIN_API") + "/blockchain/distribution"
	if err := helper.FetchJSON(url, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (b *blockchainUsecase) GetAllSellerBox() ([]dto.GetSellerBoxBc, error) {
	var results []dto.GetSellerBoxBc
	url := os.Getenv("BLOCKCHAIN_API") + "/blockchain/seller-box"
	if err := helper.FetchJSON(url, &results); err != nil {
		return nil, err
	}
	return results, nil
}

// ==================== Get By ID ====================

func (b *blockchainUsecase) GetHarvestById(id int64) (*dto.BcHarvest, error) {
	var result dto.BcHarvest
	url := fmt.Sprintf("%s/blockchain/harvest/%d", os.Getenv("BLOCKCHAIN_API"), id)
	if err := helper.FetchJSON(url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (b *blockchainUsecase) GetHarvestCollectorById(id int64) (*dto.GetCollectorHarvestBc, error) {
	var result dto.GetCollectorHarvestBc
	url := fmt.Sprintf("%s/blockchain/harvest-collector/%d", os.Getenv("BLOCKCHAIN_API"), id)
	if err := helper.FetchJSON(url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (b *blockchainUsecase) GetHarvestProcessorById(id int64) (*dto.GetProcessorHarvestBc, error) {
	var result dto.GetProcessorHarvestBc
	url := fmt.Sprintf("%s/blockchain/harvest-processor/%d", os.Getenv("BLOCKCHAIN_API"), id)
	if err := helper.FetchJSON(url, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (b *blockchainUsecase) GetDistributionById(id int64) (dto.GetDistributionBc, error) {
	var result dto.GetDistributionBc
	url := fmt.Sprintf("%s/blockchain/distribution/%d", os.Getenv("BLOCKCHAIN_API"), id)
	if err := helper.FetchJSON(url, &result); err != nil {
		return dto.GetDistributionBc{}, err
	}
	return result, nil
}

func (b *blockchainUsecase) GetSellerBoxById(id int64) (dto.GetSellerBoxBc, error) {
	var result dto.GetSellerBoxBc
	url := fmt.Sprintf("%s/blockchain/seller-box/%d", os.Getenv("BLOCKCHAIN_API"), id)
	if err := helper.FetchJSON(url, &result); err != nil {
		return dto.GetSellerBoxBc{}, err
	}
	return result, nil
}
