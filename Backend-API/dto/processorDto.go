package dto

import "time"

type CreateProcessor struct {
	ProcessorProfileId uint    `json:"-" validate:"required"`
	HarvestCollectorId uint    `json:"harvest_collector_id" example:"297083"`
	HarvestId          uint    `json:"harvest_id" example:"78321"`
	Name               string  `json:"name" validate:"required"`
	Desc               string  `json:"desc"`
	Quantity           float64 `json:"quantity" validate:"required,gte=0.1"`
	BasePrice          float64 `json:"base_price" validate:"required,gte=0.1"`
	Price              float64 `json:"price" validate:"required,gte=0.1"`
}

type UpdateProcessor struct {
	ProcessorProfileId uint    `json:"-" validate:"required"`
	ProcessorHarvestId uint    `json:"-" validate:"required"`
	Name               string  `json:"name" `
	Desc               string  `json:"desc"`
	Quantity           float64 `json:"quantity" validate:"omitempty,gte=0.1"`
	BasePrice          float64 `json:"base_price" validate:"omitempty,gte=0.1"`
	Price              float64 `json:"price" validate:"omitempty,gte=0.1"`
}

type GetListHarvestProcessor struct {
	ProcessorHarvestId   uint      `json:"id" gorm:"column:id"`
	ProcessorProfileName string    `json:"processor_profile_name"  gorm:"column:processor_profile_name" `
	CropName             string    `json:"crop_name"  gorm:"column:crop_name"`
	Name                 string    `json:"name"  gorm:"column:name"`
	Quantity             float64   `json:"quantity"  gorm:"column:quantity"`
	BasePrice            float64   `json:"base_price"  gorm:"column:base_price"`
	Price                float64   `json:"price"  gorm:"column:price"`
	Status               int16     `json:"status" gorm:"column:status" example:"1"`
	Time                 time.Time `json:"time" gorm:"column:time" example:"2025-10-09T15:04:05Z"`
}

type GetHarvestProcessorById struct {
	ProcessorHarvestId   uint      `json:"id" gorm:"column:id" example:"890"`
	ProcessorProfileName string    `json:"processor_profile_name"  gorm:"column:processor_profile_name" example:"Waluyo"`
	CropName             string    `json:"crop_name"  gorm:"column:crop_name" example:"Waluyo"`
	Name                 string    `json:"name"  gorm:"column:name"`
	Desc                 string    `json:"desc" gorm:"column:description" example:"This rice is processed into packaged rice"`
	Quantity             float64   `json:"quantity"  gorm:"column:quantity" example:"5.0"`
	BasePrice            float64   `json:"base_price"  gorm:"column:base_price" example:"20000"`
	Price                float64   `json:"price"  gorm:"column:price" example:"40000"`
	TxBlock              string    `json:"tx_block" gorm:"column:tx_block" example:""`
	Status               int16     `json:"status" gorm:"column:status" example:"1"`
	Time                 time.Time `json:"time" gorm:"column:time" example:"2025-10-09T15:04:05Z"`
}

type BcHarvestProcessor struct {
	ID                 int64  `json:"id"`
	ProcessorID        int64  `json:"processor_profile_id"`
	HarvestCollectorID int64  `json:"harvest_collector_id"`
	HarvestID          int64  `json:"harvest_id"`
	Name               string `json:"name"`
	Desc               string `json:"desc"`
	Quantity           int64  `json:"quantity"`
	BasePrice          int64  `json:"base_price"`
	Price              int64  `json:"price"`
}
