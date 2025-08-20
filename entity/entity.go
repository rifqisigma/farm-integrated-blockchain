package entity

import "time"

type Status string

const (
	Consumer    Status = "consumer"
	Farmer      Status = "farmer"
	Distributor Status = "distributor"
	Seller      Status = "seller"
)

type User struct {
	ID         uint      `gorm:"primaryKey"`
	Email      string    `gorm:"not null:unique"`
	Password   string    `gorm:"not null"`
	IsVerified bool      `gorm:"default:false"`
	Provider   string    `gorm:"default:gmail"`
	Status     Status    `gorm:"type:enum('consumer','farmer','distributor','seller')"`
	CreateTime time.Time `gorm:"autoCreateTime"`
}

type FarmerProfile struct {
	ID       uint `gorm:"primaryKey"`
	UserID   uint `gorm:"index"`
	FarmName string
	Location string
	Harvest  []Harvest `gorm:"foreignKey:FarmerProfileId:constraint:OnDelete:CASCADE"`
}

type DistributorProfile struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint `gorm:"index"`
	Warehouse    string
	Region       string
	Distribution []Distribution `gorm:"foreignKey:DistributorProfileId:constraint:OnDelete:CASCADE"`
}

type RetailerProfile struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint `gorm:"index"`
	ShopName     string
	Market       string
	RetailerCart []RetailerCart `gorm:"foreignKey:RetailerId:constraint:OnDelete:CASCADE"`
}

type Crops struct {
	ID    uint   `gorm:"primaryKey"`
	Crops string `gorm:"not null"`
}

type Harvest struct {
	ID              uint      `gorm:"primaryKey"`
	FarmerProfileId uint      `gorm:"index"`
	CropId          uint      `gorm:"index"`
	Quantity        float64   `gorm:"not null"`
	BasePrice       float64   `gorm:"not null"`
	Accept          bool      `gorm:"default:false"`
	CreateTime      time.Time `gorm:"autoCreateTime"`

	//relation
	Crops Crops `gorm:"foreignKey:ChopId:constraint:OnDelete:CASCADE"`
}

type Distribution struct {
	ID                   uint      `gorm:"primaryKey"`
	HarvestId            uint      `gorm:"index"`
	DistributorProfileId uint      `gorm:"index"`
	ApprovedByFarmer     bool      `gorm:"default:false"`
	MarkUpPrice          float64   `gorm:"not null"`
	FinalPrice           float64   `gorm:"not null"`
	BlockHash            string    `gorm:"not null"`
	CreateTime           time.Time `gorm:"autoCreateTime"`
	HasArrived           bool      `gorm:"default:false"`

	//relation
	Harvest Harvest `gorm:"foreignKey:HarvestId:constraint:OnDelete:CASCADE"`
}

type RetailerCart struct {
	ID         uint      `gorm:"primaryKey"`
	RetailerId uint      `gorm:"index"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	IsReceived bool      `gorm:"default:false"`
}
