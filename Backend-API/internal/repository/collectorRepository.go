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

type CollectorRepository interface {
	CreateHarvestCollector(ctx context.Context, input *dto.CreateCollector) error
	UpdateHarvestCollector(ctx context.Context, input *dto.UpdateCollector) error
	DeleteHarvestCollector(ctx context.Context, collectorId, collectorProfileId uint) error

	//accept
	AcceptHarvestProcessor(ctx context.Context, collectorProfileId, processorId uint) (*dto.BcHarvestProcessor, error)
	AcceptDistributor(ctx context.Context, collectorProfileId, distributorId uint) (*dto.BcDistribution, error)

	//blockchain
	UpdateBcBlockCollector(ctx context.Context, collectionId uint, txBlock string) error

	//get
	ListHarvestCollectorByCollectorId(ctx context.Context, collectorId uint) ([]dto.GetListHarvestCollector, error)
	ListHarvestCollectorFYP(ctx context.Context) ([]dto.GetListHarvestCollector, error)
	SearchHarvestCollector(ctx context.Context, search string) ([]dto.GetListHarvestCollector, error)
	GetHarvestCollectorById(ctx context.Context, id uint) (*dto.GetHarvestCollectorById, error)
}

type collectorRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewCollectorRepository(db *gorm.DB, redis *redis.Client) CollectorRepository {
	return &collectorRepository{db, redis}
}

func (r *collectorRepository) CreateHarvestCollector(ctx context.Context, input *dto.CreateCollector) error {
	tx := r.db.Begin().WithContext(ctx)
	defer tx.Rollback()

	var quantity float64
	res := tx.Debug().Model(&entity.Harvest{}).Select("quantity").Where("id = ? AND status = ?", input.HarvestId, 1).First(&quantity)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	if input.Quantity > quantity {
		tx.Rollback()
		return helper.ErrQuantityNotEnough
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	newCollector := entity.HarvestCollector{
		CollectorProfileId: input.CollectorProfileId,
		HarvestId:          input.HarvestId,
		Name:               input.Name,
		Desc:               input.Desc,
		Quantity:           input.Quantity,
		BasePrice:          input.BasePrice,
		Price:              input.Price,
		Status:             0,
	}

	res2 := tx.Debug().Model(&entity.HarvestCollector{}).Create(&newCollector)

	if res2.Error != nil {
		tx.Rollback()
		return res2.Error
	}

	if res2.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	keyCollector := fmt.Sprintf("harvestCollector:collector:%d", input.CollectorProfileId)
	keyHarvestCollector := fmt.Sprintf("harvestCollector:%d", newCollector.ID)
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
		return err
	}

	return tx.Commit().Error
}

func (r *collectorRepository) UpdateHarvestCollector(ctx context.Context, input *dto.UpdateCollector) error {
	tx := r.db.Begin().WithContext(ctx)
	defer tx.Rollback()

	var quantity float64
	update := make(map[string]interface{})

	if input.BasePrice > 0 {
		update["base_price"] = input.BasePrice
	}
	if input.Price > 0 {
		update["price"] = input.Price
	}

	if input.Name != "" {
		update["name"] = input.Name
	}
	if input.Desc != "" {
		update["desc"] = input.Desc
	}

	if input.Quantity > 0 {
		res := tx.Debug().Debug().Table("harvest_collectors as hc").
			Select("h.quantity").
			Joins("JOIN harvests AS h ON h.id  = hc.harvest_id").
			Where("hc.id = ? AND hc.status = ?", input.CollectorId, 0).Scan(&quantity)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		fmt.Println(quantity)
		if input.Quantity > quantity {
			tx.Rollback()
			return helper.ErrQuantityNotEnough
		}

		if res.RowsAffected == 0 {
			tx.Rollback()

			return gorm.ErrRecordNotFound
		}
		update["quantity"] = input.Quantity

	}

	res2 := tx.Debug().Model(&entity.HarvestCollector{}).Where("id = ? AND collector_profile_id = ? and status = ?", input.CollectorId, input.CollectorProfileId, 0).Updates(update)
	if res2.Error != nil {
		tx.Rollback()
		return res2.Error
	}

	if res2.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	keyCollector := fmt.Sprintf("harvestCollector:collector:%d", input.CollectorProfileId)
	keyHarvestCollector := fmt.Sprintf("harvestCollector:%d", input.CollectorId)
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
		return err
	}

	return tx.Commit().Error
}

func (r *collectorRepository) DeleteHarvestCollector(ctx context.Context, collectorProfileId, harvestCollectorId uint) error {
	res := r.db.WithContext(ctx).Debug().Model(&entity.HarvestCollector{}).Where("id = ? AND collector_profile_id = ? and status = ? ", harvestCollectorId, collectorProfileId, 0)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	keyCollector := fmt.Sprintf("harvestCollector:collector:%d", collectorProfileId)
	keyHarvestCollector := fmt.Sprintf("harvestCollector:%d", harvestCollectorId)
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
		return err
	}

	return nil

}

func (r *collectorRepository) AcceptHarvestProcessor(ctx context.Context, collectorId, processorHarvestId uint) (*dto.BcHarvestProcessor, error) {
	tx := r.db.Begin().WithContext(ctx)
	defer tx.Rollback()

	var v entity.HarvestProcessor

	res := tx.Debug().Model(&entity.HarvestProcessor{}).
		Preload("Harvest").
		Preload("HarvestCollector").
		Joins("JOIN harvest_collectors AS hc on hc.id  = harvest_processors.harvest_collector_id").
		Where("harvest_processors.id = ? AND harvest_processors.status = ? AND hc.collector_profile_id = ?", processorHarvestId, 0, collectorId).
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

func (r *collectorRepository) AcceptDistributor(ctx context.Context, collectorId, distributorHarvestId uint) (*dto.BcDistribution, error) {
	tx := r.db.Begin().WithContext(ctx)
	defer tx.Rollback()

	var v entity.Distribution

	res := tx.Debug().Model(&entity.Distribution{}).
		Preload("Harvest").
		Preload("HarvestCollector").
		Preload("HarvestProcessor").
		Joins("JOIN harvest_collectors AS hc ON hc.id = distributions.harvest_collector_id").
		Where("distributions.id = ? AND distributions.status = ? AND hc.collector_profile_id = ?", distributorHarvestId, 0, collectorId).
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

func (r *collectorRepository) UpdateBcBlockCollector(ctx context.Context, collectionId uint, txBlock string) error {
	tx := r.db.Begin().WithContext(ctx)
	defer tx.Rollback()

	var harvestCollector entity.HarvestCollector
	res := tx.Debug().Model(&entity.HarvestCollector{}).Select("id", "collector_profile_id").Where("id = ?", collectionId).First(&harvestCollector)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	res2 := tx.Debug().Model(&entity.HarvestCollector{}).Where("id = ?", harvestCollector.ID).Updates(map[string]interface{}{
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

	keyCollector := fmt.Sprintf("harvestCollector:collector:%d", harvestCollector.ID)
	keyHarvestCollector := fmt.Sprintf("harvestCollector:%d", harvestCollector.CollectorProfileId)
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
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
func (r *collectorRepository) ListHarvestCollectorByCollectorId(ctx context.Context, collectorId uint) ([]dto.GetListHarvestCollector, error) {
	var result []dto.GetListHarvestCollector

	key := fmt.Sprintf("harvestCollector:collector:%d", collectorId)

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err != nil {
			return result, nil
		}
	}
	res := r.db.WithContext(ctx).Debug().Table("collector_profiles AS cp").
		Select(
			"hc.id AS id",
			"hc.quantity AS quantity",
			"hc.base_price AS base_price",
			"hc.name as name",
			"hc.status AS status",
			"hc.price AS price",
			"hc.updated_at AS time",
			"c.crop_name AS crop_name",
			"cp.name AS collector_profile_name",
		).
		Joins("LEFT JOIN harvest_collectors AS hc ON hc_collector_profile_id = cp.id").
		Joins("JOIN harvests AS h ON h.id = hc.harvest_id").
		Joins("JOIN crops AS c ON c.id = h.crop_id").
		Where("cp.id = ?", collectorId).
		Scan(&result)

	if res.Error != nil {
		return nil, res.Error
	}

	if len(result) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if result[0].CollectorId == 0 {
		return nil, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(result)
	if err := r.redis.Set(ctx, key, jsonData, 5*time.Minute).Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *collectorRepository) ListHarvestCollectorFYP(ctx context.Context) ([]dto.GetListHarvestCollector, error) {
	var result []dto.GetListHarvestCollector

	key := "harvestCollector:fyp"

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err != nil {
			return result, nil
		}
	}
	res := r.db.WithContext(ctx).Debug().Table("harvest_collectors AS hc").
		Select("hc.id AS id",
			"hc.quantity AS quantity",
			"hc.name as name",
			"hc.status AS status",
			"hc.base_price AS base_price",
			"hc.price AS price",
			"hc.updated_at AS time",
			"c.crop_name AS crop_name",
			"cp.name AS collector_profile_name",
		).
		Joins("JOIN collector_profiles AS cp ON cp.id = hc.collector_profile_id").
		Joins("JOIN harvests AS h ON h.id = hc.harvest_id").
		Joins("JOIN crops AS c ON c.id = h.crop_id").
		Order("RAND()").
		Scan(&result)

	if res.Error != nil {
		return nil, res.Error
	}

	if len(result) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if result[0].CollectorId == 0 {
		return nil, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(result)
	if err := r.redis.Set(ctx, key, jsonData, 5*time.Minute).Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *collectorRepository) SearchHarvestCollector(ctx context.Context, search string) ([]dto.GetListHarvestCollector, error) {
	input := "%" + search + "%"
	var result []dto.GetListHarvestCollector

	key := fmt.Sprintf("harvestCollector:search:%s", search)

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err != nil {
			return result, nil
		}
	}

	res := r.db.WithContext(ctx).Debug().Table("harvest_collectors AS hc").
		Select("hc.id AS id",
			"hc.quantity AS quantity",
			"hc.name as name",
			"hc.status AS status",
			"hc.base_price AS base_price",
			"hc.price AS price",
			"hc.updated_at AS time",
			"c.crop_name AS crop_name",
			"cp.name AS collector_profile_name",
		).
		Joins("JOIN collector_profiles AS cp ON cp.id = hc.collector_profile_id").
		Joins("JOIN harvests AS h ON h.id = hc.harvest_id").
		Joins("JOIN crops AS c ON c.id = h.crop_id").
		Where("c.crop_name LIKE ? OR cp.name LIKE ? OR  hc.name LIKE ?", input, input, input).
		Scan(&result)

	if res.Error != nil {
		return nil, res.Error
	}

	if len(result) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if result[0].CollectorId == 0 {
		return nil, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(result)
	if err := r.redis.Set(ctx, key, jsonData, 5*time.Minute).Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *collectorRepository) GetHarvestCollectorById(ctx context.Context, id uint) (*dto.GetHarvestCollectorById, error) {
	var result dto.GetHarvestCollectorById

	key := fmt.Sprintf("harvestCollector:%d", id)

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err != nil {
			return &result, nil
		}
	}

	res := r.db.WithContext(ctx).Debug().Table("harvest_collectors AS hc").
		Select("hc.id AS id",
			"hc.quantity AS quantity",
			"hc.name as name",
			"hc.`desc`AS description",
			"hc.status AS status",
			"hc.base_price AS base_price",
			"hc.price AS price",
			"hc.tx_block AS tx_block",
			"hc.updated_at AS time",
			"c.crop_name AS crop_name",
			"cp.name AS collector_profile_name",
		).
		Joins("JOIN collector_profiles AS cp ON cp.id = hc.collector_profile_id").
		Joins("JOIN harvests AS h ON h.id = hc.harvest_id").
		Joins("JOIN crops AS c ON c.id = h.crop_id").
		Where("hc.id = ?", id).
		Scan(&result)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	jsonData, _ := json.Marshal(result)
	if err := r.redis.Set(ctx, key, jsonData, 5*time.Minute).Err(); err != nil {
		return nil, err
	}

	return &result, nil
}
