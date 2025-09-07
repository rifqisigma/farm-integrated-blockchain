package dto

import "time"

type HarvestRequest struct {
	FarmerProfileId uint    `json:"-" validate:"required"`
	CropID          uint    `json:"-" validate:"required"`
	Quantity        float64 `json:"quantity" validate:"required"`
	BasePrice       float64 `gorm:"base_price" validate:"required"`
}

type ValidateTimeHarvest struct {
	HarvestTime     time.Time `json:"harvest_time"`
	HarvestTimeSpan int       `json:"harvest_time_span"`
}

type HarvestUpdate struct {
	FarmerProfileId uint `json:"-" validate:"required"`
	HarvestId       uint `json:"-" validate:"required"`

	Quantity  float64 `json:"quantity" validate:"required"`
	BasePrice float64 `gorm:"base_price" validate:"required"`
}
type AcceptHarvest struct {
	FarmerProfileId uint `json:"-" validate:"required"`
	HarvestId       uint `json:"-" validate:"required"`
	Accepted        bool `json:"accepted" validate:"required"`
}

type AcceptFarmerForDistributor struct {
	FarmerProfieId uint `json:"-" validate:"required"`
	DistributionId uint `json:"-" validate:"required"`
	Accepted       bool `json:"accepted" validate:"required"`
}

//get
type GetListHarvest struct {
	Id         uint      `json:"id" gorm:"id"`
	CropName   string    `json:"crop_name" gorm:"crop_name"`
	FarmerName string    `json:"farmer_name" gorm:"farmer_name"`
	BasePrice  float64   `json:"base_price" gorm:"base_price"`
	Time       time.Time `json:"time" gorm:"time"`
}

type GetHarvestById struct {
	Id         uint      `json:"id" gorm:"id"`
	CropName   string    `json:"crop_name" gorm:"crop_name"`
	FarmerName string    `json:"farmer_name" gorm:"farmer_name"`
	BasePrice  float64   `json:"base_price" gorm:"base_price"`
	Quantity   float64   `json:"quantity" gorm:"quantity"`
	Accepted   bool      `json:"accepted" gorm:"accepted"`
	Time       time.Time `json:"time" gorm:"time"`
}

type DataValidationFarm struct {
	Quantity float64   `json:"quantity" gorm:"quantity"`
	Time     time.Time `json:"time" gorm:"create_time"`
}
