package repository

import (
	"context"
	"encoding/json"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/entity"
	"farm-integrated-web3/utils/helper"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type DistributorRepository interface {
	CreateDistribution(ctx context.Context, input *dto.CreateDistributionRequest) error
	UpdateDistribution(ctx context.Context, input *dto.UpdateDistributionRequest) error
	DeleteDistribution(ctx context.Context, distrbutionId, distributorId uint) error
	UpdateStatusOfDistribution(ctx context.Context, input *dto.UpdateStatusDistributionRequest) error
	ApprovedRetailerCartForRetailer(ctx context.Context, input *dto.ApprovedRetailerCart) error
	//get
	SearchDistributions(ctx context.Context, search string) ([]dto.GetDistribution, error)
	GetDistributionsByDistributorId(ctx context.Context, id uint) ([]dto.GetDistribution, error)
	GetDistributionByid(ctx context.Context, id uint) (*dto.GetDistributionById, error)
	GetDistributionFYP(ctx context.Context) ([]dto.GetDistribution, error)
}

type distributorRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewDistributorRepository(db *gorm.DB, redis *redis.Client) DistributorRepository {
	return &distributorRepository{db, redis}
}

func (r *distributorRepository) CreateDistribution(ctx context.Context, input *dto.CreateDistributionRequest) error {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var quantityHarvest float64
	res := tx.Model(&entity.Harvest{}).Select("quantity").Where("id = ? AND approved_by_farmer = ?", input.HarvestId, true).Scan(&quantityHarvest)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	if res.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	if quantityHarvest > input.Quantity {
		tx.Rollback()
		return gorm.ErrInvalidData
	}

	newDistribution := entity.Distribution{
		FinalPrice:           input.FinalPrice,
		MarkUpPrice:          input.MarkupPrice,
		HarvestId:            input.HarvestId,
		FarmerProfileId:      input.FarmerProfileId,
		DistributorProfileId: input.DistributorProfileId,
	}
	if err := tx.Model(&entity.Distribution{}).Create(&newDistribution).Error; err != nil {
		tx.Rollback()
	}

	keyDistributor := fmt.Sprintf("distribution:distributor:%d", newDistribution.DistributorProfileId)
	keyDistribution := fmt.Sprintf("distribution:%d", newDistribution.ID)
	keyFYP := fmt.Sprintln("distribution:fyp")
	_, errRedis := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keyDistribution)
		p.Del(ctx, keyDistributor)
		p.Del(ctx, keyFYP)
		keys, _ := r.redis.Keys(ctx, "distribution:search:*").Result()
		if len(keys) > 0 {
			r.redis.Del(ctx, keys...)
		}
		return nil
	})

	if errRedis != nil {
		return errRedis
	}

	return tx.Commit().Error
}

func (r *distributorRepository) UpdateDistribution(ctx context.Context, input *dto.UpdateDistributionRequest) error {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()
	var validation dto.DataValidationDistribution
	res := tx.Model(&entity.Harvest{}).Select("quantity, create_time").Where("id = ? AND is_canceled = ?", input.DistributionId, false).Scan(&validation)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	if input.Quantity > validation.Quantity {
		tx.Rollback()
		return gorm.ErrInvalidData
	}

	if time.Since(validation.Time) > 30*time.Minute {
		tx.Rollback()
		return helper.ErrInvalidTime
	}

	updateDistribution := map[string]interface{}{
		"final_price":   input.FinalPrice,
		"mark_up_price": input.MarkupPrice,
	}

	if input.Quantity > 0 {
		updateDistribution["quantity"] = input.Quantity
	}

	res2 := tx.Model(&entity.Distribution{}).Where("id = ? AND distributor_profile_id = ? ", input.DistributionId, input.DistributorProfileId).UpdateColumns(&updateDistribution)
	if res2.Error != nil {
		tx.Rollback()
		return res.Error
	}

	if res2.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	keyDistributor := fmt.Sprintf("distribution:distributor:%d", input.DistributorProfileId)
	keyDistribution := fmt.Sprintf("distribution:%d", input.DistributionId)
	keyFYP := fmt.Sprintln("distribution:fyp")
	_, errRedis := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keyDistribution)
		p.Del(ctx, keyDistributor)
		p.Del(ctx, keyFYP)
		keys, _ := r.redis.Keys(ctx, "distribution:search:*").Result()
		if len(keys) > 0 {
			r.redis.Del(ctx, keys...)
		}
		return nil
	})

	if errRedis != nil {
		return errRedis
	}

	return tx.Commit().Error
}

func (r *distributorRepository) DeleteDistribution(ctx context.Context, distrbutionId, distributorId uint) error {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()
	var validationTime time.Time
	res := tx.Model(&entity.Distribution{}).Select("create_time").Where("id = ? AND is_canceled = ? ", distributorId, false).Scan(&validationTime)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	if time.Since(validationTime) > 30*time.Minute {
		tx.Rollback()
		return helper.ErrInvalidTime
	}

	res2 := tx.Model(&entity.Distribution{}).Where("id = ? AND distributor_profile_id = ? ", distrbutionId, distributorId).UpdateColumn("is_canceled", true)
	if res2.Error != nil {
		return res2.Error
	}

	if res2.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	keyDistributor := fmt.Sprintf("distribution:distributor:%d", distributorId)
	keyDistribution := fmt.Sprintf("distribution:%d", distrbutionId)
	keyFYP := fmt.Sprintln("distribution:fyp")
	_, errRedis := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keyDistribution)
		p.Del(ctx, keyDistributor)
		p.Del(ctx, keyFYP)
		keys, _ := r.redis.Keys(ctx, "distribution:search:*").Result()
		if len(keys) > 0 {
			r.redis.Del(ctx, keys...)
		}
		return nil
	})

	if errRedis != nil {
		return errRedis
	}

	return nil

}

func (r *distributorRepository) UpdateStatusOfDistribution(ctx context.Context, input *dto.UpdateStatusDistributionRequest) error {
	err := r.db.WithContext(ctx).Model(&entity.Distribution{}).Where("id = ? AND distributor_profile_id = ? AND approved_by_farmer = ? AND is_canceled = ? ", input.DistributionId, input.DistributorProfileId, true, false).UpdateColumn("status_distribution", input.Status)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil

}

func (r *distributorRepository) ApprovedRetailerCartForRetailer(ctx context.Context, input *dto.ApprovedRetailerCart) error {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	update := make(map[string]interface{})
	update["approved_by_distributor"] = input.Approved
	if !input.Approved {
		update["is_canceled"] = true
	}
	res := tx.Model(&entity.RetailerCart{}).Where("id = ? AND distributor_profile_id = ? AND is_canceled = ?", input.RetailerCartId, input.DistributorProfileId, false).UpdateColumns(update)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *distributorRepository) SearchDistributions(ctx context.Context, search string) ([]dto.GetDistribution, error) {
	var result []dto.GetDistribution
	input := "%" + search + "%"
	key := fmt.Sprintf("distribution:search:%s", search)

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err == nil {
			return result, nil
		}
	}

	res := r.db.WithContext(ctx).Model(&entity.Distribution{}).
		Select(
			"distributions.id as id",
			"dp.name as ditributor_name",
			"fp.name as farmer_name",
			"distributions.final_price as final_price",
			"cp.crop as crop_name",
			"distributions.update_time as time",
			"rg.name as regency_name",
		).
		Joins("JOIN harvests as hs ON hs.id = distributions.harvest_id").
		Joins("JOIN crops as cs ON cs.id = hs.crop_id").
		Joins("JOIN farmer_profiles as fp ON fp.id = distributions.farmer_profile_id").
		Joins("JOIN distributor_profiles as dp ON dp.id = distributions.distributor_profile_id").
		Joins("JOIN regencies as rg ON rg.id = distributions.destination_id").
		Where("cs.crop LIKE ? OR dp.name LIKE ? OR cp.crop LIKE ? AND distributions.approved_by_farmer = ?", input, input, input, true).
		Scan(&result)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if len(result) == 0 {
		return nil, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(result)
	if err := r.redis.Set(ctx, key, jsonData, 5*time.Minute).Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *distributorRepository) GetDistributionsByDistributorId(ctx context.Context, id uint) ([]dto.GetDistribution, error) {
	var result []dto.GetDistribution
	key := fmt.Sprintf("distribution:distributor:%d", id)

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err == nil {
			return result, nil
		}
	}

	res := r.db.WithContext(ctx).Model(&entity.Distribution{}).
		Select(
			"distributions.id as id",
			"dp.name as ditributor_name",
			"fp.name as farmer_name",
			"distributions.final_price as final_price",
			"cp.crop as crop_name",
			"distributions.update_time as time",
			"rg.name as regency_name",
		).
		Joins("JOIN harvests as hs ON hs.id = distributions.harvest_id").
		Joins("JOIN crops as cs ON cs.id = hs.crop_id").
		Joins("JOIN farmer_profiles as fp ON fp.id = distributions.farmer_profile_id").
		Joins("JOIN distributor_profiles as dp ON dp.id = distributions.distributor_profile_id").
		Joins("JOIN regencies as rg ON rg.id = distributions.destination_id").
		Where("distributions.distributor_profile_id = ?", id).
		Scan(&result)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if len(result) == 0 {
		return nil, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(result)
	if err := r.redis.Set(ctx, key, jsonData, 5*time.Minute).Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *distributorRepository) GetDistributionByid(ctx context.Context, id uint) (*dto.GetDistributionById, error) {
	var result dto.GetDistributionById

	key := fmt.Sprintf("distribution:%d", id)

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err == nil {
			return &result, nil
		}
	}

	res := r.db.WithContext(ctx).Model(&entity.Distribution{}).
		Select(
			"distributions.id as id",
			"dp.name as ditributor_name",
			"fp.name as farmer_name",
			"distributions.final_price as final_price",
			"cp.crop as crop_name",
			"distributions.block_hash as block_hash",
			"distributions.has_arrived as has_arrived",
			"distributions.update_time as time",
			"rg.name as regency_name",
			"rn.name as region_name",
			"cn.name as country_name",
		).
		Joins("JOIN harvests as hs ON hs.id = distributions.harvest_id").
		Joins("JOIN crops as cs ON cs.id = hs.crop_id").
		Joins("JOIN farmer_profiles as fp ON fp.id = distributions.farmer_profile_id").
		Joins("JOIN distributor_profiles as dp ON dp.id = distributions.distributor_profile_id").
		Joins("JOIN regencies as rg ON rg.id = distributions.destination_id").
		Joins("JOIN regions as rn ON rn.id = rg.region_id").
		Joins("JOIN countries as cn ON cn.id = rn.country_id").
		Where("distributions.id = ?", id).
		Scan(&result)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	jsonData, _ := json.Marshal(result)
	r.redis.Set(ctx, key, jsonData, 5*time.Minute)

	return &result, nil

}

func (r *distributorRepository) GetDistributionFYP(ctx context.Context) ([]dto.GetDistribution, error) {
	var result []dto.GetDistribution

	key := fmt.Sprintln("distribution:fyp")

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err == nil {
			return result, nil
		}
	}

	res := r.db.WithContext(ctx).Model(&entity.Distribution{}).
		Select(
			"distributions.id as id",
			"dp.name as ditributor_name",
			"fp.name as farmer_name", "distributions.final_price as final_price",
			"cp.crop as crop_name",
			"distributions.update_time as time",
			"rg.name as regency_name",
		).
		Joins("JOIN harvests as hs ON hs.id = distributions.harvest_id").
		Joins("JOIN crops as cs ON cs.id = hs.crop_id").
		Joins("JOIN farmer_profiles as fp ON fp.id = distributions.farmer_profile_id").
		Joins("JOIN distributor_profiles as dp ON dp.id = distributions.distributor_profile_id").
		Joins("JOIN regencies as rg ON rg.id = distributions.destination_id").
		Order("RAND()").
		Scan(&result)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if len(result) == 0 {
		return nil, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(result)
	r.redis.Set(ctx, key, jsonData, 5*time.Minute)

	return result, nil
}
