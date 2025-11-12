package entity

import "time"

type Status string

const (
	None        Status = "none"
	Consumer    Status = "consumer"
	Farmer      Status = "farmer"
	Distributor Status = "distributor"
	Seller      Status = "seller"
	Processor   Status = "processor"
	Collector   Status = "collector"
	Admin       Status = "admin"
)

type User struct {
	ID         uint      `gorm:"primaryKey"`
	Email      string    `gorm:"not null;unique"`
	Password   string    `gorm:"not null"`
	IsVerified bool      `gorm:"default:false"`
	Provider   string    `gorm:"default:'gmail'"`
	Role       Status    `gorm:"type:enum('consumer','farmer','distributor','seller','processor','collector','admin', 'none');default:'none'"`
	IsDeleted  bool      `gorm:"default:false"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type Token struct {
	ID        uint   `gorm:"primaryKey"`
	UserId    uint   `gorm:"index;not null"`
	Token     string `gorm:"not null"`
	IsRevoked bool   `gorm:"default:false"`

	CreateTime time.Time `gorm:"autoCreateTime"`
	User       User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

// ----------------------------- PROFILES -----------------------------

type FarmerProfile struct {
	ID     uint `gorm:"primaryKey"`
	UserId uint `gorm:"uniqueIndex;not null"`
	Name   string

	IsDeleted bool `gorm:"default:false"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User     User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Harvests []Harvest `gorm:"foreignKey:FarmerProfileId"`
}

type DistributorProfile struct {
	ID     uint `gorm:"primaryKey"`
	UserId uint `gorm:"uniqueIndex;not null"`
	Name   string

	IsDeleted bool `gorm:"default:false"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User          User           `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Distributions []Distribution `gorm:"foreignKey:DistributorProfileId"`
}

type SellerProfile struct {
	ID     uint `gorm:"primaryKey"`
	UserId uint `gorm:"uniqueIndex;not null"`
	Name   string

	IsDeleted bool `gorm:"default:false"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User      User        `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	SellerBox []SellerBox `gorm:"foreignKey:SellerProfileId"`
}

type ConsumerProfile struct {
	ID     uint `gorm:"primaryKey"`
	UserId uint `gorm:"uniqueIndex;not null"`
	Name   string

	IsDeleted bool `gorm:"default:false"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type ProcessorProfile struct {
	ID     uint `gorm:"primaryKey"`
	UserId uint `gorm:"uniqueIndex;not null"`
	Name   string

	IsDeleted bool `gorm:"default:false"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User             User               `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	HarvestProcessor []HarvestProcessor `gorm:"foreignKey:ProcessorProfileId"`
}

type CollectorProfile struct {
	ID     uint `gorm:"primaryKey"`
	UserId uint `gorm:"uniqueIndex;not null"`
	Name   string

	IsDeleted bool `gorm:"default:false"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User             User               `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	HarvestCollector []HarvestCollector `gorm:"foreignKey:CollectorProfileId"`
}

// ----------------------------- MASTER DATA -----------------------------

type Crop struct {
	ID              uint   `gorm:"primaryKey"`
	Commodity       string `gorm:"not null"`
	CropName        string `gorm:"not null"`
	HarvestTimeSpan int    `gorm:"not null"`

	IsDeleted bool `gorm:"default:false"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Harvests []Harvest `gorm:"foreignKey:CropId"`
}

// ----------------------------- TRANSACTIONS / CARTS -----------------------------

type Harvest struct {
	ID uint `gorm:"primaryKey"`

	FarmerProfileId uint `gorm:"index"`
	CropId          uint `gorm:"index;not null"`
	RegencyId       uint `gorm:"index;not null"`

	Name      string `gorm:"not null"`
	Desc      string
	Quantity  float64 `gorm:"not null"`
	BasePrice float64 `gorm:"not null"`
	Status    int16   `gorm:"not null;default:0;check:status >= 0 AND status <= 2"`
	TxBlock   string

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	FarmerProfile FarmerProfile `gorm:"foreignKey:FarmerProfileId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Crop          Crop          `gorm:"foreignKey:CropId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Regency       Regency       `gorm:"foreignKey:RegencyId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type HarvestCollector struct {
	ID uint `gorm:"primaryKey"`

	CollectorProfileId uint `gorm:"index;not null"`
	HarvestId          uint `gorm:"index;not null"`

	Name      string  `gorm:"not null"`
	Desc      string  `gorm:"not null"`
	Quantity  float64 `gorm:"not null"`
	Price     float64 `gorm:"not null"`
	BasePrice float64 `gorm:"not null"`
	Status    int16   `gorm:"not null;default:0;check:status >= 0 AND status <= 2"`
	TxBlock   string

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	CollectorProfile CollectorProfile `gorm:"foreignKey:CollectorProfileId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Harvest          Harvest          `gorm:"foreignKey:HarvestId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type HarvestProcessor struct {
	ID uint `gorm:"primaryKey"`

	ProcessorProfileId uint `gorm:"index;not null"`

	HarvestCollectorId *uint `gorm:"index"`
	HarvestId          *uint `gorm:"index"`

	Name      string  `gorm:"not null"`
	Desc      string  `gorm:"not null"`
	Quantity  float64 `gorm:"not null"`
	BasePrice float64 `gorm:"not null"`
	Price     float64 `gorm:"not null"`
	Status    int16   `gorm:"not null;default:0;check:status >= 0 AND status <= 2"`
	TxBlock   string

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	ProcessorProfile ProcessorProfile  `gorm:"foreignKey:ProcessorProfileId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	HarvestCollector *HarvestCollector `gorm:"foreignKey:HarvestCollectorId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Harvest          *Harvest          `gorm:"foreignKey:HarvestId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type Distribution struct {
	ID uint `gorm:"primaryKey"`

	DistributorProfileId uint  `gorm:"index;not null"`
	DestinationId        uint  `gorm:"index;not null"`
	HarvestId            *uint `gorm:"index"`
	HarvestCollectorId   *uint `gorm:"index"`
	HarvestProcessorId   *uint `gorm:"index"`

	Name               string `gorm:"not null"`
	Desc               string
	Quantity           float64 `gorm:"not null"`
	BasePrice          float64 `gorm:"not null"`
	Price              float64 `gorm:"not null"`
	Transportation     string  `gorm:"type:varchar(20);not null"`
	DistributionStatus int16   `gorm:"not null;default:1;check:distribution_status >= 1 AND distribution_status <= 7"`
	Status             int16   `gorm:"not null;default:0;check:status >= 0 AND status <= 2"`
	TxBlock            string

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	DistributorProfile DistributorProfile `gorm:"foreignKey:DistributorProfileId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Destination        Regency            `gorm:"foreignKey:DestinationId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	Harvest          *Harvest          `gorm:"foreignKey:HarvestId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	HarvestCollector *HarvestCollector `gorm:"foreignKey:HarvestCollectorId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	HarvestProcessor *HarvestProcessor `gorm:"foreignKey:HarvestProcessorId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type SellerBox struct {
	ID uint `gorm:"primaryKey"`

	SellerProfileId uint `gorm:"index;not null"`
	DistributionId  uint `gorm:"index;not null"`

	Name      string  `gorm:"not null"`
	Desc      string  `gorm:"not null"`
	Quantity  float64 `gorm:"not null"`
	BasePrice float64 `gorm:"not null"`
	Price     float64 `gorm:"not null"`
	Status    int16   `gorm:"not null;default:0;check:status >= 0 AND status <= 2"`
	TxBlock   string

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Distribution  Distribution  `gorm:"foreignKey:DistributionId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	SellerProfile SellerProfile `gorm:"foreignKey:SellerProfileId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

// ----------------------------- GEO DATA -----------------------------

type Country struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`

	Regions []Region `gorm:"foreignKey:CountryId"`
}

type Region struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CountryId uint   `gorm:"not null"`

	Country   Country   `gorm:"foreignKey:CountryId"`
	Regencies []Regency `gorm:"foreignKey:RegionId"`
}

type Regency struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	RegionId uint   `gorm:"not null"`

	Region Region `gorm:"foreignKey:RegionId"`
}
