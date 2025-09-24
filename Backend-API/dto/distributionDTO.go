package dto

import "time"

type CreateDistributionRequest struct {
	FarmerProfileId      uint    `json:"-" validate:"required" example:"122"`
	DistributorProfileId uint    `json:"-" validate:"required" example:"122"`
	HarvestId            uint    `json:"-" validate:"required" example:"211"`
	Quantity             float64 `json:"quantity" validate:"required" example:"200.50"`
	MarkupPrice          float64 `json:"markup_price" validate:"required" example:"20.000"`
	FinalPrice           float64 `json:"final_price" validate:"required" example:"220.000"`
}

type UpdateDistributionRequest struct {
	DistributionId       uint    `json:"-" `
	DistributorProfileId uint    `json:"-"`
	Quantity             float64 `json:"quantity" example:"20.70"`
	MarkupPrice          float64 `json:"markup_price"  example:"50.000"`
	FinalPrice           float64 `json:"final_price"  example:"350.000"`
}

type GetDistribution struct {
	DistributionId  uint      `json:"id"  gorm:"column:id"`
	DistributorName string    `json:"distributor_name"  gorm:"column:distributor_name"`
	FarmerName      string    `json:"farmer_name" gorm:"column:farmer_name"`
	CropName        string    `json:"crop_name" gorm:"column:crop_name"`
	FinalPrice      float64   `json:"final_price" gorm:"column:final_price"`
	Time            time.Time `json:"time" gorm:"column:time"`
	RegencyName     string    `json:"regency_name" gorm:"column:regency_name" example:"Bogor"`
}

type GetDistributionById struct {
	DistributionId  uint    `json:"id"  gorm:"column:id" example:"23324"`
	DistributorName string  `json:"distributor_name"  gorm:"column:distributor_name" example:"Wanasena"`
	FarmerName      string  `json:"farmer_name" gorm:"column:farmer_name" example:"Sukirman"`
	CropName        string  `json:"crop_name" gorm:"column:crop_name" example:"Corn"`
	FinalPrice      float64 `json:"final_price" gorm:"column:final_price" example:"30000000"`
	BlockHash       string  `json:"block_hash" gorm:"column:block_hash" example:"-"`
	HasArrived      bool    `json:"has_arrived" gorm:"column:has_arrived" example:"false"`
	Country         Country
	Time            time.Time `json:"time" gorm:"column:time" example:"2025-10-09T15:04:05Z"`
}

type DataValidationDistribution struct {
	Quantity float64   `json:"quantity" gorm:"column:quantity"`
	Time     time.Time `json:"time" gorm:"column:create_time"`
}

type UpdateStatusDistributionRequest struct {
	DistributionId       uint  `json:"-" validate:"required"`
	DistributorProfileId uint  `json:"-" validate:"required"`
	Status               int32 `json:"status" validate:"required;oneof=1 2 3 4 5 6 7" example:"7"`
}
