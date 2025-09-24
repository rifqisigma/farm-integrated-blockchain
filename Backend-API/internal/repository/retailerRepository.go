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

type RetailerRepository interface {
	AddRetailerCart(ctx context.Context, input *dto.CreateRetailerCartRequest) error
	UpdateRetailerCart(ctx context.Context, input *dto.UpdateRetailerCartRequest) error
	DeleteRetailerCart(ctx context.Context, retailerCartId, retailerId uint) error
	SearchRetailerCart(ctx context.Context, search string) ([]dto.GetRetailerCart, error)
	GetRetailerCarts(ctx context.Context, retailerProfileId uint) ([]dto.GetRetailerCart, error)
	GetRetailerCartById(ctx context.Context, retailerCartId uint) (*dto.GetRetailerCartById, error)
	GetRetailerCartsFYP(ctx context.Context) ([]dto.GetRetailerCart, error)
}

type retailerRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewRetailerRepository(db *gorm.DB, redis *redis.Client) RetailerRepository {
	return &retailerRepository{db, redis}
}

func (r *retailerRepository) AddRetailerCart(ctx context.Context, input *dto.CreateRetailerCartRequest) error {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var quantity float64
	res := tx.Model(&entity.Distribution{}).Select("quantity").Where("id = ? AND approved_by_farmer = ? AND is_canceled = ? ", input.DistributionId, true, false).Scan(&quantity)
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
		return gorm.ErrInvalidData
	}

	newRetailerCart := entity.RetailerCart{
		Quantity:          input.Quantity,
		RetailerProfileId: input.RetailerProfileId,
		DistributionId:    input.DistributionId,
	}

	if err := tx.Model(&entity.RetailerCart{}).Create(&newRetailerCart).Error; err != nil {
		tx.Rollback()
		return err
	}

	keyRetailer := fmt.Sprintf("retail:retailer:%d", newRetailerCart.RetailerProfileId)
	keyRetail := fmt.Sprintf("retail:%d", newRetailerCart.ID)
	keyFYP := fmt.Sprintln("retail:fyp")
	_, err := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keyRetailer)
		p.Del(ctx, keyRetail)
		p.Del(ctx, keyFYP)
		keys, _ := r.redis.Keys(ctx, "retailer:search:*").Result()
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

func (r *retailerRepository) UpdateRetailerCart(ctx context.Context, input *dto.UpdateRetailerCartRequest) error {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var validation dto.DataValidationRetailer
	res := tx.Model(&entity.Distribution{}).Select("quantity, create_time").Where("id = ? AND approved_by_farmer = ? AND is_canceled =?", input.DistributionId, true, false).Scan(&validation)
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

	updated := make(map[string]interface{})

	if input.Quantity > 0 {
		updated["quantity"] = input.Quantity
	}

	res2 := tx.Model(&entity.RetailerCart{}).Where("id = ? AND retailer_profile_id = ? AND is_canceled = ?", input.RetailerCartId, input.RetailerProfileId, false).UpdateColumns(updated)
	if res2.Error != nil {
		return res2.Error
	}

	if res2.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	keyRetailer := fmt.Sprintf("retail:retailer:%d", input.RetailerProfileId)
	keyRetail := fmt.Sprintf("retail:%d", input.RetailerCartId)
	keyFYP := fmt.Sprintln("retail:fyp")
	_, err := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keyRetailer)
		p.Del(ctx, keyRetail)
		p.Del(ctx, keyFYP)
		keys, _ := r.redis.Keys(ctx, "retailer:search:*").Result()
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

func (r *retailerRepository) DeleteRetailerCart(ctx context.Context, retailerCartId, retailerId uint) error {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var dataTime time.Time
	res := tx.Model(&entity.RetailerCart{}).Select("create_time").Where("id = ? AND approved_by_farmer = ? AND is_canceled = ?", retailerCartId, true, false).Scan(&dataTime)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	if time.Since(dataTime) > 30*time.Minute {
		tx.Rollback()
		return helper.ErrInvalidTime
	}

	res2 := tx.Model(&entity.RetailerCart{}).Where("id = ? AND retailer_profile_id = ?", retailerCartId, retailerId).Update("is_canceled", true)
	if res2.Error != nil {
		return res2.Error
	}
	if res2.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	keyRetailer := fmt.Sprintf("retail:retailer:%d", retailerId)
	keyRetail := fmt.Sprintf("retail:%d", retailerCartId)
	keyFYP := fmt.Sprintln("retail:fyp")
	_, err := r.redis.Pipelined(ctx, func(p redis.Pipeliner) error {
		p.Del(ctx, keyRetailer)
		p.Del(ctx, keyRetail)
		p.Del(ctx, keyFYP)
		keys, _ := r.redis.Keys(ctx, "retailer:search:*").Result()
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

func (r *retailerRepository) SearchRetailerCart(ctx context.Context, search string) ([]dto.GetRetailerCart, error) {
	input := "%" + search + "%"
	var result []dto.GetRetailerCart

	key := fmt.Sprintf("retailer:search:%s", search)
	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err != nil {
			return result, nil
		}
	}

	if err := r.db.WithContext(ctx).Model(&entity.RetailerCart{}).
		Select("retailer_carts.id as id", "cp.crop as hs_name", "dp.name as distributor_name", "rs.name as retailer_name", "retailer_carts.quantity", "retailer_carts.update_time as time").
		Joins("JOIN distributions as ds ON ds.id = retailer_carts.id").
		Joins("JOIM distrbutor_profiles as dp ON dp.id = retailer_carts.distributor_profile_id").
		Joins("JOIN retailer_profiles as a rs ON rs.id = retailer_carts.retailer_profile_id").
		Joins("JOIN harvests as hs ON hs.id = ds.id").
		Joins("JOIN crops cs ON cs.id = hs.id").
		Where("cs.crop LIKE ? OR hs.name LIKE ? OR rs.name LIKE ? OR dp.name LIKE ? OR retailer_carts.name LIKE ?", input, input, input, input, input).
		Scan(result).
		Error; err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return result, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(result)
	if err := r.redis.Set(ctx, key, jsonData, 5*time.Minute).Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *retailerRepository) GetRetailerCarts(ctx context.Context, retailerProfileId uint) ([]dto.GetRetailerCart, error) {
	var result []dto.GetRetailerCart
	key := fmt.Sprintf("retail:retailer:%d", retailerProfileId)

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err != nil {
			return result, nil
		}
	}

	if err := r.db.WithContext(ctx).Model(&entity.RetailerCart{}).
		Select("retailer_carts.id as id", "cp.crop as hs_name", "dp.name as distributor_name", "rs.name as retailer_name", "retailer_carts.quantity", "retailer_carts.update_time as time").
		Joins("JOIN distributions as ds ON ds.id = retailer_carts.id").
		Joins("JOIM distrbutor_profiles as dp ON dp.id = retailer_carts.distributor_profile_id").
		Joins("JOIN retailer_profiles as a rs ON rs.id = retailer_carts.retailer_profile_id").
		Joins("JOIN harvests as hs ON hs.id = ds.id").
		Joins("JOIN crops cs ON cs.id = hs.id").
		Where("retailer_carts.retailer_profile_id = ?", retailerProfileId).
		Scan(result).
		Error; err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return result, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(result)
	if err := r.redis.Set(ctx, key, jsonData, 5*time.Minute).Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *retailerRepository) GetRetailerCartById(ctx context.Context, retailerCartId uint) (*dto.GetRetailerCartById, error) {
	var result dto.GetRetailerCartById
	key := fmt.Sprintf("retail:%d", retailerCartId)

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err != nil {
			return &result, nil
		}
	}

	res := r.db.WithContext(ctx).Model(&entity.RetailerCart{}).
		Select("retailer_carts.id as id", "cp.crop as hs_name", "dp.name as distributor_name", "rs.name as retailer_name", "retailer_carts.quantity", "retailer_carts.update_time as time", "retailer_carts.block_hash as block_hash").
		Joins("JOIN distributions as ds ON ds.id = retailer_carts.id").
		Joins("JOIM distrbutor_profiles as dp ON dp.id = retailer_carts.distributor_profile_id").
		Joins("JOIN retailer_profiles as a rs ON rs.id = retailer_carts.retailer_profile_id").
		Joins("JOIN harvests as hs ON hs.id = ds.id").
		Joins("JOIN crops cs ON cs.id = hs.id").
		Where("retailer_carts.id = ?").
		Scan(result)
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

func (r *retailerRepository) GetRetailerCartsFYP(ctx context.Context) ([]dto.GetRetailerCart, error) {
	var result []dto.GetRetailerCart
	key := fmt.Sprintln("retail:fyp")

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err != nil {
			return result, nil
		}
	}
	if err := r.db.WithContext(ctx).Model(&entity.RetailerCart{}).
		Select(
			"retailer_carts.id as id",
			"cp.crop as hs_name",
			"dp.name as distributor_name",
			"rs.name as retailer_name",
			"retailer_carts.quantity",
			"retailer_carts.update_time as time").
		Joins("JOIN distributions as ds ON ds.id = retailer_carts.id").
		Joins("JOIM distrbutor_profiles as dp ON dp.id = retailer_carts.distributor_profile_id").
		Joins("JOIN retailer_profiles as a rs ON rs.id = retailer_carts.retailer_profile_id").
		Joins("JOIN harvests as hs ON hs.id = ds.id").
		Joins("JOIN crops cs ON cs.id = hs.id").
		Order("RAND()").
		Scan(result).
		Error; err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return result, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(result)
	if err := r.redis.Set(ctx, key, jsonData, 5*time.Minute).Err(); err != nil {
		return nil, err
	}

	return result, nil
}
