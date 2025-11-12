package usecase

import (
	"context"
	"encoding/json"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/internal/repository"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type FarmerUsecase interface {
	CreateHarvest(ctx context.Context, input *dto.HarvestCreate) error
	UpdateHarvest(ctx context.Context, input *dto.HarvestUpdate) error
	DeleteHarvest(ctx context.Context, farmerProfileId, harvestId uint) error
	UpdateStatusHarvest(ctx context.Context, farmerProfileId, harvestId uint) error

	//accept
	AcceptHarvestCollector(ctx context.Context, farmerId, collectorHarvestId uint) error
	AcceptHarvestProcessor(ctx context.Context, farmerId, processorHarvestId uint) error
	AcceptDistributor(ctx context.Context, farmerId, distributorHarvestId uint) error

	//get
	ListHarvestFYP(ctx context.Context) ([]dto.GetListHarvest, error)
	ListHarvestByFarmerId(ctx context.Context, farmerId uint) ([]dto.GetListHarvest, error)
	HarvestById(ctx context.Context, harvestId uint) (*dto.GetHarvestById, error)
	SearchHarvest(ctx context.Context, search string) ([]dto.GetListHarvest, error)
}

type farmerUsecase struct {
	farmerRepo repository.FarmerRepository

	redis *redis.Client
}

func NewFarmerUsecase(farmerRepo repository.FarmerRepository, redis *redis.Client) FarmerUsecase {
	return &farmerUsecase{farmerRepo, redis}
}

func (u *farmerUsecase) CreateHarvest(ctx context.Context, input *dto.HarvestCreate) error {
	err := u.farmerRepo.CreateHarvest(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (u *farmerUsecase) UpdateHarvest(ctx context.Context, input *dto.HarvestUpdate) error {
	err := u.farmerRepo.UpdateHarvest(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (u *farmerUsecase) AcceptHarvestCollector(ctx context.Context, farmerId, collectorHarvestId uint) error {
	bcReq, err := u.farmerRepo.AcceptHarvestCollector(ctx, farmerId, collectorHarvestId)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"op":   "bc_collector",
		"data": bcReq,
	}
	dataJson, _ := json.Marshal(data)

	key := fmt.Sprintln("behind:pending:harvest_collector")
	if err := u.redis.HSet(ctx, key, "item_"+time.Now().Format("150405.000"), dataJson).Err(); err != nil {
		return err
	}

	u.redis.Expire(ctx, key, 10*time.Second)
	return nil
}
func (u *farmerUsecase) AcceptHarvestProcessor(ctx context.Context, farmerId, processorHarvestId uint) error {
	bcReq, err := u.farmerRepo.AcceptHarvestProcessor(ctx, farmerId, processorHarvestId)
	if err != nil {
		return err
	}
	data := map[string]interface{}{
		"op":   "bc_processor",
		"data": bcReq,
	}
	dataJson, _ := json.Marshal(data)

	key := fmt.Sprintln("behind:pending:harvest_processor")
	if err := u.redis.HSet(ctx, key, "item_"+time.Now().Format("150405.000"), dataJson).Err(); err != nil {
		return err
	}
	u.redis.Expire(ctx, key, 10*time.Second)
	return nil
}
func (u *farmerUsecase) AcceptDistributor(ctx context.Context, farmerId, distributorHarvestId uint) error {
	bcReq, err := u.farmerRepo.AcceptDistributor(ctx, farmerId, distributorHarvestId)
	if err != nil {
		return err
	}
	data := map[string]interface{}{
		"op":   "bc_distribution",
		"data": bcReq,
	}
	dataJson, _ := json.Marshal(data)

	key := fmt.Sprintln("behind:pending:distribution")
	if err := u.redis.HSet(ctx, key, "item_"+time.Now().Format("150405.000"), dataJson).Err(); err != nil {
		return err
	}
	u.redis.Expire(ctx, key, 10*time.Second)
	return nil
}

func (u *farmerUsecase) DeleteHarvest(ctx context.Context, farmerProfileId, harvestId uint) error {

	return u.farmerRepo.DeleteHarvest(ctx, farmerProfileId, harvestId)
}

func (u *farmerUsecase) UpdateStatusHarvest(ctx context.Context, farmerProfileId, harvestId uint) error {
	bcReq, err := u.farmerRepo.UpdateStatusHarvest(ctx, farmerProfileId, harvestId)
	if err != nil {
		return err
	}
	data := map[string]interface{}{
		"op":   "bc_harvest",
		"data": bcReq,
	}
	dataJson, _ := json.Marshal(data)

	key := fmt.Sprintln("behind:pending:harvest_farmer")
	if err := u.redis.HSet(ctx, key, "item_"+time.Now().Format("150405.000"), dataJson).Err(); err != nil {
		return err
	}

	u.redis.Expire(ctx, key, 10*time.Second)
	return nil
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
