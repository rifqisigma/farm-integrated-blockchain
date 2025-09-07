package dto

import "time"

type CreateDistributionRequest struct {
	FarmerProfileId uint    `json:"-" validate:"required"`
	HarvestId       uint    `json:"-" validate:"required"`
	Quantity        float64 `json:"quantity" validate:"required"`
	MarkupPrice     float64 `json:"markup_price" validate:"required"`
	FinalPrice      float64 `json:"final_price" validate:"required"`
}

type UpdateDistributionRequest struct {
	DistributionId       uint    `json:"-" `
	DistributorProfileId uint    `json:"-"`
	Quantity             float64 `json:"quantity" `
	MarkupPrice          float64 `json:"markup_price" `
	FinalPrice           float64 `json:"final_price" `
}

type GetDistribution struct {
	DistributionId  uint      `json:"-"  gorm:"id"`
	DistributorName string    `json:"distributor_name"  gorm:"distributor_name"`
	FarmerName      string    `json:"farmer_name" gorm:"farmer_name"`
	CropName        string    `json:"crop_name" gorm:"crop_name"`
	FinalPrice      float64   `json:"final_price" gorm:"final_price"`
	Time            time.Time `json:"time" gorm:" time"`
}

type GetDistributionById struct {
	DistributionId  uint      `json:"id"  gorm:"id"`
	DistributorName string    `json:"distributor_name"  gorm:"distributor_name"`
	FarmerName      string    `json:"farmer_name" gorm:"farmer_name"`
	CropName        string    `json:"crop_name" gorm:"crop_name"`
	FinalPrice      float64   `json:"final_price" gorm:"final_price"`
	BlockHash       string    `json:"block_hash" gorm:"block_hash"`
	HasArrived      bool      `json:"has_arrived" gorm:"has_arrived"`
	Time            time.Time `json:"time" gorm:" time"`
}

type DataValidationDistribution struct {
	Quantity float64   `json:"quantity" gorm:"quantity"`
	Time     time.Time `json:"time" gorm:"create_time"`
}

type UpdateStatusDistributionRequest struct {
	DistributionId       uint  `json:"-" validate:"required"`
	DistributorProfileId uint  `json:"-" validate:"required"`
	Status               int32 `json:"status" validate:"required;oneof=1 2 3 4 5 6 7"`
}
