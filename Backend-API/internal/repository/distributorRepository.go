package repository

import (
	"farm-integrated-web3/dto"
	"farm-integrated-web3/entity"
	"farm-integrated-web3/utils/helper"
	"time"

	"gorm.io/gorm"
)

type DistributorRepository interface {
	CreateDistribution(input *dto.CreateDistributionRequest) error
	UpdateDistribution(input *dto.UpdateDistributionRequest) error
	DeleteDistribution(distrbutionId, distributorId uint) error
	UpdateStatusOfDistribution(input *dto.UpdateStatusDistributionRequest) error
	ApprovedRetailerCartForRetailer(input *dto.ApprovedRetailerCart) error
	//get
	SearchDistributions(search string) ([]dto.GetDistribution, error)
	GetDistributionsByDistributorId(id uint) ([]dto.GetDistribution, error)
	GetDistributionByid(id uint) (*dto.GetDistributionById, error)
}

type distributorRepository struct {
	db *gorm.DB
}

func NewDistributorRepository(db *gorm.DB) DistributorRepository {
	return &distributorRepository{db}
}

func (r *distributorRepository) CreateDistribution(input *dto.CreateDistributionRequest) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

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

	if err := tx.Model(&entity.Distribution{}).Create(&entity.Distribution{
		FinalPrice:      input.FinalPrice,
		MarkUpPrice:     input.MarkupPrice,
		HarvestId:       input.HarvestId,
		FarmerProfileId: input.FarmerProfileId,
	}).Error; err != nil {
		tx.Rollback()
	}

	return tx.Commit().Error
}

func (r *distributorRepository) UpdateDistribution(input *dto.UpdateDistributionRequest) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
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

	return tx.Commit().Error
}

func (r *distributorRepository) DeleteDistribution(distrbutionId, distributorId uint) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
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

	return nil

}

func (r *distributorRepository) UpdateStatusOfDistribution(input *dto.UpdateStatusDistributionRequest) error {
	err := r.db.Model(&entity.Distribution{}).Where("id = ? AND distributor_profile_id = ? AND approved_by_farmer = ? AND is_canceled = ? ", input.DistributionId, input.DistributorProfileId, true, false).UpdateColumn("status_distribution", input.Status)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil

}

func (r *distributorRepository) SearchDistributions(search string) ([]dto.GetDistribution, error) {
	var result []dto.GetDistribution
	input := "%" + search + "%"

	res := r.db.Model(&entity.Distribution{}).
		Select("distributions.id as id", "dp.name as ditributor_name", "fp.name as farmer_name", "distributions.final_price as final_price", "cp.crop as crop_name", "update time as time").
		Joins("JOIN harvests as hs ON hs.id = distributions.harvest_id").
		Joins("JOIN crops as cs ON cs.id = hs.crop_id").
		Joins("JOIN farmer_profiles as fp ON fp.id = distributions.farmer_profile_id").
		Joins("JOIN distributor_profiles as dp ON dp.id = distributions.distributor_profile_id").
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

	return result, nil
}

func (r *distributorRepository) GetDistributionsByDistributorId(id uint) ([]dto.GetDistribution, error) {
	var result []dto.GetDistribution

	res := r.db.Model(&entity.Distribution{}).
		Select("distributions.id as id", "dp.name as ditributor_name", "fp.name as farmer_name", "distributions.final_price as final_price", "cp.crop as crop_name", "update_time as time").
		Joins("JOIN harvests as hs ON hs.id = distributions.harvest_id").
		Joins("JOIN crops as cs ON cs.id = hs.crop_id").
		Joins("JOIN farmer_profiles as fp ON fp.id = distributions.farmer_profile_id").
		Joins("JOIN distributor_profiles as dp ON dp.id = distributions.distributor_profile_id").
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

	return result, nil
}

func (r *distributorRepository) GetDistributionByid(id uint) (*dto.GetDistributionById, error) {
	var result dto.GetDistributionById

	res := r.db.Model(&entity.Distribution{}).
		Select("distributions.id as id", "dp.name as ditributor_name", "fp.name as farmer_name", "distributions.final_price as final_price", "cp.crop as crop_name", "distributions.block_hash as block_hash", "distributions.has_arrived as has_arrived", "update_tme as time").
		Joins("JOIN harvests as hs ON hs.id = distributions.harvest_id").
		Joins("JOIN crops as cs ON cs.id = hs.crop_id").
		Joins("JOIN farmer_profiles as fp ON fp.id = distributions.farmer_profile_id").
		Joins("JOIN distributor_profiles as dp ON dp.id = distributions.distributor_profile_id").
		Where("distributions.id = ?", id).
		Scan(&result)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &result, nil

}

func (r *distributorRepository) ApprovedRetailerCartForRetailer(input *dto.ApprovedRetailerCart) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

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
