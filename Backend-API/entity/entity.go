package entity

import "time"

type Status string

const (
	Consumer    Status = "consumer"
	Farmer      Status = "farmer"
	Distributor Status = "distributor"
	Retailer    Status = "retailer"
	Admin       Status = "admin"
)

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Email      string `gorm:"not null;unique"`
	Password   string `gorm:"not null"`
	IsVerified bool   `gorm:"default:false"`
	Provider   string `gorm:"default:'gmail'"`
	Role       Status `gorm:"type:enum('consumer','farmer','distributor','retailer','admin');not null"`
	IsDeleted  bool   `gorm:"default:false"`

	CreateTime time.Time `gorm:"autoCreateTime"`
}

type Token struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"index;not null"`
	Token     string `gorm:"not null"`
	IsRevoked bool   `gorm:"default:false"`

	CreateTime time.Time `gorm:"autoCreateTime"`
	User       User      `foreignKey:"UserId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

type FarmerProfile struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint `gorm:"index;unique"`
	Name       string
	Harvest    []Harvest `gorm:"foreignKey:FarmerProfileId"`
	IsDeleted  bool      `gorm:"default:false"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`

	User User `foreignKey:"UserId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

type DistributorProfile struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint `gorm:"index;unique"`
	Name         string
	Distribution []Distribution `gorm:"foreignKey:DistributorProfileId"`
	IsDeleted    bool           `gorm:"default:false"`
	CreateTime   time.Time      `gorm:"autoCreateTime"`
	UpdateTime   time.Time      `gorm:"autoUpdateTime"`

	User User `foreignKey:"UserId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

type RetailerProfile struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint `gorm:"index;unique"`
	Name         string
	RetailerCart []RetailerCart `gorm:"foreignKey:RetailerProfileId"`
	IsDeleted    bool           `gorm:"default:false"`
	CreateTime   time.Time      `gorm:"autoCreateTime"`
	UpdateTime   time.Time      `gorm:"autoUpdateTime"`

	User User `foreignKey:"UserId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
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
	Harvest         []Harvest `gorm:"foreignKey:CropId;"`
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
	RegencyId       uint      `gorm:"index; not null"`

	//relation
	FarmerProfile FarmerProfile `gorm:"foreignKey:FarmerProfileId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Crop          Crop          `gorm:"foreignKey:CropId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Regency       Regency       `gorm:"foreignKey:RegencyId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
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
	DestinationId        uint      `gorm:"index; not null"`
	//relation
	Harvest            Harvest            `gorm:"foreignKey:HarvestId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	DistributorProfile DistributorProfile `gorm:"foreignKey:DistributorProfileId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	FarmerProfile      FarmerProfile      `gorm:"foreignKey:FarmerProfileId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Destination        Regency            `gorm:"foreignKey:DestinationId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
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

	//relasi
	DistributorProfile DistributorProfile `gorm:"foreignKey:DistributorProfileId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Distribution       Distribution       `gorm:"foreignKey:DistributionId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	RetailerProfile    RetailerProfile    `gorm:"foreignKey:RetailerProfileId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

type Country struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
	// Relasi
	Regions []Region
}

type Region struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CountryID uint   `gorm:"not null"`
	Country   Country
	// Relasi
	Regencies []Regency
}

type Regency struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	RegionID uint   `gorm:"not null"`
	Region   Region
}
