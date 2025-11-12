package dto

import "math/big"

type GetHarvestBc struct {
	Id              *big.Int `json:"id"`
	FarmerProfileId *big.Int `json:"farmer_profile_id"`
	CropId          *big.Int `json:"crop_id"`
	RegencyId       *big.Int `json:"regency_id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Quantity        *big.Int `json:"quantity"`
	BasePrice       *big.Int `json:"base_price"`
	CreatedAt       *big.Int `json:"created_at"`
}

type GetCollectorHarvestBc struct {
	Id                 *big.Int `json:"id"`
	CollectorProfileId *big.Int `json:"collector_profile_id"`
	HarvestId          *big.Int `json:"harvest_id"`
	Name               string   `json:"name"`
	Desc               string   `json:"desc"`
	Quantity           *big.Int `json:"quantity"`
	Price              *big.Int `json:"price"`
	BasePrice          *big.Int `json:"base_price"`
	CreatedAt          *big.Int `json:"created_at"`
}

type GetProcessorHarvestBc struct {
	Id                 *big.Int `json:"id"`
	ProcessorProfileId *big.Int `json:"processor_profile_id"`
	HarvestCollectorId *big.Int `json:"harvest_collector_id"`
	HarvestId          *big.Int `json:"harvest_id"`
	Name               string   `json:"name"`
	Desc               string   `json:"desc"`
	Quantity           *big.Int `json:"quantity"`
	BasePrice          *big.Int `json:"base_price"`
	Price              *big.Int `json:"price"`
	CreatedAt          *big.Int `json:"created_at"`
}

type GetDistributionBc struct {
	Id                   *big.Int `json:"id"`
	DistributorProfileId *big.Int `json:"distributor_profile_id"`
	DestinationId        *big.Int `json:"destination_id"`
	HarvestId            *big.Int `json:"harvest_id"`
	HarvestCollectorId   *big.Int `json:"harvest_collector_id"`
	HarvestProcessorId   *big.Int `json:"harvest_processor_id"`
	Name                 string   `json:"name"`
	Desc                 string   `json:"desc"`
	Quantity             *big.Int `json:"quantity"`
	BasePrice            *big.Int `json:"base_price"`
	Price                *big.Int `json:"price"`
	Transportation       string   `json:"transportation"`
	CreatedAt            *big.Int `json:"created_at"`
}

type GetSellerBoxBc struct {
	Id              *big.Int `json:"id"`
	SellerProfileId *big.Int `json:"seller_profile_id"`
	DistributionId  *big.Int `json:"distribution_id"`
	Name            string   `json:"name"`
	Desc            string   `json:"desc"`
	Quantity        *big.Int `json:"quantity"`
	BasePrice       *big.Int `json:"base_price"`
	Price           *big.Int `json:"price"`
	CreatedAt       *big.Int `json:"created_at"`
}
