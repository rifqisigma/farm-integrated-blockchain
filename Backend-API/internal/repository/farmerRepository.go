package repository

import (
	"context"
	"encoding/json"
	"errors"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/entity"
	"farm-integrated-web3/utils/helper"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type FarmerRepository interface {
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

type farmerRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewFarmerRepository(db *gorm.DB, redis *redis.Client) FarmerRepository {
	return &farmerRepository{db, redis}
}

func (r *farmerRepository) CreateHarvest(ctx context.Context, input *dto.HarvestRequest) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var regency entity.Regency
		res := tx.WithContext(ctx).Where("name = ?", input.CountryRequest.RegionRequest.RegencyRequest.Name).First(&regency)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			var country entity.Country
			if err := tx.WithContext(ctx).WithContext(ctx).Where("name = ?", input.RegionRequest.Name).
				FirstOrCreate(&country, entity.Country{Name: input.CountryRequest.Name}).Error; err != nil {
				return err
			}

			var region entity.Region
			if err := tx.WithContext(ctx).Where("name = ? AND country_id = ?", input.CountryRequest.RegionRequest.Name, country.ID).
				FirstOrCreate(&region, entity.Region{Name: input.CountryRequest.RegionRequest.Name, CountryID: country.ID}).Error; err != nil {
				return err
			}

			regency = entity.Regency{Name: input.CountryRequest.RegionRequest.RegencyRequest.Name, RegionID: region.ID}
			if err := tx.WithContext(ctx).Create(&regency).Error; err != nil {
				return err
			}
		}

		var dataTime dto.ValidateTimeHarvest
		err := tx.Model(&entity.Harvest{}).
			Select("update_time as harvest_time", "crops.harvest_time_span as harvest_time_span").
			Joins("JOIN crops ON crops.id = harvests.crop_id").
			Where("farmer_profile_id = ?", input.FarmerProfileId).
			Order("harvests.update_time desc").
			First(&dataTime).Error

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if err == nil {
			rentang := time.Duration(dataTime.HarvestTimeSpan) * 24 * time.Hour
			if time.Since(dataTime.HarvestTime) <= rentang {
				return helper.ErrInvalidTime
			}
		}

		newHarvest := entity.Harvest{
			FarmerProfileId: input.FarmerProfileId,
			CropId:          input.CropID,
			Quantity:        input.Quantity,
			BasePrice:       input.BasePrice,
			Name:            input.Name,
			RegencyId:       regency.ID,
		}

		if err := tx.WithContext(ctx).Create(&newHarvest).Error; err != nil {
			return err
		}

		keyFarmer := fmt.Sprintf("farm:farmer:%d", newHarvest.FarmerProfileId)
		keyHarvest := fmt.Sprintf("farm:%d", newHarvest.ID)
		keyFYP := fmt.Sprintln("farm:fyp")
		_, errRedis := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
			p.Del(ctx, keyFarmer)
			p.Del(ctx, keyHarvest)
			p.Del(ctx, keyFYP)
			keys, _ := r.redis.Keys(ctx, "farm:search:*").Result()
			if len(keys) > 0 {
				r.redis.Del(ctx, keys...)
			}
			return nil
		})

		if errRedis != nil {
			return err
		}

		return nil
	})
}

func (r *farmerRepository) UpdateHarvest(ctx context.Context, input *dto.HarvestUpdate) error {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var regency entity.Regency
	if err := tx.WithContext(ctx).Where("name = ?", input.CountryRequest.RegionRequest.RegencyRequest.Name).First(&regency).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			var country entity.Country
			if err := tx.WithContext(ctx).Where("name = ?", input.CountryRequest.Name).
				FirstOrCreate(&country, entity.Country{Name: input.CountryRequest.Name}).Error; err != nil {
				return err
			}

			var region entity.Region
			if err := tx.WithContext(ctx).Where("name = ? AND country_id = ?", input.CountryRequest.RegionRequest.Name, country.ID).
				FirstOrCreate(&region, entity.Region{Name: input.CountryRequest.RegionRequest.Name, CountryID: country.ID}).Error; err != nil {
				return err
			}

			regency = entity.Regency{Name: input.CountryRequest.RegionRequest.RegencyRequest.Name, RegionID: region.ID}
			if err := tx.WithContext(ctx).Create(&regency).Error; err != nil {
				return err
			}

		} else {
			return err
		}

	}

	var validation dto.DataValidationFarm
	res := tx.Model(&entity.Harvest{}).Select("create_time, quantity").Where("id = ? AND is_canceled = ?", input.HarvestId, false).Scan(&validation)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	if time.Since(validation.Time) > 30*time.Minute {
		tx.Rollback()
		return helper.ErrInvalidTime
	}

	if input.Quantity > validation.Quantity {
		tx.Rollback()
		return gorm.ErrInvalidData
	}

	updatedHarvest := make(map[string]interface{})
	if input.Quantity > 0 {
		updatedHarvest["quantity"] = input.Quantity
	}

	if input.BasePrice > 0 {
		updatedHarvest["base_price"] = input.BasePrice
	}

	if input.Name != "" {
		updatedHarvest["name"] = input.Name
	}

	if regency.ID != 0 {
		updatedHarvest["regency_id"] = regency.ID
	}

	err2 := tx.Model(&entity.Harvest{}).Where("id = ? AND farmer_profile_id = ?", input.HarvestId, input.FarmerProfileId).UpdateColumns(&updatedHarvest)
	if err2.Error != nil {
		tx.Rollback()
		return err2.Error
	}

	if err2.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	keyFarmer := fmt.Sprintf("farm:farmer:%d", input.FarmerProfileId)
	keyHarvest := fmt.Sprintf("farm:%d", input.HarvestId)
	keyFYP := fmt.Sprintln("farm:fyp")
	_, err := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keyFarmer)
		p.Del(ctx, keyHarvest)
		p.Del(ctx, keyFYP)
		keys, _ := r.redis.Keys(ctx, "farm:search:*").Result()
		if len(keys) > 0 {
			r.redis.Del(ctx, keys...)
		}
		return nil
	})

	if err != nil {
		return err
	}

	return tx.Commit().Error
}

func (r *farmerRepository) DeleteHarvest(ctx context.Context, farmerProfileId, harvestId uint) error {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()
	var validationTime time.Time
	res := tx.Model(&entity.Harvest{}).Select("create_time").Where("id = ?", harvestId).Scan(&validationTime)
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

	err2 := r.db.WithContext(ctx).Model(&entity.Harvest{}).Where("id = ? AND farmer_profile_id ", harvestId, farmerProfileId).UpdateColumn("is_canceled", true)
	if err2.Error != nil {
		return err2.Error
	}

	if err2.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	keyFarmer := fmt.Sprintf("farm:farmer:%d", farmerProfileId)
	keyHarvest := fmt.Sprintf("farm:%d", harvestId)
	keyFYP := fmt.Sprintln("farm:fyp")
	_, err := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keyFarmer)
		p.Del(ctx, keyHarvest)
		p.Del(ctx, keyFYP)
		keys, _ := r.redis.Keys(ctx, "farm:search:*").Result()
		if len(keys) > 0 {
			r.redis.Del(ctx, keys...)
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *farmerRepository) AcceptedFarmerForDistributor(ctx context.Context, input *dto.AcceptFarmerForDistributor) error {

	update := make(map[string]interface{})
	update["approved_by_farmer"] = input.Accepted
	if !input.Accepted {
		update["is_canceled"] = true
	}
	res := r.db.WithContext(ctx).Model(&entity.Distribution{}).Where("distributions.id = ? AND distribution.farmer_profile_id = ? AND is_canceled = ?", input.DistributionId, input.FarmerProfieId, false).UpdateColumns(update)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	keyDistributor := fmt.Sprintf("distribution:distributor:%d", input.DistributorProfileId)
	keyDistribution := fmt.Sprintf("distribution:%d", input.DistributionId)
	_, errRedis := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keyDistribution)
		p.Del(ctx, keyDistributor)
		return nil
	})
	if errRedis != nil {
		return errRedis
	}

	return nil
}

func (r *farmerRepository) ListHarvestByFarmerId(ctx context.Context, farmerId uint) ([]dto.GetListHarvest, error) {

	key := fmt.Sprintf("farm:farmer:%d", farmerId)
	var harvests []dto.GetListHarvest

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &harvests); err != nil {
			return harvests, nil
		}
	}

	res := r.db.WithContext(ctx).Model(&entity.Harvest{}).
		Select(
			"cp.crop as crop_name",
			"fp.name as farmer_name",
			"harvests.id as id",
			"harvests.update_time as time",
			"harvests.base_price as base_price",
			"rg.name as regency_name").
		Joins("JOIN crops as cp ON cp.id = harvests.crop_id").
		Joins("JOIN farmer_profiles as fp ON fp.id = harvests.farmer_profile_id").
		Joins("JOIN regencies as rg ON rg.id = harvests.regency_id").
		Where("harvests.farmer_profile_id = ?", farmerId).
		Scan(&harvests)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if len(harvests) == 0 {
		return harvests, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(&harvests)
	r.redis.Set(ctx, key, jsonData, 5*time.Minute)
	return harvests, nil
}

func (r *farmerRepository) ListHarvestFYP(ctx context.Context) ([]dto.GetListHarvest, error) {
	var harvests []dto.GetListHarvest
	key := fmt.Sprintln("farm:fyp")

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &harvests); err == nil {
			return harvests, nil
		}
	}

	res := r.db.WithContext(ctx).Model(&entity.Harvest{}).
		Select(
			"cp.crop as crop_name",
			"fp.name as farmer_name",
			"harvests.id as id",
			"harvests.update_time as time",
			"harvests.base_price",
			"rg.name as regency_name",
		).
		Joins("JOIN crops as cp ON cp.id = harvests.crop_id").
		Joins("JOIN farmer_profiles as fp ON fp.id = harvests.farmer_profile_id").
		Joins("JOIN regencies as rg ON rg.id = harvests.regency_id").
		Order("RAND()").
		Scan(&harvests)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if len(harvests) == 0 {
		return harvests, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(harvests)
	r.redis.Set(ctx, key, jsonData, 5*time.Minute)
	return harvests, nil
}

func (r *farmerRepository) HarvestById(ctx context.Context, harvestId uint) (*dto.GetHarvestById, error) {
	var harvest dto.GetHarvestById
	key := fmt.Sprintf("farm:%d", harvestId)

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &harvest); err == nil {
			return &harvest, nil
		}
	}
	res := r.db.WithContext(ctx).Model(&entity.Harvest{}).
		Select(
			"cp.crop as crop_name",
			"fp.name as farmer_name",
			"harvests.id as id",
			"harvests.base_price as base_price",
			"harvests.update_time as time",
			"harvests.quantity as quantity",
			"rg.name as regency_name",
			"rn.name as region_name",
			"cn.name as country_name",
		).
		Joins("JOIN crops as cp ON cp.id = harvests.crop_id").
		Joins("JOIN farmer_profiles as fp ON fp.id = harvests.farmer_profile_id").
		Joins("JOIN regencies as rg ON rg.id = harvests.regency_id").
		Joins("JOIN regions as rn ON rn.id = rg.region_id").
		Joins("JOIN countries as cn ON cn.id = rn.country_id").
		Where("harvests.id = ?", harvestId).
		Scan(&harvest)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	jsonData, _ := json.Marshal(harvest)
	r.redis.Set(ctx, key, jsonData, 5*time.Minute)

	return &harvest, nil
}

func (r *farmerRepository) SearchHarvest(ctx context.Context, search string) ([]dto.GetListHarvest, error) {
	var harvests []dto.GetListHarvest
	key := fmt.Sprintf("farm:search:%s", search)

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &harvests); err != nil {
			return harvests, nil
		}
	}

	input := "%" + search + "%"
	res := r.db.WithContext(ctx).Model(&entity.Harvest{}).
		Select(
			"cp.crop as crop_name",
			"fp.name as farmer_name",
			"harvests.id as id",
			"harvests.update_time as time",
			"harvests.base_price as base_price",
			"rg.name as regency_name",
		).
		Joins("JOIN crops as cp ON cp.id = harvests.crop_id").
		Joins("JOIN farmer_profiles as fp ON fp.id = harvests.farmer_profile_id").
		Joins("JOIN regencies as rg ON rg.id = harvests.regency_id").
		Where("cp.crop LIKE ? OR fp.name LIKE ?", input, input).
		Scan(&harvests)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if len(harvests) == 0 {
		return harvests, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(harvests)
	if err := r.redis.Set(ctx, key, jsonData, 5*time.Minute).Err(); err != nil {
		return nil, err
	}

	return harvests, nil
}
