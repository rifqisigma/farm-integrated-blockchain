package dto

import "time"

type CreateDistribution struct {
	DistributorProfileId uint    `json:"-" validate:"required" `
	HarvestCollectorId   uint    `json:"harvest_collector_id" example:"13"`
	HarvestProcessorId   uint    `json:"harvest_processor_id" example:"54"`
	HarvestId            uint    `json:"harvest_id" example:"13"`
	Name                 string  `json:"name" validate:"required" example:"distribution-coffe 897"`
	Desc                 string  `json:"desc" example:"coffe batak to jakarta"`
	Transportation       string  `json:"transportation" validate:"required" example:"air"`
	Quantity             float64 `json:"quantity" validate:"required,gte=0.1" example:"200.50"`
	BasePrice            float64 `json:"base_price" validate:"required,gte=0.1" example:"20.000"`
	Price                float64 `json:"price" validate:"required,gte=0.1" example:"220.000"`
	CountryRequest       `json:"country" validate:"required"`
}

type UpdateDistribution struct {
	DistributionId       uint    `json:"-" `
	DistributorProfileId uint    `json:"-"`
	Name                 string  `json:"name"  example:"distribution-coffe 897"`
	Desc                 string  `json:"desc" example:"coffe batak to jakarta"`
	Transportation       string  `json:"transportation" example:"air"`
	Quantity             float64 `json:"quantity" validate:"omitempty,gte=0.1" example:"20.70"`
	BasePrice            float64 `json:"base_price" validate:"omitempty,gte=0.1" example:"20.000"`
	Price                float64 `json:"price" validate:"omitempty,gte=0.1"  example:"220.000"`
	CountryRequest       `json:"country"`
}

type GetListDistribution struct {
	DistributionId  uint      `json:"id"  gorm:"column:id"`
	DistributorName string    `json:"distributor_name"  gorm:"column:distributor_name"`
	CropName        string    `json:"crop_name" gorm:"column:crop_name"`
	Name            string    `json:"name" gorm:"name" example:"distribution-coffe 897"`
	Transportation  string    `json:"transportation"  gorm:"column:transportation"   example:"air"`
	Quantity        float64   `json:"quantity"  gorm:"column:quantity"  example:"20.70"`
	BasePrice       float64   `json:"base_price"  gorm:"column:base_price" example:"20.000"`
	Price           float64   `json:"price" gorm:"column:price" example:"220.000"`
	RegencyName     string    `json:"destination_regency" gorm:"column:regency_name" example:"Bogor"`
	Status          int16     `json:"status" gorm:"column:status" example:"1"`
	Time            time.Time `json:"time" gorm:"column:time" example:"2025-10-09T15:04:05Z"`
}

type GetDistributionById struct {
	DistributionId     uint      `json:"id"  gorm:"column:id" example:"23324"`
	CropName           string    `json:"crop_name" gorm:"column:crop_name" example:"Corn"`
	Name               string    `json:"name" gorm:"name" example:"distribution-coffe 897"`
	Desc               string    `json:"desc" gorm:"column:description" example:"coffe batak to jakarta"`
	Transportation     string    `json:"transportation"  gorm:"column:transportation"   example:"air"`
	Quantity           float64   `json:"quantity"  gorm:"column:quantity"  example:"20.70"`
	BasePrice          float64   `json:"base_price"  gorm:"column:base_price" example:"20.000"`
	Price              float64   `json:"price" gorm:"column:price" example:"220.000"`
	TxBlock            string    `json:"tx_block" gorm:"column:tx_block" example:""`
	Status             int16     `json:"status" gorm:"column:status" example:"1"`
	DistributionStatus int16     `json:"distribution_status" gorm:"column:distribution_status" example:"5"`
	RegencyName        string    `json:"destination_regency" gorm:"column:regency_name" example:"Bogor"`
	RegionName         string    `json:"destination_region" gorm:"column:region_name" example:"Jawa barat"`
	Time               time.Time `json:"time" gorm:"column:time" example:"2025-10-09T15:04:05Z"`
}

type UpdateStatusDistribution struct {
	DistributionId       uint  `json:"-" validate:"required"`
	DistributorProfileId uint  `json:"-" validate:"required"`
	Status               int32 `json:"status" validate:"required,oneof=1 2 3 4 5 6 7" example:"7"`
}

type BcDistribution struct {
	ID               int64  `json:"id"`
	DistributorID    int64  `json:"distributor_profile_id"`
	DestinationID    int64  `json:"destination_id"`
	HarvestID        int64  `json:"harvest_id"`
	HarvestCollector int64  `json:"harvest_collector_id"`
	HarvestProcessor int64  `json:"harvest_processor_id"`
	Name             string `json:"name"`
	Desc             string `json:"desc"`
	Quantity         int64  `json:"quantity"`
	BasePrice        int64  `json:"base_price"`
	Price            int64  `json:"price"`
	Transportation   string `json:"transportation"`
}
