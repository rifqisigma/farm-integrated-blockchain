package dto

type ResponseMessage struct {
	Message string `json:"message" example:"this is message"`
}

type ResponseError struct {
	Error string `json:"error" example:"this is error"`
}

type ResponseLogin struct {
	AccessToken  string `json:"access_token" example:"this is access token"`
	RefreshToken string `json:"refresh_token" example:"this is refresh token"`
}

type ResponseRefreshToken struct {
	RefreshToken string `json:"refresh_token" example:"this is refresh token"`
}

type ResponseAccessToken struct {
	AccessToken string `json:"access_token" example:"this is access token"`
}

type CountryRequest struct {
	Name          string        `json:"name" example:"Indonesia" `
	RegionRequest RegionRequest `json:"region" `
}

type RegionRequest struct {
	Name           string         `json:"name" example:"Jawa Barat"  `
	RegencyRequest RegencyRequest `json:"regency" `
}

type RegencyRequest struct {
	Name string `json:"name" example:"Bogor"`
}

type Country struct {
	Name   string `json:"name" example:"Indonesia" gorm:"column:country_name"`
	Region Region `json:"region"`
}

type Region struct {
	Name    string `json:"name" example:"Jawa Barat" gorm:"column:region_name"`
	Regency `json:"regency"`
}

type Regency struct {
	Name string `json:"name" example:"Bogor"  gorm:"column:regency_name"`
}
