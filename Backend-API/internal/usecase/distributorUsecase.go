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

type DistributorUsecase interface {
	CreateDistribution(ctx context.Context, input *dto.CreateDistribution) error
	UpdateDistribution(ctx context.Context, input *dto.UpdateDistribution) error
	DeleteDistribution(ctx context.Context, distrbutionId, distributorId uint) error
	UpdateStatusOfDistribution(ctx context.Context, input *dto.UpdateStatusDistribution) error

	//accept
	AcceptSeller(ctx context.Context, distributorProfileId, sellerBoxId uint) error

	//get
	GetListDistributionFYP(ctx context.Context) ([]dto.GetListDistribution, error)
	SearchDistributions(ctx context.Context, search string) ([]dto.GetListDistribution, error)
	GetListDistributionsByDistributorId(ctx context.Context, id uint) ([]dto.GetListDistribution, error)
	GetDistributionById(ctx context.Context, id uint) (*dto.GetDistributionById, error)
}

type distributorUsecase struct {
	distributorRepo repository.DistributorRepository
	redis           *redis.Client
}

func NewDistributorUsecase(distributorRepo repository.DistributorRepository, redis *redis.Client) DistributorUsecase {
	return &distributorUsecase{distributorRepo, redis}
}

func (u *distributorUsecase) CreateDistribution(ctx context.Context, input *dto.CreateDistribution) error {
	return u.distributorRepo.CreateDistribution(ctx, input)
}

func (u *distributorUsecase) UpdateDistribution(ctx context.Context, input *dto.UpdateDistribution) error {
	return u.distributorRepo.UpdateDistribution(ctx, input)
}

func (u *distributorUsecase) DeleteDistribution(ctx context.Context, distrbutionId, distributorId uint) error {
	return u.distributorRepo.DeleteDistribution(ctx, distrbutionId, distributorId)
}

func (u *distributorUsecase) SearchDistributions(ctx context.Context, search string) ([]dto.GetListDistribution, error) {
	return u.distributorRepo.SearchDistributions(ctx, search)
}

func (u *distributorUsecase) GetListDistributionsByDistributorId(ctx context.Context, id uint) ([]dto.GetListDistribution, error) {
	return u.distributorRepo.GetListDistributionsByDistributorId(ctx, id)
}

func (u *distributorUsecase) GetDistributionById(ctx context.Context, id uint) (*dto.GetDistributionById, error) {
	return u.distributorRepo.GetDistributionById(ctx, id)
}

func (u *distributorUsecase) UpdateStatusOfDistribution(ctx context.Context, input *dto.UpdateStatusDistribution) error {
	return u.distributorRepo.UpdateStatusOfDistribution(ctx, input)
}

func (u *distributorUsecase) AcceptSeller(ctx context.Context, distributorProfileId, sellerBoxId uint) error {
	bcReq, err := u.distributorRepo.AcceptSeller(ctx, distributorProfileId, sellerBoxId)
	if err != nil {
		return err
	}
	data := map[string]interface{}{
		"op":   "bc_seller",
		"data": bcReq,
	}
	dataJson, _ := json.Marshal(data)

	key := fmt.Sprintln("behind:pending:seller")
	if err := u.redis.HSet(ctx, key, "item_"+time.Now().Format("150405.000"), dataJson).Err(); err != nil {
		return err
	}

	u.redis.Expire(ctx, key, 10*time.Second)
	return nil
}

func (u *distributorUsecase) GetListDistributionFYP(ctx context.Context) ([]dto.GetListDistribution, error) {
	return u.distributorRepo.GetListDistributionFYP(ctx)
}
