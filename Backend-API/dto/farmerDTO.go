package dto

import "time"

type HarvestRequest struct {
	FarmerProfileId uint    `json:"-" validate:"required"`
	CropID          uint    `json:"-" validate:"required"`
	Name            string  `json:"name" validate:"required" example:"Jagung ke-1"`
	Quantity        float64 `json:"quantity" validate:"required" example:"100"`
	BasePrice       float64 `gorm:"base_price" validate:"required" example:"200000000"`
	CountryRequest  `json:"country" validate:"required"`
}

type ValidateTimeHarvest struct {
	HarvestTime     time.Time `json:"harvest_time"`
	HarvestTimeSpan int       `json:"harvest_time_span"`
}

type HarvestUpdate struct {
	FarmerProfileId uint    `json:"-" validate:"required"`
	HarvestId       uint    `json:"-" validate:"required"`
	Name            string  `json:"name" validate:"required" example:"Jagung ke-1"`
	Quantity        float64 `json:"quantity" validate:"required" example:"100"`
	BasePrice       float64 `gorm:"base_price" validate:"required" example:"200000000"`
	CountryRequest  `json:"country" validate:"required"`
}
type AcceptHarvest struct {
	FarmerProfileId uint `json:"-" validate:"required"`
	HarvestId       uint `json:"-" validate:"required"`
	Accepted        bool `json:"accepted" validate:"required" example:"true"`
}

type AcceptFarmerForDistributor struct {
	FarmerProfieId       uint `json:"-" validate:"required" example:"644"`
	DistributionId       uint `json:"-" validate:"required" example:"1231"`
	DistributorProfileId uint `json:"-" validate:"required" example:"313"`
	Accepted             bool `json:"accepted" validate:"required" example:"true"`
}

//get
type GetListHarvest struct {
	Id          uint      `json:"id" gorm:"column:id" example:"1"`
	CropName    string    `json:"crop_name" gorm:"column:crop_name" example:"rice"`
	FarmerName  string    `json:"farmer_name" gorm:"column:farmer_name" example:"Kartosuryo"`
	BasePrice   float64   `json:"base_price" gorm:"column:base_price" example:"100000000"`
	RegencyName string    `json:"regency_name" gorm:"column:regency_name" example:"Bogor"`
	Time        time.Time `json:"time" gorm:"column:time" example:"2025-10-09T15:04:05Z"`
	Country     Country
}

type GetHarvestById struct {
	Id          uint      `json:"id" gorm:"column:id" example:"192"`
	CropName    string    `json:"crop_name" gorm:"column:crop_name" example:"Corn"`
	FarmerName  string    `json:"farmer_name" gorm:"column:farmer_name" example:"Wayulo"`
	BasePrice   float64   `json:"base_price" gorm:"column:base_price" example:"100000000"`
	Quantity    float64   `json:"quantity" gorm:"column:quantity" example:"100"`
	Accepted    bool      `json:"accepted" gorm:"column:accepted" example:"true"`
	RegencyName string    `json:"regency_name" gorm:"column:regency_name" example:"Bogor"`
	Time        time.Time `json:"time" gorm:"column:time" example:"2025-10-09T15:04:05Z"`
}

type DataValidationFarm struct {
	Quantity float64   `json:"quantity" gorm:"column:quantity"`
	Time     time.Time `json:"time" gorm:"column:create_time"`
}
