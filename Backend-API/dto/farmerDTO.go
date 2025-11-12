package dto

import "time"

type HarvestCreate struct {
	FarmerProfileId uint    `json:"-" validate:"required"`
	Name            string  `json:"name" validate:"required" example:"Jagung ke-1"`
	Quantity        float64 `json:"quantity" validate:"required" example:"100"`
	BasePrice       float64 `json:"base_price" validate:"required" example:"200000000"`
	Desc            string  `json:"desc" example:"this is fruit form Bogor "`
	CountryRequest  `json:"country" validate:"required"`
	CropCreate      `json:"crop" validate:"required"`
}

type HarvestUpdate struct {
	FarmerProfileId uint    `json:"-" validate:"required"`
	HarvestId       uint    `json:"-" validate:"required"`
	Name            string  `json:"name"  example:"Jagung ke-1"`
	Quantity        float64 `json:"quantity" validate:"omitempty,gte=0.1" example:"100"`
	BasePrice       float64 `json:"base_price" gorm:"base_price" example:"200000000"`
	Desc            string  `json:"desc" example:"this is fruit form Bogor "`
	CountryRequest  `json:"country"`
}

//get
type GetListHarvest struct {
	Id          uint      `json:"id" gorm:"column:id" example:"1"`
	Name        string    `json:"name" gorm:"column:name" example:"coffe ke-908"`
	CropName    string    `json:"crop_name" gorm:"column:crop_name" example:"rice"`
	FarmerName  string    `json:"farmer_name" gorm:"column:farmer_name" example:"Kartosuryo"`
	BasePrice   float64   `json:"base_price" gorm:"column:base_price" example:"100000000"`
	RegencyName string    `json:"regency_name" gorm:"column:regency_name" example:"Bogor"`
	Status      int16     `json:"status" gorm:"column:status" example:"1"`
	Quantity    float64   `json:"quantity" gorm:"column:quantity" example:"100"`
	Time        time.Time `json:"time" gorm:"column:time" example:"2025-10-09T15:04:05Z"`
}

type GetHarvestById struct {
	Id          uint    `json:"id" gorm:"column:id" example:"192"`
	CropName    string  `json:"crop_name" gorm:"column:crop_name" example:"Corn"`
	Commodity   string  `json:"commodity" gorm:"column:commodity" example:"Yellow Corn"`
	FarmerName  string  `json:"farmer_name" gorm:"column:farmer_name" example:"Wayulo"`
	Name        string  `json:"name" gorm:"column:name" example:"coffe ke-908"`
	Desc        string  `json:"desc" gorm:"column:description" example:"this is coffee beans form central of java"`
	BasePrice   float64 `json:"base_price" gorm:"column:base_price" example:"100000000"`
	Quantity    float64 `json:"quantity" gorm:"column:quantity" example:"100"`
	TxBlock     string  `json:"tx_block" gorm:"column:tx_block" example:""`
	RegencyName string  `json:"regency_name" gorm:"column:regency_name" example:"Bogor"`
	RegionName  string  `json:"region_name" gorm:"column:region_name" example:"Jawa barat"`
	Status      int16   `json:"status" gorm:"column:status" example:"1"`

	Time time.Time `json:"time" gorm:"column:time" example:"2025-10-09T15:04:05Z"`
}

type CropCreate struct {
	Commodity       string `json:"commodity" validate:"required" example:"Beras"`
	Name            string `json:"name" validate:"required" example:"padi"`
	HarvestTimeSpan int    `json:"harvest_time_span" example:"120"`
}

type BcHarvest struct {
	ID          int64  `json:"id"`
	FarmerID    int64  `json:"farmer_profile_id"`
	CropID      int64  `json:"crop_id"`
	RegencyID   int64  `json:"regency_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
	BasePrice   int64  `json:"base_price"`
}
