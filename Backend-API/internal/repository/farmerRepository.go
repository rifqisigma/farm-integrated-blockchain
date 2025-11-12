package repository

import (
	"context"
	"encoding/json"
	"errors"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/entity"
	"farm-integrated-web3/utils/helper"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type FarmerRepository interface {
	CreateHarvest(ctx context.Context, input *dto.HarvestCreate) error
	UpdateHarvest(ctx context.Context, input *dto.HarvestUpdate) error
	DeleteHarvest(ctx context.Context, farmerProfileId, harvestId uint) error
	UpdateStatusHarvest(ctx context.Context, farmerProfileId, harvestId uint) (*dto.BcHarvest, error)

	//accept
	AcceptHarvestCollector(ctx context.Context, farmerId, collectorHarvestId uint) (*dto.BcHarvestCollector, error)
	AcceptHarvestProcessor(ctx context.Context, farmerId, processorHarvestId uint) (*dto.BcHarvestProcessor, error)
	AcceptDistributor(ctx context.Context, farmerId, distributorHarvestId uint) (*dto.BcDistribution, error)

	//updatebc
	UpdateBcBlockHarvest(ctx context.Context, harvestId uint, txBlock string) error

	//get
	ListHarvestFYP(ctx context.Context) ([]dto.GetListHarvest, error)
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

func (r *farmerRepository) CreateHarvest(ctx context.Context, input *dto.HarvestCreate) error {
	tx := r.db.Begin().WithContext(ctx)
	var regency entity.Regency
	res := tx.Debug().Model(&entity.Regency{}).Where("name = ?", input.CountryRequest.RegionRequest.RegencyRequest.Name).First(&regency)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		var country entity.Country
		if err := tx.WithContext(ctx).WithContext(ctx).Where("name = ?", input.RegionRequest.Name).
			FirstOrCreate(&country, entity.Country{Name: input.CountryRequest.Name}).Error; err != nil {
			return err
		}

		var region entity.Region
		if err := tx.WithContext(ctx).Where("name = ? AND country_id = ?", input.CountryRequest.RegionRequest.Name, country.ID).
			FirstOrCreate(&region, entity.Region{Name: input.CountryRequest.RegionRequest.Name, CountryId: country.ID}).Error; err != nil {
			return err
		}

		regency = entity.Regency{Name: input.CountryRequest.RegionRequest.RegencyRequest.Name, RegionId: region.ID}
		if err := tx.WithContext(ctx).FirstOrCreate(&regency).Error; err != nil {
			return err
		}
	} else if res.Error != nil {
		// kalau error bukan record not found, return errornya
		return res.Error
	}

	var crop entity.Crop
	if err := tx.WithContext(ctx).Where("crop_name = ? AND commodity = ?", strings.ToLower(input.CropCreate.Name), strings.ToLower(input.CropCreate.Commodity)).
		FirstOrCreate(&crop, entity.Crop{CropName: strings.ToLower(input.CropCreate.Name), Commodity: strings.ToLower(input.CropCreate.Commodity), HarvestTimeSpan: input.CropCreate.HarvestTimeSpan}).Error; err != nil {
		return err
	}

	newHarvest := entity.Harvest{
		FarmerProfileId: input.FarmerProfileId,
		CropId:          crop.ID,
		Quantity:        input.Quantity,
		BasePrice:       input.BasePrice,
		Name:            input.Name,
		Desc:            input.Desc,
		Status:          0,
		RegencyId:       regency.ID,
	}

	if err := tx.WithContext(ctx).Create(&newHarvest).Error; err != nil {
		return err
	}

	keyFarmer := fmt.Sprintf("farm:farmer:%d", newHarvest.FarmerProfileId)
	keyHarvest := fmt.Sprintf("farm:%d", newHarvest.ID)
	keyFYP := "farm:fyp"
	_, errRedis := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keyFarmer)
		p.Del(ctx, keyHarvest)
		p.Del(ctx, keyFYP)
		keys, _ := r.redis.Keys(ctx, "farm:search:*").Result()
		if len(keys) > 0 {
			p.Del(ctx, keys...)
		}
		return nil
	})

	if errRedis != nil {
		return errRedis
	}

	return tx.Commit().Error
}

func (r *farmerRepository) UpdateHarvest(ctx context.Context, input *dto.HarvestUpdate) error {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()
	var regency entity.Regency
	if input.RegionRequest.Name != "" {
		if err := tx.WithContext(ctx).Where("name = ?", input.CountryRequest.RegionRequest.RegencyRequest.Name).First(&regency).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				var country entity.Country
				if err := tx.WithContext(ctx).Where("name = ?", input.CountryRequest.Name).
					FirstOrCreate(&country, entity.Country{Name: input.CountryRequest.Name}).Error; err != nil {
					return err
				}

				var region entity.Region
				if err := tx.WithContext(ctx).Where("name = ? AND country_id = ?", input.CountryRequest.RegionRequest.Name, country.ID).
					FirstOrCreate(&region, entity.Region{Name: input.CountryRequest.RegionRequest.Name, CountryId: country.ID}).Error; err != nil {
					return err
				}

				regency = entity.Regency{Name: input.CountryRequest.RegionRequest.RegencyRequest.Name, RegionId: region.ID}
				if err := tx.WithContext(ctx).Create(&regency).Error; err != nil {
					return err
				}

			} else {
				return err
			}
		}
	}

	var quantity float64
	res := tx.Debug().Model(&entity.Harvest{}).Select("quantity").Where("id = ? AND status = ?", input.HarvestId, 0).First(&quantity)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	if input.Quantity > quantity {
		tx.Rollback()
		return helper.ErrQuantityNotEnough
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

	if input.Desc != "" {
		updatedHarvest["desc"] = input.Desc
	}

	if regency.ID != 0 {
		updatedHarvest["regency_id"] = regency.ID
	}

	err2 := tx.Debug().Model(&entity.Harvest{}).Where("id = ? AND farmer_profile_id = ?", input.HarvestId, input.FarmerProfileId).Updates(&updatedHarvest)
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
	keyFYP := "farm:fyp"
	_, err := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keyFarmer)
		p.Del(ctx, keyHarvest)
		p.Del(ctx, keyFYP)
		keys, _ := r.redis.Keys(ctx, "farm:search:*").Result()
		if len(keys) > 0 {
			p.Del(ctx, keys...)
		}
		return nil
	})

	if err != nil {
		return err
	}

	return tx.Commit().Error
}

func (r *farmerRepository) DeleteHarvest(ctx context.Context, farmerProfileId, harvestId uint) error {
	res := r.db.WithContext(ctx).Debug().Model(&entity.Harvest{}).Where("id = ? AND farmer_profile_id = ? AND status = ?", harvestId, farmerProfileId, 0).Update("status", 2)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	keyFarmer := fmt.Sprintf("farm:farmer:%d", farmerProfileId)
	keyHarvest := fmt.Sprintf("farm:%d", harvestId)
	keyFYP := "farm:fyp"
	_, err := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keyFarmer)
		p.Del(ctx, keyHarvest)
		p.Del(ctx, keyFYP)
		keys, _ := r.redis.Keys(ctx, "farm:search:*").Result()
		if len(keys) > 0 {
			p.Del(ctx, keys...)
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *farmerRepository) UpdateStatusHarvest(ctx context.Context, farmerProfileId, harvestId uint) (*dto.BcHarvest, error) {
	tx := r.db.Begin().WithContext(ctx)

	var harvest entity.Harvest
	if res2 := tx.Debug().Model(&entity.Harvest{}).Where("id = ? AND farmer_profile_id = ?", harvestId, farmerProfileId).First(&harvest); res2.Error != nil {
		if errors.Is(res2.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, gorm.ErrRecordNotFound
		}
		tx.Rollback()
		return nil, res2.Error
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	keyFarmer := fmt.Sprintf("farm:farmer:%d", farmerProfileId)
	keyHarvest := fmt.Sprintf("farm:%d", harvestId)
	keyFYP := "farm:fyp"
	_, errRedis := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keyFarmer)
		p.Del(ctx, keyHarvest)
		p.Del(ctx, keyFYP)
		keys, _ := r.redis.Keys(ctx, "farm:search:*").Result()
		if len(keys) > 0 {
			p.Del(ctx, keys...)
		}
		return nil
	})

	if errRedis != nil {
		tx.Rollback()
		return nil, errRedis
	}

	return &dto.BcHarvest{
		ID:          int64(harvest.ID),
		FarmerID:    int64(harvest.FarmerProfileId),
		CropID:      int64(harvest.CropId),
		RegencyID:   int64(harvest.RegencyId),
		Name:        harvest.Name,
		Description: harvest.Desc,
		Quantity:    int64(harvest.Quantity),
		BasePrice:   int64(harvest.BasePrice),
	}, nil
}

// accept
func (r *farmerRepository) AcceptHarvestCollector(ctx context.Context, farmerId, collectorHarvestId uint) (*dto.BcHarvestCollector, error) {
	tx := r.db.Begin().WithContext(ctx)
	defer tx.Rollback()

	var v entity.HarvestCollector

	res := tx.Debug().Model(&entity.HarvestCollector{}).
		Preload("Harvest").
		Joins("JOIN harvests AS h ON h.id = harvest_collectors.harvest_id").
		Where("harvest_collectors.id = ? AND harvest_collectors.status = ? AND h.farmer_profile_id = ?", collectorHarvestId, 0, farmerId).
		First(&v)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, gorm.ErrRecordNotFound
		}
		tx.Rollback()
		return nil, res.Error
	}

	if v.Quantity > v.Harvest.Quantity {
		tx.Rollback()
		return nil, helper.ErrQuantityNotEnough
	}

	keyCollector := fmt.Sprintf("harvestCollector:collector:%d", v.CollectorProfileId)
	keyHarvestCollector := fmt.Sprintf("harvestCollector:%d", v.ID)
	keyFYP := "harvestCollector:fyp"
	_, err := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keyCollector)
		p.Del(ctx, keyHarvestCollector)
		p.Del(ctx, keyFYP)
		keys, _ := r.redis.Keys(ctx, "harvestCollector:search:*").Result()
		if len(keys) > 0 {
			p.Del(ctx, keys...)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	newQuantity := v.Harvest.Quantity - v.Quantity
	res3 := tx.Debug().Model(&entity.Harvest{}).Where("id = ? AND status = ?", v.HarvestId, 1).Update("quantity", newQuantity)
	if res3.Error != nil {
		return nil, res3.Error
	}

	if res3.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if err := tx.Commit().Error; err != nil {
		return nil, tx.Commit().Error
	}
	return &dto.BcHarvestCollector{
		ID:          int64(v.ID),
		CollectorID: int64(v.CollectorProfileId),
		HarvestID:   int64(v.HarvestId),
		Name:        v.Name,
		Desc:        v.Desc,
		Quantity:    int64(v.Quantity),
		Price:       int64(v.Quantity),
		BasePrice:   int64(v.BasePrice),
	}, nil
}

func (r *farmerRepository) AcceptHarvestProcessor(ctx context.Context, farmerId, processorHarvestId uint) (*dto.BcHarvestProcessor, error) {
	tx := r.db.Begin().WithContext(ctx)
	defer tx.Rollback()

	var v entity.HarvestProcessor

	res := tx.Debug().Model(&entity.HarvestProcessor{}).
		Preload("Harvest").
		Preload("HarvestCollector").
		Joins("JOIN harvests AS h ON h.id = harvest_processors.harvest_id").
		Where("harvest_processors.id = ? AND harvest_processors.status = ? AND h.farmer_profile_id = ?", processorHarvestId, 0, farmerId).
		First(&v)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, gorm.ErrRecordNotFound
		}
		tx.Rollback()
		return nil, res.Error
	}

	var newQuantity float64
	if v.HarvestCollector != nil {
		if v.Quantity > v.HarvestCollector.Quantity {
			tx.Rollback()
			return nil, helper.ErrQuantityNotEnough
		}

		newQuantity = v.HarvestCollector.Quantity - v.Quantity
		res3 := tx.Debug().Model(&entity.HarvestCollector{}).Where("id  = ? AND status = ?", v.HarvestCollectorId, 1).Update("quantity", newQuantity)
		if res3.Error != nil {
			return nil, res3.Error
		}

		if res3.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
		}
	} else {
		if v.Quantity > v.Harvest.Quantity {
			tx.Rollback()
			return nil, helper.ErrQuantityNotEnough
		}

		newQuantity = v.Harvest.Quantity - v.Quantity
		res3 := tx.Debug().Model(&entity.Harvest{}).Where("id = ? AND status = ?", v.HarvestId, 1).Update("quantity", newQuantity)
		if res3.Error != nil {
			return nil, res3.Error
		}
	}

	keyProcessor := fmt.Sprintf("harvestProcessor:processor:%d", v.ProcessorProfileId)
	keyHarvestProcessor := fmt.Sprintf("harvestProcessor:%d", v.ID)
	keyFYP := "harvestProcessor:fyp"
	_, err := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keyProcessor)
		p.Del(ctx, keyHarvestProcessor)
		p.Del(ctx, keyFYP)
		keys, _ := r.redis.Keys(ctx, "harvestProcessor:search:*").Result()
		if len(keys) > 0 {
			p.Del(ctx, keys...)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &dto.BcHarvestProcessor{
		ID:                 int64(v.ID),
		ProcessorID:        int64(v.ProcessorProfileId),
		HarvestCollectorID: int64(helper.Int64OrZero(v.HarvestCollectorId)),
		HarvestID:          int64(helper.Int64OrZero(v.HarvestId)),
		Name:               v.Name,
		Desc:               v.Desc,
		Quantity:           int64(v.Quantity),
		BasePrice:          int64(v.BasePrice),
		Price:              int64(v.Price),
	}, nil

}

func (r *farmerRepository) AcceptDistributor(ctx context.Context, farmerId, distributorHarvestId uint) (*dto.BcDistribution, error) {
	tx := r.db.Begin().WithContext(ctx)
	defer tx.Rollback()

	var v entity.Distribution

	res := tx.Debug().
		Model(&entity.Distribution{}).
		Joins("JOIN harvests h ON h.id = distributions.harvest_id").
		Where("distributions.id = ? AND distributions.status = ? AND h.farmer_profile_id = ?", distributorHarvestId, 0, farmerId).
		Preload("Harvest").
		Preload("HarvestCollector").
		Preload("HarvestProcessor").
		First(&v)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, gorm.ErrRecordNotFound
		}
		tx.Rollback()
		return nil, res.Error
	}

	var newQuantity float64
	if v.Harvest != nil {
		if v.Quantity > v.Harvest.Quantity {
			tx.Rollback()
			return nil, helper.ErrQuantityNotEnough
		}

		newQuantity = v.Harvest.Quantity - v.Quantity
		res3 := tx.Debug().Model(&entity.Harvest{}).Where("id = ? AND status = ?", v.Harvest.ID, 1).Update("quantity", newQuantity)
		if res3.Error != nil {
			return nil, res3.Error
		}

		if res3.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
		}
	} else if v.HarvestCollector != nil {
		if v.Quantity > v.HarvestCollector.Quantity {
			tx.Rollback()
			return nil, helper.ErrQuantityNotEnough
		}

		newQuantity = v.HarvestCollector.Quantity - v.Quantity
		res3 := tx.Debug().Model(&entity.HarvestCollector{}).Where("id = ? AND status = ?", v.HarvestCollector.ID, 1).Update("quantity", newQuantity)
		if res3.Error != nil {
			return nil, res3.Error
		}

		if res3.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
		}
	} else if v.HarvestProcessor != nil {
		if v.Quantity > v.HarvestProcessor.Quantity {
			tx.Rollback()
			return nil, helper.ErrQuantityNotEnough
		}

		newQuantity = v.HarvestProcessor.Quantity - v.Quantity
		res3 := tx.Debug().Model(&entity.HarvestProcessor{}).Where("id = ? AND status = ?", v.HarvestProcessor.ID, 1).Update("quantity", newQuantity)
		if res3.Error != nil {
			return nil, res3.Error
		}

		if res3.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
		}
	}

	keyDistributor := fmt.Sprintf("distribution:distributor:%d", v.DistributorProfileId)
	keyDistribution := fmt.Sprintf("distribution:%d", v.ID)
	keyFYP := "distribution:fyp"
	_, errRedis := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keyDistribution)
		p.Del(ctx, keyDistributor)
		p.Del(ctx, keyFYP)
		keys, _ := r.redis.Keys(ctx, "distribution:search:*").Result()
		if len(keys) > 0 {
			p.Del(ctx, keys...)
		}
		return nil
	})

	if errRedis != nil {
		return nil, errRedis
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &dto.BcDistribution{
		ID:               int64(v.ID),
		DistributorID:    int64(v.DistributorProfileId),
		DestinationID:    int64(v.DestinationId),
		HarvestID:        int64(helper.Int64OrZero(v.HarvestId)),
		HarvestCollector: int64(helper.Int64OrZero(v.HarvestCollectorId)),
		HarvestProcessor: int64(helper.Int64OrZero(v.HarvestCollectorId)),
		Name:             v.Name,
		Desc:             v.Desc,
		Quantity:         int64(v.Quantity),
		BasePrice:        int64(v.BasePrice),
		Price:            int64(v.Price),
		Transportation:   v.Transportation,
	}, nil
}

func (r *farmerRepository) UpdateBcBlockHarvest(ctx context.Context, harvestId uint, txBlock string) error {
	tx := r.db.Begin().WithContext(ctx)
	defer tx.Rollback()

	var harvest entity.Harvest
	res := tx.Debug().Model(&entity.Harvest{}).Select("id", "farmer_profile_id").Where("id = ?", harvestId).First(&harvest)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}
	res2 := tx.Debug().Model(&entity.Harvest{}).Where("id = ?", harvest.ID).Updates(map[string]interface{}{
		"tx_block": txBlock,
		"status":   1,
	})
	if res2.Error != nil {
		tx.Rollback()
		return res2.Error
	}

	if res2.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRegistered
	}

	keyFarmer := fmt.Sprintf("farm:farmer:%d", harvest.FarmerProfileId)
	keyHarvest := fmt.Sprintf("farm:%d", harvest.ID)
	keyFYP := "farm:fyp"
	_, errRedis := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keyFarmer)
		p.Del(ctx, keyHarvest)
		p.Del(ctx, keyFYP)
		keys, _ := r.redis.Keys(ctx, "farm:search:*").Result()
		if len(keys) > 0 {
			p.Del(ctx, keys...)
		}
		return nil
	})

	if errRedis != nil {
		tx.Rollback()
		return errRedis
	}

	return tx.Commit().Error
}

// get
func (r *farmerRepository) ListHarvestByFarmerId(ctx context.Context, farmerId uint) ([]dto.GetListHarvest, error) {

	key := fmt.Sprintf("farm:farmer:%d", farmerId)
	var harvests []dto.GetListHarvest

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &harvests); err == nil {
			return harvests, nil
		}
	}

	res := r.db.WithContext(ctx).Debug().Table("farmer_profiles as fp").
		Select(
			"cp.crop_name AS crop_name",
			"rg.name AS regency_name",
			"fp.name AS farmer_name",
			"h.id AS id",
			"h.name AS name",
			"h.status AS status",
			"h.quantity AS quantity",
			"h.updated_at AS time",
			"h.base_price AS base_price",
		).
		Joins("JOIN harvests AS h ON h.farmer_profile_id = fp.id").
		Joins("JOIN crops AS cp ON cp.id = h.crop_id").
		Joins("JOIN regencies AS rg ON rg.id = h.regency_id").
		Where("fp.id = ?", farmerId).
		Scan(&harvests)

	if res.Error != nil {
		return nil, res.Error
	}

	if len(harvests) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if harvests[0].Id == 0 {
		return nil, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(&harvests)
	r.redis.Set(ctx, key, jsonData, 5*time.Minute)
	return harvests, nil
}

func (r *farmerRepository) ListHarvestFYP(ctx context.Context) ([]dto.GetListHarvest, error) {
	var harvests []dto.GetListHarvest
	key := "farm:fyp"

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &harvests); err == nil {
			return harvests, nil
		}
	}

	res := r.db.WithContext(ctx).Debug().Model(&entity.Harvest{}).
		Select(
			"cp.crop_name AS crop_name",
			"fp.name AS farmer_name",
			"harvests.id AS id",
			"harvests.name AS name",
			"harvests.status AS status",
			"harvests.updated_at AS time",
			"harvests.base_price",
			"harvests.quantity AS quantity",
			"rg.name AS regency_name",
		).
		Joins("JOIN crops AS cp ON cp.id = harvests.crop_id").
		Joins("JOIN farmer_profiles AS fp ON fp.id = harvests.farmer_profile_id").
		Joins("JOIN regencies AS rg ON rg.id = harvests.regency_id").
		Order("RAND()").
		Scan(&harvests)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.Error != nil {
		return nil, res.Error
	}

	if len(harvests) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if harvests[0].Id == 0 {
		return nil, gorm.ErrEmptySlice
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
	res := r.db.WithContext(ctx).Debug().Model(&entity.Harvest{}).
		Select(
			"cp.crop_name AS crop_name",
			"cp.commodity AS commodity",
			"fp.name AS farmer_name",
			"harvests.id AS id",
			"harvests.name AS name",
			"harvests.status AS status",
			"harvests.`desc`AS description",
			"harvests.base_price AS base_price",
			"harvests.updated_at AS time",
			"harvests.quantity AS quantity",
			"harvests.tx_block AS tx_block",
			"rg.name AS regency_name",
			"rn.name AS region_name",
			"cn.name AS country_name",
		).
		Joins("JOIN crops AS cp ON cp.id = harvests.crop_id").
		Joins("JOIN farmer_profiles AS fp ON fp.id = harvests.farmer_profile_id").
		Joins("JOIN regencies AS rg ON rg.id = harvests.regency_id").
		Joins("JOIN regions AS rn ON rn.id = rg.region_id").
		Joins("JOIN countries AS cn ON cn.id = rn.country_id").
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
		if err := json.Unmarshal([]byte(val), &harvests); err == nil {
			return harvests, nil
		}
	}

	input := "%" + search + "%"
	res := r.db.WithContext(ctx).Debug().Model(&entity.Harvest{}).
		Select(
			"cp.crop_name AS crop_name",
			"fp.name AS farmer_name",
			"harvests.id AS id",
			"harvests.updated_at AS time",
			"harvests.name AS name",
			"harvests.status AS status",
			"harvests.quantity AS quantity",
			"harvests.base_price AS base_price",
			"rg.name AS regency_name",
		).
		Joins("JOIN crops AS cp ON cp.id = harvests.crop_id").
		Joins("JOIN farmer_profiles AS fp ON fp.id = harvests.farmer_profile_id").
		Joins("JOIN regencies AS rg ON rg.id = harvests.regency_id").
		Where("cp.crop_name LIKE ? OR fp.name LIKE ? OR cp.commodity LIKE  ? OR harvests.name LIKE ?", input, input, input, input).
		Scan(&harvests)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.Error != nil {
		return nil, res.Error
	}

	if len(harvests) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if harvests[0].Id == 0 {
		return nil, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(harvests)
	if err := r.redis.Set(ctx, key, jsonData, 5*time.Minute).Err(); err != nil {
		return nil, err
	}

	return harvests, nil
}
