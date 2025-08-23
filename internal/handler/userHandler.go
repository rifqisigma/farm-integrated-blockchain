package handler

import (
	"encoding/json"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/internal/usecase"
	"farm-integrated-web3/utils/helper"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userUC    usecase.UserUsecase
	validator *validator.Validate
}

func NewUserHandler(userUC usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUC, validator.New()}
}

func (h *UserHandler) CreateConsumer(w http.ResponseWriter, r *http.Response) {
	var input dto.CreateConsumerProfile
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userUC.CreateConsumerProfile(&input); err != nil {

		return
	}
}
