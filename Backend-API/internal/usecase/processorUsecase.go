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

type ProcessorUsecase interface {
	CreateProcessor(ctx context.Context, input *dto.CreateProcessor) error
	UpdateProcessor(ctx context.Context, input *dto.UpdateProcessor) error
	DeleteProcessor(ctx context.Context, procesorId, processoProfileId uint) error

	//accept
	AcceptDistributor(ctx context.Context, processorProfileId, distributorHarvestId uint) error

	//get
	ListHarvestProcessorByProcessorId(ctx context.Context, processorId uint) ([]dto.GetListHarvestProcessor, error)
	SearchHarvestProcessor(ctx context.Context, search string) ([]dto.GetListHarvestProcessor, error)
	ListHarvestProcessorFYP(ctx context.Context) ([]dto.GetListHarvestProcessor, error)
	GetHarvestProcessorById(ctx context.Context, id uint) (*dto.GetHarvestProcessorById, error)
}

type processorUsecase struct {
	processorRepo repository.ProcessorRepository
	redis         *redis.Client
}

func NewProcessoUsecase(processorRepo repository.ProcessorRepository, redis *redis.Client) ProcessorUsecase {
	return &processorUsecase{processorRepo, redis}
}

func (u *processorUsecase) CreateProcessor(ctx context.Context, input *dto.CreateProcessor) error {

	return u.processorRepo.CreateProcessor(ctx, input)
}

func (u *processorUsecase) UpdateProcessor(ctx context.Context, input *dto.UpdateProcessor) error {

	return u.processorRepo.UpdateProcessor(ctx, input)
}

func (u *processorUsecase) DeleteProcessor(ctx context.Context, processorId, processorData uint) error {

	return u.processorRepo.DeleteProcessor(ctx, processorId, processorData)
}

func (u *processorUsecase) AcceptDistributor(ctx context.Context, processorProfileId, distributorHarvestId uint) error {
	bcReq, err := u.processorRepo.AcceptDistributor(ctx, processorProfileId, distributorHarvestId)
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

func (u *processorUsecase) ListHarvestProcessorByProcessorId(ctx context.Context, processorId uint) ([]dto.GetListHarvestProcessor, error) {

	return u.processorRepo.ListHarvestProcessorByProcessorId(ctx, processorId)
}

func (u *processorUsecase) SearchHarvestProcessor(ctx context.Context, search string) ([]dto.GetListHarvestProcessor, error) {

	return u.processorRepo.SearchHarvestProcessor(ctx, search)
}

func (u *processorUsecase) ListHarvestProcessorFYP(ctx context.Context) ([]dto.GetListHarvestProcessor, error) {
	return u.processorRepo.ListHarvestProcessorFYP(ctx)
}

func (u *processorUsecase) GetHarvestProcessorById(ctx context.Context, id uint) (*dto.GetHarvestProcessorById, error) {
	return u.processorRepo.GetHarvestProcessorById(ctx, id)
}
