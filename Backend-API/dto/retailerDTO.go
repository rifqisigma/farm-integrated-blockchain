package dto

import "time"

type CreateRetailerCartRequest struct {
	RetailerProfileId uint    `json:"-" gorm:"retailer_profile_id" validate:"required"`
	DistributionId    uint    `json:"-" gorm:"distribution_id" validate:"required"`
	Quantity          float64 `json:"quantity" validate:"required"`
}

type UpdateRetailerCartRequest struct {
	DistributionId    uint    `json:"-" gorm:"distribution_id" validate:"required"`
	RetailerCartId    uint    `json:"-" gorm:"id" validate:"required"`
	RetailerProfileId uint    `json:"-" gorm:"retailer_profile_id" validate:"required"`
	Quantity          float64 `json:"quantity" gorm:"quantity" validate:"required"`
}

type DataValidationRetailerCart struct {
	Quantity   float64 `json:"quantity" gorm:"quantity"`
	IsCanceled bool    `json:"is_canceled" gorm:"is_canceled"`
}

type GetRetailerCart struct {
	ID              uint      `json:"id" gorm:"id"`
	DistributorName uint      `json:"distributor_name" gorm:"distributor_name"`
	RetailerName    string    `json:"retailer_name" gorm:"retailer_name"`
	HarvestName     uint      `json:"harvest_name" gorm:"hatvest_name"`
	Quantity        float64   `json:"quantity" gorm:"quantity"`
	UpdateTime      time.Time `json:"time" gorm:"time"`
}

type GetRetailerCartById struct {
	ID              uint      `json:"id" gorm:"id"`
	DistributorName uint      `json:"distributor_name" gorm:"distributor_name"`
	FarmerName      string    `json:"farmer_name" gorm:"farmer_name"`
	HarvestName     uint      `json:"harvest_name" gorm:"hatvest_name"`
	Quantity        float64   `json:"quantity" gorm:"quantity"`
	BlockHash       string    `gorm:"not null"`
	UpdateTime      time.Time `json:"time" gorm:"time"`
}

type ApprovedRetailerCart struct {
	RetailerCartId       uint `json:"-" validate:"required"`
	DistributorProfileId uint `json:"-" validate:"required"`
	Approved             bool `json:"approved" validate:"required"`
}

type DataValidationRetailer struct {
	Quantity float64   `json:"quantity" gorm:"quantity"`
	Time     time.Time `json:"time" gorm:"create_time"`
}
