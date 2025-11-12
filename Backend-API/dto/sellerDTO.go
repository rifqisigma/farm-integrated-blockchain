package dto

import "time"

type CreateSellerBox struct {
	SellerProfileId uint    `json:"-" gorm:"seller_profile_id" validate:"required" example:"324"`
	DistributionId  uint    `json:"-" gorm:"distribution_id" validate:"required" example:"765"`
	Name            string  `json:"name" validate:"required" example:"distribution-coffe 897"`
	Desc            string  `json:"desc" validate:"required" example:"coffe batak to jakarta"`
	Quantity        float64 `json:"quantity" validate:"required" example:"200.50"`
	BasePrice       float64 `json:"base_price" validate:"required" example:"20.000"`
	Price           float64 `json:"price" validate:"required" example:"25.000"`
}

type UpdateSellerBox struct {
	SellerBoxId     uint    `json:"-" gorm:"id" validate:"required"`
	SellerProfileId uint    `json:"-" gorm:"seller_profile_id" validate:"required"`
	Name            string  `json:"name"  example:"distribution-coffe 897"`
	Desc            string  `json:"desc"  example:"coffe batak to jakarta"`
	Quantity        float64 `json:"quantity" validate:"omitempty,gte=0.1" example:"200.50"`
	BasePrice       float64 `json:"base_price" validate:"omitempty,gte=0.1" example:"20.000"`
	Price           float64 `json:"price" example:"25.000"`
}

type GetSellerBox struct {
	ID         uint      `json:"id" gorm:"column:id"`
	SellerName string    `json:"seller_name" gorm:"seller_name" example:"Wowo"`
	CropName   string    `json:"crop_name" gorm:"crop_name" example:"coffe"`
	Name       string    `json:"name" gorm:"column:name" example:"coffe beans"`
	BasePrice  float64   `json:"base_price" gorm:"column:base_price" example:"89000"`
	Price      float64   `json:"price" validate:"required" example:"25.000"`
	Quantity   float64   `json:"quantity" gorm:"quantity" example:"90.20"`
	UpdateTime time.Time `json:"time" gorm:"time" example:"2025-10-09T15:04:05Z"`
}

type GetSellerBoxById struct {
	ID         uint      `json:"id" gorm:"column:id"`
	SellerName string    `json:"seller_name" gorm:"seller_name" example:"Wowo"`
	CropName   string    `json:"crop_name" gorm:"crop_name" example:"coffe"`
	Name       string    `json:"name" gorm:"column:name" example:"coffe beans"`
	Desc       string    `json:"desc" gorm:"column:description" example:"coffe beans from tegal"`
	Price      float64   `json:"price" validate:"required" example:"25.000"`
	BasePrice  float64   `json:"base_price" gorm:"column:base_price" example:"89000"`
	Quantity   float64   `json:"quantity" gorm:"quantity" example:"90.20"`
	Status     int       `json:"status" gorm:"column:status" example:"1"`
	TxBlock    string    `json:"tx_block" gorm:"column:tx_block" example:""`
	UpdateTime time.Time `json:"time" gorm:"time" example:"2025-10-09T15:04:05Z"`
}

type BcSellerBox struct {
	ID             int64  `json:"id"`
	SellerID       int64  `json:"seller_profile_id"`
	DistributionID int64  `json:"distribution_id"`
	Name           string `json:"name"`
	Desc           string `json:"desc"`
	Quantity       int64  `json:"quantity"`
	BasePrice      int64  `json:"base_price"`
	Price          int64  `json:"price"`
}
