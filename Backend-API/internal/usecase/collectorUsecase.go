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

type CollectorUsecase interface {
	CreateHarvestCollector(ctx context.Context, input *dto.CreateCollector) error
	UpdateHarvestCollector(ctx context.Context, input *dto.UpdateCollector) error
	DeleteHarvestCollector(ctx context.Context, collectorId, collectorProfileId uint) error

	//accept
	AcceptHarvestProcessor(ctx context.Context, collectorProfileId, processorId uint) error
	AcceptDistributor(ctx context.Context, collectorProfileId, distributorId uint) error

	//get
	ListHarvestCollectorByCollectorId(ctx context.Context, collectorId uint) ([]dto.GetListHarvestCollector, error)
	ListHarvestCollectorFYP(ctx context.Context) ([]dto.GetListHarvestCollector, error)
	SearchHarvestCollector(ctx context.Context, search string) ([]dto.GetListHarvestCollector, error)
	GetHarvestCollectorById(ctx context.Context, id uint) (*dto.GetHarvestCollectorById, error)
}

type collectorUsecase struct {
	collectorRepo repository.CollectorRepository
	redis         *redis.Client
}

func NewCollectorUsecase(collectorRepo repository.CollectorRepository, redis *redis.Client) CollectorUsecase {
	return &collectorUsecase{collectorRepo, redis}
}

func (u *collectorUsecase) CreateHarvestCollector(ctx context.Context, input *dto.CreateCollector) error {
	return u.collectorRepo.CreateHarvestCollector(ctx, input)
}

func (u *collectorUsecase) UpdateHarvestCollector(ctx context.Context, input *dto.UpdateCollector) error {
	return u.collectorRepo.UpdateHarvestCollector(ctx, input)
}

func (u *collectorUsecase) DeleteHarvestCollector(ctx context.Context, collectorId, collectorProfileId uint) error {
	return u.collectorRepo.DeleteHarvestCollector(ctx, collectorId, collectorProfileId)
}

func (u *collectorUsecase) AcceptHarvestProcessor(ctx context.Context, collectorProfileId, processorId uint) error {
	bcReq, err := u.collectorRepo.AcceptHarvestProcessor(ctx, collectorProfileId, processorId)
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

func (u *collectorUsecase) AcceptDistributor(ctx context.Context, collectorProfileId, distributorId uint) error {
	bcReq, err := u.collectorRepo.AcceptDistributor(ctx, collectorProfileId, distributorId)
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

func (u *collectorUsecase) ListHarvestCollectorByCollectorId(ctx context.Context, collectorId uint) ([]dto.GetListHarvestCollector, error) {
	return u.collectorRepo.ListHarvestCollectorByCollectorId(ctx, collectorId)
}

func (u *collectorUsecase) ListHarvestCollectorFYP(ctx context.Context) ([]dto.GetListHarvestCollector, error) {
	return u.collectorRepo.ListHarvestCollectorFYP(ctx)
}

func (u *collectorUsecase) SearchHarvestCollector(ctx context.Context, search string) ([]dto.GetListHarvestCollector, error) {
	return u.collectorRepo.SearchHarvestCollector(ctx, search)
}

func (u *collectorUsecase) GetHarvestCollectorById(ctx context.Context, id uint) (*dto.GetHarvestCollectorById, error) {
	return u.collectorRepo.GetHarvestCollectorById(ctx, id)
}
