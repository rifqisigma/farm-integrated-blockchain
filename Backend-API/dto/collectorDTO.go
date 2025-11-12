package dto

import "time"

type CreateCollector struct {
	CollectorProfileId uint    `json:"-" validate:"required" example:"122"`
	HarvestId          uint    `json:"-" validate:"required" example:"211"`
	Name               string  `json:"name" validate:"required" example:"G_Beans"`
	Desc               string  `json:"desc" example:"harvest coffe beans will save in G_Beans 8975"`
	Quantity           float64 `json:"quantity" validate:"required,gte=0.1" example:"200.50"`
	BasePrice          float64 `json:"base_price" validate:"required,gte=0.1" example:"20.000"`
	Price              float64 `json:"price" validate:"required,gte=0.1" example:"220.000"`
}

type UpdateCollector struct {
	CollectorId        uint `json:"-" validate:"required" example:"122"`
	CollectorProfileId uint `json:"-" validate:"required" example:"122"`

	Name      string  `json:"name"  example:"G_Beans"`
	Desc      string  `json:"desc" example:"harvest coffe beans will save in G_Beans 8975"`
	Quantity  float64 `json:"quantity" validate:"omitempty,gte=0.1" example:"200.50"`
	BasePrice float64 `json:"base_price" validate:"omitempty,gte=0.1" example:"20.000"`
	Price     float64 `json:"price" validate:"omitempty,gte=0.1" example:"220.000"`
}

type GetListHarvestCollector struct {
	CollectorId          uint      `json:"collector_id" gorm:"column:id" example:"122"`
	CollectorProfileName string    `json:"collector_profile_name" gorm:"column:collector_profile_name"   example:"122"`
	CropName             string    `json:"crop_name" gorm:"column:crop_name" example:"3083"`
	Name                 string    `json:"name" gorm:"column:name"  example:"G_Beans"`
	Quantity             float64   `json:"quantity" gorm:"column:quantity"  example:"200.50"`
	BasePrice            float64   `json:"base_price" gorm:"column:base_price"  example:"20.000"`
	Price                float64   `json:"price" gorm:"column:price"  example:"220.000"`
	Status               int16     `json:"status" gorm:"column:status" example:"1"`
	Time                 time.Time `json:"time" gorm:"column:time" example:"2024-06-01T15:04:05Z"`
}

type GetHarvestCollectorById struct {
	CollectorId          uint      `json:"collector_id" gorm:"column:id" example:"122"`
	CollectorProfileName string    `json:"collector_profile_name" gorm:"column:collector_profile_name"   example:"122"`
	CropName             string    `json:"crop_name" gorm:"column:crop_name" example:"3083"`
	Name                 string    `json:"name" gorm:"column:name"  example:"G_Beans"`
	Desc                 string    `json:"desc" gorm:"column:description" example:"harvest coffe beans will save in G_Beans 8975"`
	Quantity             float64   `json:"quantity" gorm:"column:quantity"  example:"200.50"`
	BasePrice            float64   `json:"base_price" gorm:"column:base_price"  example:"20.000"`
	Price                float64   `json:"price" gorm:"column:price"  example:"220.000"`
	Status               int16     `json:"status" gorm:"column:status" example:"1"`
	TxBlock              string    `json:"tx_block" gorm:"column:tx_block" example:""`
	Time                 time.Time `json:"time" gorm:"column:time" example:"2024-06-01T15:04:05Z"`
}

type BcHarvestCollector struct {
	ID          int64  `json:"id"`
	CollectorID int64  `json:"collector_profile_id"`
	HarvestID   int64  `json:"harvest_id"`
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	Quantity    int64  `json:"quantity"`
	Price       int64  `json:"price"`
	BasePrice   int64  `json:"base_price"`
}
