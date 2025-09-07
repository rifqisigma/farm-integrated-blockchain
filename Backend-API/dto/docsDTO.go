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
	AccessToken string `json:"token_token" example:"this is access token"`
}
