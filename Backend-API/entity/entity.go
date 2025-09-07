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
	IsDeleted  bool   `gorm:"default:false"`

	CreateTime time.Time `gorm:"autoCreateTime"`
}

type Token struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"index;not null"`
	Token     string `gorm:"not null"`
	IsRevoked bool   `gorm:"default:false"`

	CreateTime time.Time `gorm:"autoCreateTime"`
}

type FarmerProfile struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint `gorm:"index;unique"`
	Name       string
	Harvest    []Harvest `gorm:"foreignKey:FarmerProfileId"`
	IsDeleted  bool      `gorm:"default:false"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}

type DistributorProfile struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint `gorm:"index;unique"`
	Name         string
	Distribution []Distribution `gorm:"foreignKey:DistributorProfileId"`
	IsDeleted    bool           `gorm:"default:false"`
	CreateTime   time.Time      `gorm:"autoCreateTime"`
	UpdateTime   time.Time      `gorm:"autoUpdateTime"`
}

type RetailerProfile struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint `gorm:"index;unique"`
	Name         string
	RetailerCart []RetailerCart `gorm:"foreignKey:RetailerId"`
	IsDeleted    bool           `gorm:"default:false"`
	CreateTime   time.Time      `gorm:"autoCreateTime"`
	UpdateTime   time.Time      `gorm:"autoUpdateTime"`
}

type ConsumerProfile struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"index;unique"`
	Name       string    `gorm:"not null"`
	IsDeleted  bool      `gorm:"default:false"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}

type Crop struct {
	ID              uint      `gorm:"primaryKey"`
	Crop            string    `gorm:"not null"`
	HarvestTimeSpan int       `gorm:"not null"`
	CreateTime      time.Time `gorm:"autoCreateTime"`
	IsDeleted       bool      `gorm:"default:false"`
	Harvest         []Harvest `gorm:"foreignKey:CropId"`
}

type Harvest struct {
	ID              uint      `gorm:"primaryKey"`
	Name            string    `gorm:"not null"`
	FarmerProfileId uint      `gorm:"index;not null"`
	CropId          uint      `gorm:"index;not null"`
	Quantity        float64   `gorm:"not null"`
	BasePrice       float64   `gorm:"not null"`
	BlockHash       string    `gorm:"not null"`
	CreateTime      time.Time `gorm:"autoCreateTime"`
	UpdateTime      time.Time `gorm:"autoUpdateTime"`
	IsCanceled      bool      `gorm:"default:false"`
	//relation
	FarmerProfile FarmerProfile `gorm:"foreignKey:FarmerProfileId"`
	Crops         Crop          `gorm:"foreignKey:CropId"`
}

type Distribution struct {
	ID                   uint      `gorm:"primaryKey"`
	Name                 string    `gorm:"not null"`
	HarvestId            uint      `gorm:"index;not null"`
	DistributorProfileId uint      `gorm:"index;not null"`
	FarmerProfileId      uint      `gorm:"index;not null"`
	ApprovedByFarmer     bool      `gorm:"default:false"`
	MarkUpPrice          float64   `gorm:"not null"`
	FinalPrice           float64   `gorm:"not null"`
	Quantity             float64   `gorm:"not null"`
	BlockHash            string    `gorm:"not null"`
	CreateTime           time.Time `gorm:"autoCreateTime"`
	UpdateTime           time.Time `gorm:"autoUpdateTime"`
	StatusDistribution   int32     `gorm:"not null" validate:"min=1,max=7"`
	IsCanceled           bool      `gorm:"default:false"`
	//relation
	Harvest            Harvest            `gorm:"foreignKey:HarvestId"`
	DistributorProfile DistributorProfile `gorm:"foreignKey:DistributorProfileId"`
	FarmerProfile      FarmerProfile      `gorm:"foreignKey:FarmerProfileId"`
}

type RetailerCart struct {
	ID                    uint      `gorm:"primaryKey"`
	Name                  string    `gorm:"not null"`
	RetailerProfileId     uint      `gorm:"index;not null"`
	DistributorProfileId  uint      `gorm:"index;not null"`
	DistributionId        uint      `gorm:"index;not null"`
	BlockHash             string    `gorm:"not null"`
	ApprovedByDistributor bool      `gorm:"default:false"`
	IsReceived            bool      `gorm:"default:false"`
	Quantity              float64   `gorm:"not null"`
	CreateTime            time.Time `gorm:"autoCreateTime"`
	UpdateTime            time.Time `gorm:"autoUpdateTime"`
	IsCanceled            bool      `gorm:"default:false"`

	DistributorProfile DistributorProfile `gorm:"foreignKey:DistributorProfileId"`
	Distribution       Distribution       `gorm:"foreignKey:DistributionId"`
	RetailerProfile    RetailerProfile    `gorm:"foreignKey:RetailerId"`
}
