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

type SellerRepository interface {
	AddSellerBox(ctx context.Context, input *dto.CreateSellerBox) error
	UpdateSellerBox(ctx context.Context, input *dto.UpdateSellerBox) error
	DeleteSellerBox(ctx context.Context, SellerBoxId, retailerId uint) error

	//blockchain
	UpdateBcBlockSellerBox(ctx context.Context, sellerBoxId uint, txBlock string) error
	//get
	SearchSellerBox(ctx context.Context, search string) ([]dto.GetSellerBox, error)
	ListGetSellerBoxsbySellerId(ctx context.Context, SellerProfileId uint) ([]dto.GetSellerBox, error)
	GetSellerBoxById(ctx context.Context, SellerBoxId uint) (*dto.GetSellerBoxById, error)
	ListGetSellerBoxsbySellerIdFYP(ctx context.Context) ([]dto.GetSellerBox, error)
}

type sellerRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewSellerRepository(db *gorm.DB, redis *redis.Client) SellerRepository {
	return &sellerRepository{db, redis}
}

func (r *sellerRepository) AddSellerBox(ctx context.Context, input *dto.CreateSellerBox) error {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var quantity float64
	res := tx.Debug().Model(&entity.Distribution{}).Select("quantity").Where("id = ? AND status = ? ", input.DistributionId, 1).First(&quantity)
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

	newSellerBox := entity.SellerBox{
		Quantity:        input.Quantity,
		SellerProfileId: input.SellerProfileId,
		DistributionId:  input.DistributionId,
		Name:            input.Name,
		Desc:            input.Desc,
		BasePrice:       input.BasePrice,
		Price:           input.Price,
	}

	if err := tx.Debug().Create(&newSellerBox).Error; err != nil {
		tx.Rollback()
		return err
	}

	keySeller := fmt.Sprintf("sellerBox:seller:%d", newSellerBox.SellerProfileId)
	keySellerBox := fmt.Sprintf("sellerBox:%d", newSellerBox.ID)
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
		return err
	}

	return tx.Commit().Error
}

func (r *sellerRepository) UpdateSellerBox(ctx context.Context, input *dto.UpdateSellerBox) error {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	updated := make(map[string]interface{})
	if input.Quantity != 0 {
		var quantity float64
		res := tx.Debug().Table("seller_boxes as sb").Select("d.quantity").
			Joins("JOIN distributions AS d ON d.id =  sb.distribution_id").
			Where("sb.id = ? AND d.status = ?", input.SellerBoxId, 1).Scan(&quantity)
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

		if input.Quantity > 0 {
			updated["quantity"] = input.Quantity
		}

	}

	if input.Name != "" {
		updated["name"] = input.Name

	}
	if input.Desc != "" {
		updated["desc"] = input.Desc

	}
	if input.BasePrice != 0 {
		updated["base_price"] = input.BasePrice
	}

	if input.Price != 0 {
		updated["price"] = input.Price
	}

	res2 := tx.Debug().Model(&entity.SellerBox{}).Where("id = ? AND seller_profile_id = ? AND status = ?", input.SellerBoxId, input.SellerProfileId, 0).Updates(updated)
	if res2.Error != nil {
		return res2.Error
	}

	if res2.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	keySeller := fmt.Sprintf("sellerBox:seller:%d", input.SellerProfileId)
	keySellerBox := fmt.Sprintf("sellerBox:%d", input.SellerBoxId)
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
		return err
	}

	return tx.Commit().Error
}

func (r *sellerRepository) DeleteSellerBox(ctx context.Context, sellerBoxId, retailerId uint) error {

	res := r.db.Debug().Model(&entity.SellerBox{}).Where("id = ? AND seller_profile_id = ?", sellerBoxId, retailerId).Update("status", 2)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	keySeller := fmt.Sprintf("sellerBox:seller:%d", retailerId)
	keySellerBox := fmt.Sprintf("sellerBox:%d", sellerBoxId)
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
		return err
	}

	return nil
}

func (r *sellerRepository) UpdateBcBlockSellerBox(ctx context.Context, sellerBoxId uint, txBlock string) error {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	var sellerBox entity.SellerBox
	res := tx.Debug().Model(&entity.SellerBox{}).Select("id", "seller_profile_id").Where("id = ?", sellerBoxId).First(&sellerBox)
	if res.Error != nil {
		tx.Commit()
		return res.Error
	}
	if res.RowsAffected == 0 {
		tx.Commit()
		return gorm.ErrRecordNotFound
	}
	res2 := r.db.Debug().Model(&entity.SellerBox{}).Where("id = ?", sellerBox.ID).Updates(map[string]interface{}{
		"tx_block": txBlock,
		"status":   1,
	})
	if res2.Error != nil {
		tx.Commit()
		return res2.Error
	}

	if res2.RowsAffected == 0 {
		tx.Commit()
		return gorm.ErrRegistered
	}

	keySeller := fmt.Sprintf("sellerBox:seller:%d", sellerBox.SellerProfileId)
	keySellerBox := fmt.Sprintf("sellerBox:%d", sellerBox.ID)
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
		return err
	}

	return tx.Commit().Error
}

func (r *sellerRepository) SearchSellerBox(ctx context.Context, search string) ([]dto.GetSellerBox, error) {
	input := "%" + search + "%"
	var result []dto.GetSellerBox

	key := fmt.Sprintf("sellerBox:search:%s", search)
	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err != nil {
			return result, nil
		}
	}

	res := r.db.WithContext(ctx).Debug().Table("seller_profiles as sp").
		Select(
			"sb.id AS id",
			"COALESCE(c1.crop_name, c2.crop_name, c3.crop_name) AS crop_name",
			"sp.name AS seller_name",
			"sb.name AS name",
			"sb.base_price AS base_price",
			"sb.price AS price",
			"sb.quantity AS quantity",
			"sb.updated_at AS time",
		).
		Joins("LEFT JOIN seller_boxes AS sb ON sb.seller_profile_id = sp.id").
		Joins("LEFT JOIN distributions AS d ON d.id = sb.distribution_id").
		Joins("LEFT JOIN harvest_processors as hp ON hp.id = d.harvest_processor_id").
		Joins("LEFT JOIN harvest_collectors as hc ON hc.id = d.harvest_collector_id").
		Joins("LEFT JOIN harvests AS h1 ON h1.id = d.harvest_id").
		Joins("LEFT JOIN harvests AS h2 ON h2.id = hp.harvest_id").
		Joins("LEFT JOIN harvests AS h3 ON h3.id = hc.harvest_id").
		Joins("LEFT JOIN crops c1 ON c1.id = h1.crop_id").
		Joins("LEFT JOIN crops c2 ON c2.id = h2.crop_id").
		Joins("LEFT JOIN crops c3 ON c3.id = h3.crop_id").
		Where("c1.crop_name LIKE ? OR c2.crop_name LIKE ? OR c3.crop_name LIKE ? OR sp.name LIKE ? OR sb.name LIKE ?", input, input, input, input, input).
		Scan(&result)
	if res.Error != nil {
		return nil, res.Error
	}

	if len(result) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if result[0].ID == 0 {
		return nil, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(result)
	if err := r.redis.Set(ctx, key, jsonData, 5*time.Minute).Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *sellerRepository) ListGetSellerBoxsbySellerId(ctx context.Context, SellerProfileId uint) ([]dto.GetSellerBox, error) {
	var result []dto.GetSellerBox
	key := fmt.Sprintf("sellerBox:seller:%d", SellerProfileId)

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err != nil {
			return result, nil
		}
	}

	res := r.db.WithContext(ctx).Debug().Table("seller_profiles AS sp").
		Select(
			"sb.id AS id",
			"COALESCE(c1.crop_name, c2.crop_name, c3.crop_name) AS crop_name",
			"sp.name AS seller_name",
			"sb.name AS name",
			"sb.base_price AS base_price",
			"sb.price AS price",
			"sb.quantity AS quantity",
			"sb.updated_at AS time",
		).
		Joins("LEFT JOIN seller_boxes AS sb ON sb.seller_profile_id = sp.id").
		Joins("LEFT JOIN distributions AS d ON d.id = sb.distribution_id").
		Joins("LEFT JOIN harvest_processors AS hp ON hp.id = d.harvest_processor_id").
		Joins("LEFT JOIN harvest_collectors AS hc ON hc.id = d.harvest_collector_id").
		Joins("LEFT JOIN harvests AS h1 ON h1.id = d.harvest_id").
		Joins("LEFT JOIN harvests AS h2 ON h2.id = hp.harvest_id").
		Joins("LEFT JOIN harvests AS h3 ON h3.id = hc.harvest_id").
		Joins("LEFT JOIN crops AS c1 ON c1.id = h1.crop_id").
		Joins("LEFT JOIN crops AS c2 ON c2.id = h2.crop_id").
		Joins("LEFT JOIN crops AS c3 ON c3.id = h3.crop_id").
		Where("sp.id = ?", SellerProfileId).
		Scan(&result)

	if res.Error != nil {
		return nil, res.Error
	}

	if len(result) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if result[0].ID == 0 {
		return nil, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(result)
	if err := r.redis.Set(ctx, key, jsonData, 5*time.Minute).Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *sellerRepository) GetSellerBoxById(ctx context.Context, SellerBoxId uint) (*dto.GetSellerBoxById, error) {
	var result dto.GetSellerBoxById
	key := fmt.Sprintf("sellerBox:%d", SellerBoxId)

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err != nil {
			return &result, nil
		}
	}

	res := r.db.WithContext(ctx).Debug().Table("seller_boxes AS sb").
		Select(`
			sb.id AS id,
			rs.name AS seller_name,
			COALESCE(c1.crop_name, c2.crop_name, c3.crop_name) AS crop_name,
			sb.name AS name,
			sb.desc AS description,
			sb.base_price AS base_price,
			"sb.price AS price",
			sb.tx_block AS tx_block,
			sb.quantity AS quantity,
			sb.status AS status,
			sb.updated_at AS time
		`).
		Joins("LEFT JOIN seller_profiles AS rs ON rs.id = sb.seller_profile_id").
		Joins("LEFT JOIN distributions AS d ON d.id = sb.distribution_id").
		Joins("LEFT JOIN harvests AS h1 ON h1.id = d.harvest_id").
		Joins("LEFT JOIN crops c1 ON c1.id = h1.crop_id").
		Joins("LEFT JOIN harvest_processors AS hp ON hp.id = d.harvest_processor_id").
		Joins("LEFT JOIN harvests AS h2 ON h2.id = hp.harvest_id").
		Joins("LEFT JOIN crops c2 ON c2.id = h2.crop_id").
		Joins("LEFT JOIN harvest_collectors AS hc ON hc.id = d.harvest_collector_id").
		Joins("LEFT JOIN harvests AS h3 ON h3.id = hc.harvest_id").
		Joins("LEFT JOIN crops c3 ON c3.id = h3.crop_id").
		Where("sb.id = ?", SellerBoxId).
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

func (r *sellerRepository) ListGetSellerBoxsbySellerIdFYP(ctx context.Context) ([]dto.GetSellerBox, error) {
	var result []dto.GetSellerBox
	key := fmt.Sprintln("sellerBox:fyp")

	val, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &result); err != nil {
			return result, nil
		}
	}
	res := r.db.WithContext(ctx).Debug().Table("seller_profiles as sp").
		Select(
			"sb.id AS id",
			"COALESCE(c1.crop_name, c2.crop_name, c3.crop_name) AS crop_name",
			"sp.name AS seller_name",
			"sb.name AS name",
			"sb.base_price AS base_price",
			"sb.price AS price",
			"sb.quantity AS quantity",
			"sb.updated_at AS time",
		).
		Joins("LEFT JOIN seller_boxes AS sb ON sb.seller_profile_id = sp.id").
		Joins("LEFT JOIN distributions AS d ON d.id = sb.distribution_id").
		Joins("LEFT JOIN harvest_processors as hp ON hp.id = d.harvest_processor_id").
		Joins("LEFT JOIN harvest_collectors as hc ON hc.id = d.harvest_collector_id").
		Joins("LEFT JOIN harvests AS h1 ON h1.id = d.harvest_id").
		Joins("LEFT JOIN harvests AS h2 ON h2.id = hp.harvest_id").
		Joins("LEFT JOIN harvests AS h3 ON h3.id = hc.harvest_id").
		Joins("LEFT JOIN crops c1 ON c1.id = h1.crop_id").
		Joins("LEFT JOIN crops c2 ON c2.id = h2.crop_id").
		Joins("LEFT JOIN crops c3 ON c3.id = h3.crop_id").
		Order("RAND()").
		Scan(&result)

	if res.Error != nil {
		return nil, res.Error
	}

	if len(result) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	if result[0].ID == 0 {
		return nil, gorm.ErrEmptySlice
	}

	jsonData, _ := json.Marshal(result)
	if err := r.redis.Set(ctx, key, jsonData, 5*time.Minute).Err(); err != nil {
		return nil, err
	}

	return result, nil
}
