package repository

import (
	"errors"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/entity"
	"farm-integrated-web3/utils/helper"
	"time"

	"gorm.io/gorm"
)

type FarmerRepository interface {
	CreateHarvest(input *dto.HarvestRequest) error
	UpdateHarvest(input *dto.HarvestUpdate) error
	DeleteHarvest(farmerProfileId, harvestId uint) error
	AcceptedFarmerForDistributor(input *dto.AcceptFarmerForDistributor) error

	//get
	ListHarvestByFarmerId(farmerId uint) ([]dto.GetListHarvest, error)
	HarvestById(harvestId uint) (*dto.GetHarvestById, error)
	SearchHarvest(search string) ([]dto.GetListHarvest, error)
}

type farmerRepository struct {
	db *gorm.DB
}

func NewFarmerRepository(db *gorm.DB) FarmerRepository {
	return &farmerRepository{db}
}

func (r *farmerRepository) CreateHarvest(input *dto.HarvestRequest) error {
	tx := r.db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	newHarvest := entity.Harvest{
		FarmerProfileId: input.FarmerProfileId,
		CropId:          input.CropID,
		Quantity:        input.Quantity,
		BasePrice:       input.BasePrice,
	}

	var dataTime dto.ValidateTimeHarvest
	if err := tx.Model(&entity.Harvest{}).
		Select("update_time as harvest_time", "crops.harvest_time_span as harvest_time_span").
		Where("farmer_profile_id = ?", input.FarmerProfileId).
		Joins("JOIN crops ON crops.id = harvest.crop_id").
		Order("harvests.update_time desc").
		First(&dataTime).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := tx.Create(&newHarvest).Error; err != nil {
				tx.Rollback()
				return err
			}
			return nil
		}

		tx.Rollback()
		return err
	}

	rentang := time.Duration(dataTime.HarvestTimeSpan) * 24 * time.Hour
	selisih := time.Since(dataTime.HarvestTime)

	if selisih > rentang {
		if err := tx.Create(&newHarvest).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		return helper.ErrInvalidTime
	}

	return tx.Commit().Error

}

func (r *farmerRepository) UpdateHarvest(input *dto.HarvestUpdate) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
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

	err2 := tx.Model(&entity.Harvest{}).Where("id = ? AND farmer_profile_id = ?", input.HarvestId, input.FarmerProfileId).UpdateColumns(&updatedHarvest)
	if err2.Error != nil {
		tx.Rollback()
		return err2.Error
	}

	if err2.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	return tx.Commit().Error
}

func (r *farmerRepository) DeleteHarvest(farmerProfileId, harvestId uint) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
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

	err2 := r.db.Model(&entity.Harvest{}).Where("id = ? AND farmer_profile_id ", harvestId, farmerProfileId).UpdateColumn("is_canceled", true)
	if err2.Error != nil {
		return err2.Error
	}

	if err2.RowsAffected == 0 {
		tx.Rollback()
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *farmerRepository) AcceptedFarmerForDistributor(input *dto.AcceptFarmerForDistributor) error {

	update := make(map[string]interface{})
	update["approved_by_farmer"] = input.Accepted
	if !input.Accepted {
		update["is_canceled"] = true
	}
	res := r.db.Model(&entity.Distribution{}).Where("distributions.id = ? AND distribution.farmer_profile_id = ? AND is_canceled = ?", input.DistributionId, input.FarmerProfieId, false).UpdateColumns(update)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *farmerRepository) ListHarvestByFarmerId(farmerId uint) ([]dto.GetListHarvest, error) {
	var harvests []dto.GetListHarvest

	res := r.db.Model(&entity.Harvest{}).
		Select("cp.crop as crop_name", "fp.name as farmer_name", "harvests.id as id", "update_time as time", "base_price").
		Joins("JOIN crops as cp ON cp.id = harvests.crop_id").
		Joins("JOIN farmer_profiles as fp ON fp.id = harvests.farmer_profile_id").
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

	return harvests, nil
}

func (r *farmerRepository) HarvestById(harvestId uint) (*dto.GetHarvestById, error) {
	var harvest dto.GetHarvestById
	res := r.db.Model(&entity.Harvest{}).
		Select("cp.crop as crop_name", "fp.name as farmer_name", "harvests.id", "base_price", "update_time as time", "quantity").
		Joins("JOIN crops as cp ON cp.id = harvests.crop_id").
		Joins("JOIN farmer_profiles as fp ON fp.id = harvests.farmer_profile_id").
		Where("harvests.id = ?", harvestId).
		Scan(&harvest)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &harvest, nil
}

func (r *farmerRepository) SearchHarvest(search string) ([]dto.GetListHarvest, error) {
	var harvests []dto.GetListHarvest

	input := "%" + search + "%"
	res := r.db.Model(&entity.Harvest{}).
		Select("cp.crop as crop_name", "fp.name as farmer_name", "harvets.id", "update_time as time", "base_price").
		Joins("JOIN crops as cp ON cp.id = harvests.crop_id").
		Joins("JOIN farmer_profiles as fp ON fp.id = harvests.farmer_profile_id").
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

	return harvests, nil
}
