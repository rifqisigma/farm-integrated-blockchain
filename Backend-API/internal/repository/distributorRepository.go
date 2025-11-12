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

type DistributorRepository interface {
	CreateDistribution(ctx context.Context, input *dto.CreateDistribution) error
	UpdateDistribution(ctx context.Context, input *dto.UpdateDistribution) error
	DeleteDistribution(ctx context.Context, distrbutionId, distributorId uint) error
	UpdateStatusOfDistribution(ctx context.Context, input *dto.UpdateStatusDistribution) error

	//accept
	AcceptSeller(ctx context.Context, distributorProfileId, sellerBoxId uint) (*dto.BcSellerBox, error)

	//blockchain
	UpdateBcBlockDistribution(ctx context.Context, distrbutionId uint, txBlock string) error

	//get
	SearchDistributions(ctx context.Context, search string) ([]dto.GetListDistribution, error)
	GetListDistributionsByDistributorId(ctx context.Context, id uint) ([]dto.GetListDistribution, error)
	GetDistributionById(ctx context.Context, id uint) (*dto.GetDistributionById, error)
	GetListDistributionFYP(ctx context.Context) ([]dto.GetListDistribution, error)
}

type distributorRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewDistributorRepository(db *gorm.DB, redis *redis.Client) DistributorRepository {
	return &distributorRepository{db, redis}
}

func (r *distributorRepository) CreateDistribution(ctx context.Context, input *dto.CreateDistribution) error {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var regency entity.Regency
	res := tx.WithContext(ctx).Where("name = ?", input.CountryRequest.RegionRequest.RegencyRequest.Name).First(&regency)
	if res.Error != nil {
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
			if err := tx.WithContext(ctx).Create(&regency).Error; err != nil {
				return err
			}
		}
		return res.Error
	}

	newDistribution := entity.Distribution{
		Name:                 input.Name,
		Desc:                 input.Desc,
		Transportation:       input.Transportation,
		Quantity:             input.Quantity,
		Price:                input.Price,
		BasePrice:            input.BasePrice,
		DistributorProfileId: input.DistributorProfileId,
		DestinationId:        regency.ID,
	}

	var unkwonEntity interface{}
	var id uint
	if input.HarvestId != 0 {
		unkwonEntity = entity.Harvest{}
		newDistribution.HarvestId = &input.HarvestId
		id = input.HarvestId
	} else if input.HarvestCollectorId != 0 {
		unkwonEntity = entity.HarvestCollector{}
		newDistribution.HarvestCollectorId = &input.HarvestCollectorId
		id = input.HarvestCollectorId
	} else if input.HarvestProcessorId != 0 {
		unkwonEntity = entity.HarvestProcessor{}
		newDistribution.HarvestProcessorId = &input.HarvestProcessorId
		id = input.HarvestProcessorId
	} else {
		return gorm.ErrRecordNotFound
	}

	var quantity float64
	res2 := tx.Debug().Model(&unkwonEntity).Select("quantity").Where("id = ? AND status = ?", id, 1).First(&quantity)
	if res2.Error != nil {
		tx.Rollback()
		return res2.Error
	}
	if res2.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	if input.Quantity > quantity {
		tx.Rollback()
		return helper.ErrQuantityNotEnough
	}

	if err := tx.Debug().Model(&entity.Distribution{}).Create(&newDistribution).Error; err != nil {
		tx.Rollback()
	}

	keyDistributor := fmt.Sprintf("distribution:distributor:%d", newDistribution.DistributorProfileId)
	keyDistribution := fmt.Sprintf("distribution:%d", newDistribution.ID)
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
		return errRedis
	}

	return tx.Commit().Error
}

func (r *distributorRepository) UpdateDistribution(ctx context.Context, input *dto.UpdateDistribution) error {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	updateDistribution := make(map[string]interface{})

	if input.Quantity > 0 {
		var quantity float64
		res := tx.Debug().Table("distributions AS d").
			Joins("LEFT JOIN harvest_processors AS hp ON  hp.id  =  d.harvest_processor_id").
			Joins("LEFT JOIN harvest_collectors AS hc ON  hc.id  =  d.harvest_collector_id").
			Joins("LEFT JOIN harvests AS h1 ON  h1.id  =  d.harvest_id").
			Joins("LEFT JOIN harvests AS h2 ON  h2.id  =  hp.harvest_id").
			Joins("LEFT JOIN harvests AS h3  ON  h3.id  =  hc.harvest_id").
			Select("COALESCE(h1.quantity, h2.quantity, h3.quantity) AS quantity").
			Where("d.id = ? AND d.status = ?", input.DistributionId, 1).Scan(&quantity)
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

		updateDistribution["quantity"] = input.Quantity
	}

	if input.Name != "" {
		updateDistribution["name"] = input.Name
	}
	if input.Desc != "" {
		updateDistribution["desc"] = input.Desc
	}
	if input.Transportation != "" {
		updateDistribution["transportation"] = input.Transportation
	}
	if input.Price != 0 {
		updateDistribution["price"] = input.Price
	}
	if input.BasePrice != 0 {
		updateDistribution["base_price"] = input.BasePrice
	}

	if input.CountryRequest.RegionRequest.RegencyRequest.Name != "" {
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
				FirstOrCreate(&region, entity.Region{Name: input.CountryRequest.RegionRequest.Name, CountryId: country.ID}).Error; err != nil {
				return err
			}

			regency = entity.Regency{Name: input.CountryRequest.RegionRequest.RegencyRequest.Name, RegionId: region.ID}
			if err := tx.WithContext(ctx).Create(&regency).Error; err != nil {
				return err
			}
		}

		if input.CountryRequest.RegionRequest.RegencyRequest.Name != "" {
			updateDistribution["destination_id"] = regency.ID
		}
	}

	res2 := tx.Debug().Model(&entity.Distribution{}).Where("id = ? AND distributor_profile_id = ? AND status = ?", input.DistributionId, input.DistributorProfileId, 0).Updates(&updateDistribution)
	if res2.Error != nil {
		tx.Rollback()
		return res2.Error
	}

	if res2.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	keyDistributor := fmt.Sprintf("distribution:distributor:%d", input.DistributorProfileId)
	keyDistribution := fmt.Sprintf("distribution:%d", input.DistributionId)
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
		return errRedis
	}

	return tx.Commit().Error
}

func (r *distributorRepository) DeleteDistribution(ctx context.Context, distrbutionId, distributorId uint) error {
	res2 := r.db.Debug().Model(&entity.Distribution{}).Where("id = ? AND distributor_profile_id = ? AND status = ?", distrbutionId, distributorId, 0).Update("status", 2)
	if res2.Error != nil {
		return res2.Error
	}

	if res2.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	keyDistributor := fmt.Sprintf("distribution:distributor:%d", distributorId)
	keyDistribution := fmt.Sprintf("distribution:%d", distrbutionId)
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
		return errRedis
	}

	return nil

}

func (r *distributorRepository) UpdateStatusOfDistribution(ctx context.Context, input *dto.UpdateStatusDistribution) error {
	tx := r.db.Begin().WithContext(ctx)
	defer tx.Rollback()
	res := tx.Debug().Model(&entity.Distribution{}).Select("id", "distributor_profile_id").Where("id = ? AND distributor_profile_id = ? AND status = ?", input.DistributionId, input.DistributorProfileId, 1).First(&entity.Distribution{})
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	if res.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	err := tx.Debug().Model(&entity.Distribution{}).Where("id = ? AND distributor_profile_id = ? AND status = ?", input.DistributionId, input.DistributorProfileId, 1).Update("distribution_status", input.Status)
	if err.Error != nil {
		tx.Rollback()
		return err.Error
	}
	if err.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}
	return tx.Commit().Error

}

func (r *distributorRepository) AcceptSeller(ctx context.Context, distributorProfileId, sellerBoxId uint) (*dto.BcSellerBox, error) {
	tx := r.db.Begin().WithContext(ctx)
	defer tx.Rollback()

	var v entity.SellerBox

	res := tx.Debug().Model(&entity.SellerBox{}).
		Preload("Distribution").
		Joins("JOIN distributions AS d ON d.id = seller_boxes.distribution_id").
		Where("seller_boxes.id = ? AND d.status = ? AND d.distributor_profile_id = ? AND d.distribution_status  =  ?", sellerBoxId, 1, distributorProfileId, 7).
		First(&v)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, gorm.ErrRecordNotFound
		}
		tx.Rollback()
		return nil, res.Error
	}

	if v.Quantity > v.Distribution.Quantity {
		tx.Rollback()
		return nil, helper.ErrQuantityNotEnough
	}

	newQuantity := v.Distribution.Quantity - v.Quantity
	res3 := tx.Debug().Model(&entity.Distribution{}).Where("id = ? AND status = ?", v.Distribution.ID, 1).Update("quantity", newQuantity)
	if res3.Error != nil {
		tx.Rollback()
		return nil, res3.Error
	}

	if res3.RowsAffected == 0 {
		tx.Rollback()
		return nil, gorm.ErrRecordNotFound
	}

	keySeller := fmt.Sprintf("sellerBox:seller:%d", v.SellerProfileId)
	keySellerBox := fmt.Sprintf("sellerBox:%d", v.ID)
	keyFYP := "sellerBox:fyp"
	_, err := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keySeller)
		p.Del(ctx, keySellerBox)
		p.Del(ctx, keyFYP)
		keys, _ := r.redis.Keys(ctx, "sellerBox:search:*").Result()
		if len(keys) > 0 {
			p.Del(ctx, keys...)
		}
		return nil
	})

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &dto.BcSellerBox{
		ID:             int64(v.ID),
		SellerID:       int64(v.SellerProfileId),
		DistributionID: int64(v.DistributionId),
		Name:           v.Name,
		Desc:           v.Desc,
		Quantity:       int64(v.Quantity),
		BasePrice:      int64(v.BasePrice),
		Price:          int64(v.Price),
	}, nil
}

func (r *distributorRepository) UpdateBcBlockDistribution(ctx context.Context, distrbutionId uint, txBlock string) error {
	tx := r.db.Begin().WithContext(ctx)
	defer tx.Rollback()

	var distribution entity.Distribution
	res := tx.Debug().Model(&entity.Distribution{}).Select("id", "distributor_profile_id").Where("id = ?", distrbutionId).First(&distribution)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	if res.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	res2 := tx.Debug().Model(&entity.Distribution{}).Where("id = ?", distribution.ID).Updates(map[string]interface{}{
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

	keyDistributor := fmt.Sprintf("distribution:distributor:%d", distribution.DistributorProfileId)
	keyDistribution := fmt.Sprintf("distribution:%d", distribution.ID)
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
		return errRedis
	}

	return tx.Commit().Error
}

func (r *distributorRepository) SearchDistributions(ctx context.Context, search string) ([]dto.GetListDistribution, error) {
	var result []dto.GetListDistribution
	input := "%" + search + "%"
	key := fmt.Sprintf("distribution:search:%s", search)

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err == nil {
			return result, nil
		}
	}

	res := r.db.WithContext(ctx).Debug().Table("distributions AS d").
		Select(
			"d.id",
			"dp.name AS distributor_name",
			"COALESCE(c1.crop_name, c2.crop_name, c3.crop_name) AS crop_name",
			"rg.name AS destination",
			"d.name AS name",
			"d.price AS price",
			"d.base_price AS base_price",
			"d.quantity AS quantity",
			"d.transportation AS transportation",
			"d.status AS status",
			"rg.name AS regency_name",
			"d.updated_at AS time",
		).
		Joins("JOIN regencies AS rg ON rg.id = d.destination_id").
		Joins("JOIN distributor_profiles AS dp ON dp.id = d.distributor_profile_id").
		Joins("LEFT JOIN harvest_processors AS hp ON  hp.id  =  d.harvest_processor_id").
		Joins("LEFT JOIN harvest_collectors AS hc ON  hc.id  =  d.harvest_collector_id").
		Joins("LEFT JOIN harvests AS h1 ON  h1.id  =  d.harvest_id").
		Joins("LEFT JOIN harvests AS h2 ON  h2.id  =  hp.harvest_id").
		Joins("LEFT JOIN harvests AS h3  ON  h3.id  =  hc.harvest_id").
		Joins("LEFT JOIN crops AS c1  ON  c1.id  =  h1.crop_id").
		Joins("LEFT JOIN crops AS c2  ON  c2.id  =  h2.crop_id").
		Joins("LEFT JOIN crops AS c3  ON  c3.id  =  h3.crop_id").
		Where("c1.crop_name LIKE ? OR  c2.crop_name LIKE ? OR  c2.crop_name LIKE ? OR dp.name LIKE ? OR d.name LIKE ?", input, input, input, input, input).
		Scan(&result)
	if res.Error != nil {
		return nil, res.Error
	}

	if len(result) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if result[0].DistributionId == 0 {
		return nil, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(result)
	if err := r.redis.Set(ctx, key, jsonData, 5*time.Minute).Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *distributorRepository) GetListDistributionsByDistributorId(ctx context.Context, id uint) ([]dto.GetListDistribution, error) {
	var result []dto.GetListDistribution
	key := fmt.Sprintf("distribution:distributor:%d", id)

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err == nil {
			return result, nil
		}
	}

	res := r.db.WithContext(ctx).Table("distributions AS d").
		Select(
			"d.id",
			"dp.name AS distributor_name",
			"COALESCE(c1.crop_name, c2.crop_name, c3.crop_name) AS crop_name",
			"rg.name AS destination",
			"d.name AS name",
			"d.price AS price",
			"d.base_price AS base_price",
			"d.quantity AS quantity",
			"d.transportation AS transportation",
			"d.status as status",
			"rg.name AS regency_name",
			"d.updated_at AS time",
		).
		Joins("JOIN regencies AS rg ON rg.id = d.destination_id").
		Joins("JOIN distributor_profiles AS dp ON dp.id = d.distributor_profile_id").
		Joins("LEFT JOIN harvest_processors AS hp ON  hp.id = d.harvest_processor_id").
		Joins("LEFT JOIN harvest_collectors AS hc ON  hc.id = d.harvest_collector_id").
		Joins("LEFT JOIN harvests AS h1 ON  h1.id  =  d.harvest_id").
		Joins("LEFT JOIN harvests AS h2 ON  h2.id  =  hp.harvest_id").
		Joins("LEFT JOIN harvests AS h3  ON  h3.id  =  hc.harvest_id").
		Joins("LEFT JOIN crops AS c1  ON  c1.id  =  h1.crop_id").
		Joins("LEFT JOIN crops AS c2  ON  c2.id  =  h2.crop_id").
		Joins("LEFT JOIN crops AS c3  ON  c3.id  =  h3.crop_id").
		Where("").
		Scan(&result)

	if res.Error != nil {
		return nil, res.Error
	}

	if len(result) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if result[0].DistributionId == 0 {
		return nil, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(result)
	if err := r.redis.Set(ctx, key, jsonData, 5*time.Minute).Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *distributorRepository) GetDistributionById(ctx context.Context, id uint) (*dto.GetDistributionById, error) {
	var result dto.GetDistributionById

	key := fmt.Sprintf("distribution:%d", id)

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err == nil {
			return &result, nil
		}
	}

	res := r.db.WithContext(ctx).Debug().Table("distributions AS d").
		Select(
			"d.id AS id",
			"dp.name AS ditributor_name",
			"COALESCE(c1.crop_name, c2.crop_name, c3.crop_name) AS crop_name",
			"d.name AS name",
			"d.price AS price",
			"d.base_price AS base_price",
			"d.`desc`AS description",
			"d.quantity AS quantity",
			"d.transportation AS transportation",
			"d.status AS status",
			"d.distribution_status AS distribution_status",
			"d.tx_block AS tx_block",
			"d.updated_at AS time",
			"rg.name AS regency_name",
			"rn.name AS region_name",
		).
		Joins("JOIN distributor_profiles AS dp ON dp.id = d.distributor_profile_id").
		Joins("JOIN regencies AS rg ON rg.id = d.destination_id").
		Joins("JOIN regions AS rn ON rn.id = rg.region_id").
		Joins("LEFT JOIN harvest_processors AS hp ON  hp.id  =  d.harvest_processor_id").
		Joins("LEFT JOIN harvest_collectors AS hc ON  hc.id  =  d.harvest_collector_id").
		Joins("LEFT JOIN harvests AS h1 ON  h1.id  =  d.harvest_id").
		Joins("LEFT JOIN harvests AS h2 ON  h2.id  =  hp.harvest_id").
		Joins("LEFT JOIN harvests AS h3  ON  h3.id  =  hc.harvest_id").
		Joins("LEFT JOIN crops AS c1  ON  c1.id  =  h1.crop_id").
		Joins("LEFT JOIN crops AS c2  ON  c2.id  =  h2.crop_id").
		Joins("LEFT JOIN crops AS c3  ON  c3.id  =  h3.crop_id").
		Where("d.id = ?", id).
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

func (r *distributorRepository) GetListDistributionFYP(ctx context.Context) ([]dto.GetListDistribution, error) {
	var result []dto.GetListDistribution

	key := "distribution:fyp"

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err == nil {
			return result, nil
		}
	}

	res := r.db.WithContext(ctx).Debug().Table("distributions as d").
		Select(
			"d.id",
			"dp.name AS distributor_name",
			"COALESCE(c1.crop_name, c2.crop_name, c3.crop_name) AS crop_name",
			"rg.name AS destination",
			"d.name AS name",
			"d.price AS price",
			"d.base_price AS base_price",
			"d.quantity AS quantity",
			"d.transportation AS transportation",
			"d.status as status",
			"rg.name AS regency_name",
			"d.updated_at AS time",
		).
		Joins("JOIN regencies AS rg ON rg.id = d.destination_id").
		Joins("JOIN distributor_profiles AS dp ON dp.id = d.distributor_profile_id").
		Joins("LEFT JOIN harvest_processors AS hp ON  hp.id = d.harvest_processor_id").
		Joins("LEFT JOIN harvest_collectors AS hc ON  hc.id = d.harvest_collector_id").
		Joins("LEFT JOIN harvests AS h1 ON  h1.id  =  d.harvest_id").
		Joins("LEFT JOIN harvests AS h2 ON  h2.id  =  hp.harvest_id").
		Joins("LEFT JOIN harvests AS h3  ON  h3.id  =  hc.harvest_id").
		Joins("LEFT JOIN crops AS c1  ON  c1.id  =  h1.crop_id").
		Joins("LEFT JOIN crops AS c2  ON  c2.id  =  h2.crop_id").
		Joins("LEFT JOIN crops AS c3  ON  c3.id  =  h3.crop_id").
		Order("RAND()").
		Scan(&result)
	if res.Error != nil {
		return nil, res.Error
	}

	if len(result) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if result[0].DistributionId == 0 {
		return nil, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(result)
	r.redis.Set(ctx, key, jsonData, 5*time.Minute)

	return result, nil
}
