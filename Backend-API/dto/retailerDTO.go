package dto

import "time"

type CreateRetailerCartRequest struct {
	RetailerProfileId uint    `json:"-" gorm:"retailer_profile_id" validate:"required" example:"324"`
	DistributionId    uint    `json:"-" gorm:"distribution_id" validate:"required" example:"765"`
	Quantity          float64 `json:"quantity" validate:"required" example:"600.50"`
}

type UpdateRetailerCartRequest struct {
	DistributionId    uint    `json:"-" gorm:"distribution_id" validate:"required" `
	RetailerCartId    uint    `json:"-" gorm:"id" validate:"required"`
	RetailerProfileId uint    `json:"-" gorm:"retailer_profile_id" validate:"required"`
	Quantity          float64 `json:"quantity" gorm:"quantity" validate:"required" example:"34.00"`
}

type DataValidationRetailerCart struct {
	Quantity   float64 `json:"quantity" gorm:"quantity"`
	IsCanceled bool    `json:"is_canceled" gorm:"is_canceled"`
}

type GetRetailerCart struct {
	ID              uint      `json:"id" gorm:"id"`
	DistributorName string    `json:"distributor_name" gorm:"distributor_name" example:"Gopher"`
	RetailerName    string    `json:"retailer_name" gorm:"retailer_name" example:"Siti"`
	HarvestName     string    `json:"harvest_name" gorm:"hatvest_name" example:""`
	Quantity        float64   `json:"quantity" gorm:"quantity" example:"35.320"`
	UpdateTime      time.Time `json:"time" gorm:"time" example:"2025-10-09T15:04:05Z"`
}

type GetRetailerCartById struct {
	ID              uint      `json:"id" gorm:"id"`
	DistributorName string    `json:"distributor_name" gorm:"distributor_name" example:"Valen"`
	FarmerName      string    `json:"farmer_name" gorm:"farmer_name" example:"Wowo"`
	HarvestName     string    `json:"harvest_name" gorm:"hatvest_name" example:""`
	Quantity        float64   `json:"quantity" gorm:"quantity" example:"90.20"`
	BlockHash       string    `json:"block_hash" gorm:"not null" example:"-"`
	UpdateTime      time.Time `json:"time" gorm:"time" example:"2025-10-09T15:04:05Z"`
}

type ApprovedRetailerCart struct {
	RetailerCartId       uint `json:"-" validate:"required"`
	DistributorProfileId uint `json:"-" validate:"required"`
	Approved             bool `json:"approved" validate:"required" example:"true"`
}

type DataValidationRetailer struct {
	Quantity float64   `json:"quantity" gorm:"quantity"`
	Time     time.Time `json:"time" gorm:"create_time"`
}
