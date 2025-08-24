package entity

import "time"

type Status string

const (
	Consumer    Status = "consumer"
	Farmer      Status = "farmer"
	Distributor Status = "distributor"
	Retailer    Status = "retailer"
)

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Email      string `gorm:"not null;unique"`
	Password   string `gorm:"not null"`
	IsVerified bool   `gorm:"default:false"`
	Provider   string `gorm:"default:'gmail'"`
	Role       Status `gorm:"type:enum('consumer','farmer','distributor','retailer');not null"`

	CreateTime time.Time `gorm:"autoCreateTime"`
}

type Token struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"index;not null;unique"`
	Token     string `gorm:"not null"`
	IsRevoked bool   `gorm:"default:false"`

	CreateTime time.Time `gorm:"autoCreateTime"`
}

type FarmerProfile struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint `gorm:"index"`
	Name    string
	Harvest []Harvest `gorm:"foreignKey:FarmerProfileId;constraint:OnDelete:CASCADE"`

	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}

type DistributorProfile struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint `gorm:"index"`
	Name         string
	Distribution []Distribution `gorm:"foreignKey:DistributorProfileId;constraint:OnDelete:CASCADE"`

	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}

type RetailerProfile struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint `gorm:"index"`
	Name         string
	RetailerCart []RetailerCart `gorm:"foreignKey:RetailerId;constraint:OnDelete:CASCADE"`

	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}

type ConsumerProfile struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `gorm:"index"`
	Name   string `gorm:"not null"`

	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
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
	UpdateTime      time.Time `gorm:"autoUpdateTime"`

	//relation
	Crops Crops `gorm:"foreignKey:CropId;constraint:OnDelete:CASCADE"`
}

type Distribution struct {
	ID                   uint      `gorm:"primaryKey"`
	HarvestId            uint      `gorm:"index"`
	DistributorProfileId uint      `gorm:"index"`
	ApprovedByFarmer     bool      `gorm:"default:false"`
	MarkUpPrice          float64   `gorm:"not null"`
	FinalPrice           float64   `gorm:"not null"`
	BlockHash            string    `gorm:"not null"`
	HasArrived           bool      `gorm:"default:false"`
	CreateTime           time.Time `gorm:"autoCreateTime"`
	UpdateTime           time.Time `gorm:"autoUpdateTime"`

	//relation
	Harvest Harvest `gorm:"foreignKey:HarvestId;constraint:OnDelete:CASCADE"`
}

type RetailerCart struct {
	ID         uint      `gorm:"primaryKey"`
	RetailerId uint      `gorm:"index"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
	IsReceived bool      `gorm:"default:false"`
}
