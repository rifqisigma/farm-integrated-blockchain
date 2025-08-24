package handler

import (
	"encoding/json"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/internal/usecase"
	"farm-integrated-web3/utils/helper"
	"farm-integrated-web3/utils/middleware"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userUC    usecase.UserUsecase
	validator *validator.Validate
}

func NewUserHandler(userUC usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUC, validator.New()}
}

func (h *UserHandler) CreateConsumer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid user id")
		return
	}
	if id == 0 {
		helper.HttpError(w, http.StatusBadRequest, "email is required")
		return
	}

	var input dto.CreateConsumerProfile
	input.UserId = uint(id)
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userUC.CreateConsumerProfile(&input); err != nil {
		switch err {
		case helper.ErrUserNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
			return
		}

		helper.HttpWriter(w, http.StatusCreated, "profile created successfully")
	}
}

func (h *UserHandler) CreateFarmer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid user id")
		return
	}
	if id == 0 {
		helper.HttpError(w, http.StatusBadRequest, "email is required")
		return
	}

	var input dto.CreateFarmerProfile
	input.UserId = uint(id)
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userUC.CreateFarmerProfile(&input); err != nil {
		switch err {
		case helper.ErrUserNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusCreated, "profile created successfully")
}

func (h *UserHandler) CreateDistributor(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid user id")
		return
	}
	if id == 0 {
		helper.HttpError(w, http.StatusBadRequest, "email is required")
		return
	}
	var input dto.CreateDistributorProfile
	input.UserId = uint(id)
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userUC.CreateDistributorProfile(&input); err != nil {
		switch err {
		case helper.ErrUserNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusCreated, "profile created successfully")
}

func (h *UserHandler) CreateRetailer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid user id")
		return
	}
	if id == 0 {
		helper.HttpError(w, http.StatusBadRequest, "email is required")
		return
	}

	var input dto.CreateRetailerProfile
	input.UserId = uint(id)
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userUC.CreateRetailerProfile(&input); err != nil {
		switch err {
		case helper.ErrUserNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	helper.HttpWriter(w, http.StatusCreated, "profile created successfully")
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := h.userUC.Logout(claims.UserID); err != nil {
		switch err {
		case helper.ErrUserNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, "logout successful")
}

func (h *UserHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := h.userUC.DeleteAccount(claims.UserID); err != nil {
		switch err {
		case helper.ErrUserNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, "account deleted successfully")
}

func (h *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var input dto.UserChangePassword
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if input.NewPassword != input.ConfirmNewPassword {
		helper.HttpError(w, http.StatusBadRequest, "new password and confirm password do not match")
		return
	}

	if err := h.userUC.ChangePassword(&input); err != nil {
		switch err {
		case helper.ErrUserNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, "password changed successfully")
}

func (h *UserHandler) UpdateConsumerProfile(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateConsumerProfile
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	input.UserId = claims.UserID

	if err := h.userUC.UpdateConsumerProfile(&input); err != nil {
		switch err {
		case helper.ErrUserNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, "profile updated successfully")
}
func (h *UserHandler) UpdateFarmerProfile(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateFarmerProfile
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	input.UserId = claims.UserID

	if err := h.userUC.UpdateFarmerProfile(&input); err != nil {
		switch err {
		case helper.ErrUserNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, "profile updated successfully")
}

func (h *UserHandler) UpdateDistributorProfile(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateDistributorProfile
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	input.UserId = claims.UserID

	if err := h.userUC.UpdateDistributorProfile(&input); err != nil {
		switch err {
		case helper.ErrUserNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, "profile updated successfully")
}

func (h *UserHandler) UpdateRetailerProfile(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateRetailerProfile
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	input.UserId = claims.UserID

	if err := h.userUC.UpdateRetailerProfile(&input); err != nil {
		switch err {
		case helper.ErrUserNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, "profile updated successfully")
}
