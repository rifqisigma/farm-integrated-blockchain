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

type ProcessorRepository interface {
	CreateProcessor(ctx context.Context, input *dto.CreateProcessor) error
	UpdateProcessor(ctx context.Context, input *dto.UpdateProcessor) error
	DeleteProcessor(ctx context.Context, procesorId, processoProfileId uint) error

	//accept
	AcceptDistributor(ctx context.Context, processorId, distributorHarvestId uint) (*dto.BcDistribution, error)

	//blockchain
	UpdateBcBlockProcessor(ctx context.Context, processorId uint, txBlock string) error
	//get
	ListHarvestProcessorByProcessorId(ctx context.Context, processorId uint) ([]dto.GetListHarvestProcessor, error)
	SearchHarvestProcessor(ctx context.Context, search string) ([]dto.GetListHarvestProcessor, error)
	ListHarvestProcessorFYP(ctx context.Context) ([]dto.GetListHarvestProcessor, error)
	GetHarvestProcessorById(ctx context.Context, id uint) (*dto.GetHarvestProcessorById, error)
}

type processorRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewProcessorRepository(db *gorm.DB, redis *redis.Client) ProcessorRepository {
	return &processorRepository{db, redis}
}

// w
func (r *processorRepository) CreateProcessor(ctx context.Context, input *dto.CreateProcessor) error {
	tx := r.db.Begin().WithContext(ctx)
	defer tx.Rollback()

	var quantity float64

	newProcessor := entity.HarvestProcessor{
		ProcessorProfileId: input.ProcessorProfileId,
		Name:               input.Name,
		Desc:               input.Desc,
		Quantity:           input.Quantity,
		BasePrice:          input.BasePrice,
		Price:              input.Price,
		Status:             0,
	}

	var unkwonEntity interface{}
	var idQuantity uint
	if input.HarvestId != 0 {
		idQuantity = input.HarvestId
		unkwonEntity = entity.Harvest{}
		newProcessor.HarvestId = &input.HarvestId

	} else if input.HarvestCollectorId != 0 {
		idQuantity = input.HarvestCollectorId
		unkwonEntity = entity.HarvestCollector{}
		newProcessor.HarvestCollectorId = &input.HarvestCollectorId
	}
	res := tx.Debug().Model(&unkwonEntity).Select("quantity").Where("id = ? AND status = ?", idQuantity, 1).First(&quantity)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return gorm.ErrRecordNotFound
		}
		tx.Rollback()
		return res.Error
	}

	if input.Quantity > quantity {
		tx.Rollback()
		return helper.ErrQuantityNotEnough
	}

	res2 := tx.Debug().Model(&entity.HarvestProcessor{}).Create(&newProcessor)

	if res2.Error != nil {
		tx.Rollback()
		return res2.Error
	}

	if res2.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	keyProcessor := fmt.Sprintf("harvestProcessor:processor:%d", input.ProcessorProfileId)
	keyHarvestProcessor := fmt.Sprintf("harvestProcessor:%d", newProcessor.ID)
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
		return err
	}

	return tx.Commit().Error

}

func (r *processorRepository) UpdateProcessor(ctx context.Context, input *dto.UpdateProcessor) error {
	tx := r.db.Begin().WithContext(ctx)
	defer tx.Rollback()

	update := make(map[string]interface{})

	if input.Quantity != 0 {
		var result struct {
			Quantity float64
		}

		res := tx.Debug().Table("harvest_processors as hp").
			Select("COALESCE(h1.quantity, h2.quantity) AS quantity").
			Joins("LEFT JOIN harvest_collectors AS hc ON hc.id =  hp.harvest_collector_id").
			Joins("LEFT JOIN harvests AS h1 ON h1.id = hp.harvest_id").
			Joins("LEFT JOIN harvests AS h2 ON h2.id = hc.harvest_id").
			Where("hp.id = ? AND hp.status = ? AND h1.status = ? OR h2.status = ?", input.ProcessorHarvestId, 0, 1, 1).Scan(&result)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		fmt.Println(result.Quantity)

		if input.Quantity > result.Quantity {
			tx.Rollback()
			return helper.ErrQuantityNotEnough
		}

	}

	if input.BasePrice != 0 {
		update["base_price"] = input.BasePrice
	}

	if input.Price != 0 {
		update["price"] = input.Price
	}

	if input.Name != "" {
		update["name"] = input.Name
	}

	if input.Desc != "" {
		update["desc"] = input.Desc
	}

	res2 := tx.Debug().Model(&entity.HarvestProcessor{}).Where("id = ? AND processor_profile_id = ? AND status  = ?", input.ProcessorHarvestId, input.ProcessorProfileId, 0).Updates(&update)

	if res2.Error != nil {
		tx.Rollback()
		return res2.Error
	}

	if res2.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	keyProcessor := fmt.Sprintf("harvestProcessor:processor:%d", input.ProcessorProfileId)
	keyHarvestProcessor := fmt.Sprintf("harvestProcessor:%d", input.ProcessorHarvestId)
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
		return err
	}

	return tx.Commit().Error
}

func (r *processorRepository) DeleteProcessor(ctx context.Context, procesorId, processoProfileId uint) error {
	tx := r.db.Begin().WithContext(ctx)
	defer tx.Rollback()

	res := tx.Debug().Model(&entity.HarvestProcessor{}).Where("id = ? AND processor_profile_id = ? AND status = ?", procesorId, processoProfileId, 0).Update("status", 2)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	if res.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	keyProcessor := fmt.Sprintf("harvestProcessor:processor:%d", processoProfileId)
	keyHarvestProcessor := fmt.Sprintf("harvestProcessor:%d", procesorId)
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
		return err
	}

	return tx.Commit().Error
}

// accept
func (r *processorRepository) AcceptDistributor(ctx context.Context, processorId, distributorHarvestId uint) (*dto.BcDistribution, error) {
	tx := r.db.Begin().WithContext(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var v entity.Distribution

	res := tx.Debug().Model(&entity.Distribution{}).
		Preload("Harvest").
		Preload("HarvestCollector").
		Preload("HarvestProcessor").
		Joins("JOIN harvest_processors AS hp ON hp.id = distributions.harvest_processor_id").
		Where("distributions.id = ? AND distributions.status = ? AND hp.processor_profile_id = ?", distributorHarvestId, 0, processorId).
		First(&v)
	if res.Error != nil {
		tx.Rollback()
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, res.Error
	}

	var newQuantity float64

	if v.Harvest != nil {
		if v.Quantity > v.Harvest.Quantity {
			tx.Rollback()
			return nil, helper.ErrQuantityNotEnough
		}

		newQuantity = v.Harvest.Quantity - v.Quantity
		res3 := tx.Debug().Model(&entity.Harvest{}).
			Where("id = ? AND status = ?", v.Harvest.ID, 1).
			Update("quantity", newQuantity)
		if res3.Error != nil {
			tx.Rollback()
			return nil, res3.Error
		}
		if res3.RowsAffected == 0 {
			tx.Rollback()
			return nil, gorm.ErrRecordNotFound
		}

	} else if v.HarvestCollector != nil {
		if v.Quantity > v.HarvestCollector.Quantity {
			tx.Rollback()
			return nil, helper.ErrQuantityNotEnough
		}

		newQuantity = v.HarvestCollector.Quantity - v.Quantity
		res3 := tx.Debug().Model(&entity.HarvestCollector{}).
			Where("id = ? AND status = ?", v.HarvestCollector.ID, 1).
			Update("quantity", newQuantity)
		if res3.Error != nil {
			tx.Rollback()
			return nil, res3.Error
		}
		if res3.RowsAffected == 0 {
			tx.Rollback()
			return nil, gorm.ErrRecordNotFound
		}

	} else if v.HarvestProcessor != nil {
		if v.Quantity > v.HarvestProcessor.Quantity {
			tx.Rollback()
			return nil, helper.ErrQuantityNotEnough
		}

		newQuantity = v.HarvestProcessor.Quantity - v.Quantity
		res3 := tx.Debug().Model(&entity.HarvestProcessor{}).
			Where("id = ? AND status = ?", v.HarvestProcessor.ID, 1).
			Update("quantity", newQuantity)
		if res3.Error != nil {
			tx.Rollback()
			return nil, res3.Error
		}
		if res3.RowsAffected == 0 {
			tx.Rollback()
			return nil, gorm.ErrRecordNotFound
		}

	} else {
		tx.Rollback()
		return nil, errors.New("invalid distribution source")
	}

	keyDistributor := fmt.Sprintf("distribution:distributor:%d", v.DistributorProfileId)
	keyDistribution := fmt.Sprintf("distribution:%d", v.ID)
	keyFYP := "distribution:fyp"

	keys, _ := r.redis.Keys(ctx, "distribution:search:*").Result()
	_, errRedis := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keyDistribution)
		p.Del(ctx, keyDistributor)
		p.Del(ctx, keyFYP)
		if len(keys) > 0 {
			p.Del(ctx, keys...)
		}
		return nil
	})
	if errRedis != nil {
		tx.Rollback()
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

func (r *processorRepository) UpdateBcBlockProcessor(ctx context.Context, processorId uint, txBlock string) error {
	tx := r.db.Begin().WithContext(ctx)
	defer tx.Rollback()

	var harvestProcessor entity.HarvestProcessor
	res := tx.Debug().Model(&entity.HarvestProcessor{}).Select("id", "processor_profile_id").Where("id = ?", processorId).First(&harvestProcessor)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	res2 := tx.Debug().Model(&entity.HarvestProcessor{}).Where("id = ?", harvestProcessor.ID).Updates(map[string]interface{}{
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

	keyProcessor := fmt.Sprintf("harvestProcessor:processor:%d", harvestProcessor.ProcessorProfileId)
	keyHarvestProcessor := fmt.Sprintf("harvestProcessor:%d", harvestProcessor.ID)
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
		return err
	}

	return tx.Commit().Error
}

// get
func (r *processorRepository) ListHarvestProcessorByProcessorId(ctx context.Context, processorId uint) ([]dto.GetListHarvestProcessor, error) {
	var result []dto.GetListHarvestProcessor

	key := fmt.Sprintf("harvestProcessor:processor:%d", processorId)

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err == nil {
			return result, nil
		}
	}
	res := r.db.WithContext(ctx).Debug().Table(" processor_profiles AS pp").
		Select(
			"hp.name AS name",
			"pp.name  AS processor_profile_name",
			"COALESCE(c1.crop_name, c2.crop_name) AS crop_name",
			"hp.id AS id",
			"hp.quantity AS quantity",
			"hp.base_price AS base_price",
			"hp.price AS price",
			"hp.status AS status",
			"hp.updated_at AS time",
		).
		Joins("LEFT JOIN harvest_processors AS hp ON hp.processor_profile_id = pp.id").
		Joins("LEFT JOIN harvest_collectors AS hc ON hc.id = hp.harvest_collector_id").
		Joins("LEFT JOIN harvests AS h1 ON h1.id = hp.harvest_id").
		Joins("LEFT JOIN harvests AS h2 ON h2.id = hc.harvest_id").
		Joins("LEFT JOIN crops AS c1 ON c1.id = h1.crop_id").
		Joins("LEFT JOIN crops AS c2 ON c2.id = h2.crop_id").
		Where("hp.processor_profile_id = ?", processorId).
		Scan(&result)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.Error != nil {
		return nil, res.Error
	}

	if len(result) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if result[0].ProcessorHarvestId == 0 {
		return nil, gorm.ErrEmptySlice
	}
	jsonData, _ := json.Marshal(result)
	r.redis.Set(ctx, key, jsonData, 5*time.Minute)

	return result, nil
}

func (r *processorRepository) ListHarvestProcessorFYP(ctx context.Context) ([]dto.GetListHarvestProcessor, error) {
	var result []dto.GetListHarvestProcessor

	key := "harvestProcessor:fyp"

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err == nil {
			return result, nil
		}
	}

	res := r.db.WithContext(ctx).Debug().Table("harvest_processors AS hp").
		Select(
			"hp.name AS name",
			"pp.name  AS processor_profile_name",
			"COALESCE(c1.crop_name, c2.crop_name) AS crop_name",
			"hp.id AS id",
			"hp.quantity AS quantity",
			"hp.base_price AS base_price",
			"hp.price AS price",
			"hp.status AS status",
			"hp.updated_at AS time",
		).
		Joins("JOIN processor_profiles AS pp ON pp.id = hp.processor_profile_id").
		Joins("LEFT JOIN harvest_collectors AS hc ON hc.id = hp.harvest_collector_id").
		Joins("LEFT JOIN harvests AS h1 ON h1.id = hp.harvest_id"). // langsung
		Joins("LEFT JOIN harvests AS h2 ON h2.id = hc.harvest_id"). // lewat collector
		Joins("LEFT JOIN crops AS c1 ON c1.id = h1.crop_id").
		Joins("LEFT JOIN crops AS c2 ON c2.id = h2.crop_id").
		Order("RAND()").
		Scan(&result)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.Error != nil {
		return nil, res.Error
	}

	if len(result) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if result[0].ProcessorHarvestId == 0 {
		return nil, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(result)
	r.redis.Set(ctx, key, jsonData, 5*time.Minute)

	return result, nil
}

func (r *processorRepository) SearchHarvestProcessor(ctx context.Context, search string) ([]dto.GetListHarvestProcessor, error) {
	input := "%" + search + "%"
	var result []dto.GetListHarvestProcessor

	key := fmt.Sprintf("harvestProcessor:search:%s", search)

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err == nil {
			return result, nil
		}
	}
	res := r.db.WithContext(ctx).Debug().Table("harvest_processors AS hp").
		Select(
			"hp.name AS name",
			"pp.name  AS processor_profile_name",
			"COALESCE(c1.crop_name, c2.crop_name) AS crop_name",
			"hp.id AS id",
			"hp.quantity AS quantity",
			"hp.base_price AS base_price",
			"hp.price AS price",
			"hp.status AS status",
			"hp.updated_at AS time").
		Joins("JOIN processor_profiles AS pp ON pp.id = hp.processor_profile_id").
		Joins("LEFT JOIN harvest_collectors AS hc ON hc.id = hp.harvest_collector_id").
		Joins("LEFT JOIN harvests AS h1 ON h1.id = hp.harvest_id"). // langsung
		Joins("LEFT JOIN harvests AS h2 ON h2.id = hc.harvest_id"). // lewat collector
		Joins("LEFT JOIN crops AS c1 ON c1.id = h1.crop_id").
		Joins("LEFT JOIN crops AS c2 ON c2.id = h2.crop_id").
		Where("hp.name LIKE ? OR hp.processor_profile_id LIKE ? OR c1.crop_name LIKE ? OR c2.crop_name LIKE ?", input, input, input, input).
		Scan(&result)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.Error != nil {
		return nil, res.Error
	}

	if len(result) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if result[0].ProcessorHarvestId == 0 {
		return nil, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(result)
	r.redis.Set(ctx, key, jsonData, 5*time.Minute)

	return result, nil
}

func (r *processorRepository) GetHarvestProcessorById(ctx context.Context, id uint) (*dto.GetHarvestProcessorById, error) {
	var result dto.GetHarvestProcessorById

	key := fmt.Sprintf("harvestProcessor:%d", id)

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err == nil {
			return &result, nil
		}
	}

	res := r.db.WithContext(ctx).Debug().Table("harvest_processors AS hp").
		Select(
			"hp.name AS name",
			"hp.`desc`AS description",
			"pp.name  AS processor_profile_name",
			"COALESCE(c1.crop_name, c2.crop_name) AS crop_name",
			"hp.id AS id",
			"hp.quantity AS quantity",
			"hp.base_price AS base_price",
			"hp.price AS price",
			"hp.status AS status",
			"hp.tx_block AS tx_block",
			"hp.updated_at AS time",
		).
		Joins("JOIN processor_profiles AS pp ON pp.id = hp.processor_profile_id").
		Joins("LEFT JOIN harvest_collectors AS hc ON hc.id = hp.harvest_collector_id").
		Joins("LEFT JOIN harvests AS h1 ON h1.id = hp.harvest_id"). // langsung
		Joins("LEFT JOIN harvests AS h2 ON h2.id = hc.harvest_id"). // lewat collector
		Joins("LEFT JOIN crops AS c1 ON c1.id = h1.crop_id").
		Joins("LEFT JOIN crops AS c2 ON c2.id = h2.crop_id").
		Order("RAND()").
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
