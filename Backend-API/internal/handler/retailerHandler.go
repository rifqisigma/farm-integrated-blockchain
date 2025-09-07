package handler

import (
	"encoding/json"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/entity"
	"farm-integrated-web3/internal/usecase"
	"farm-integrated-web3/utils/helper"
	"farm-integrated-web3/utils/middleware"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type RetailerHandler struct {
	retailerUC usecase.RetailerUsecase
	validator  *validator.Validate
}

func NewRetailerHandler(retailerUC usecase.RetailerUsecase) *RetailerHandler {
	return &RetailerHandler{retailerUC, validator.New()}
}

// Add Retailer Cart godoc
// @Summary Add retailer cart.
// @Description This endpoint for retailer add the cart of distribution form distribution.
// @Tags Retailer
// @Accept json
// @Produce json
// @param distribution path integer true "distribution id"
// @param request body dto.CreateRetailerCartRequest true "request body"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /retail/distribution/{distribution} [post]
// @Security BearerAuth
func (h *RetailerHandler) AddRetailerCart(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Retailer) {
		helper.HttpError(w, http.StatusForbidden, "you are not allowed")
		return
	}

	vars := mux.Vars(r)
	distributionId, _ := strconv.Atoi(vars["distribution"])

	var input dto.CreateRetailerCartRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	input.RetailerProfileId = claims.ProfileId
	input.DistributionId = uint(distributionId)
	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.retailerUC.AddRetailerCart(&input); err != nil {
		switch err {
		case gorm.ErrInvalidData:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}

		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)

}

// Update Retailer Cart godoc
// @Summary Update retailer cart.
// @Description This endpoint for retailer update the cart of distribution form distribution.
// @Tags Retailer
// @Accept json
// @Produce json
// @param distribution path integer true "distribution id"
// @param retailer path integer true "retailer id"
// @param request body dto.UpdateRetailerCartRequest true "request body"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /retail/{retailer}/distribution/{distribution} [patch]
// @Security BearerAuth
func (h *RetailerHandler) UpdateRetailerCart(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Retailer) {
		helper.HttpError(w, http.StatusForbidden, "you are not allowed")
		return
	}

	vars := mux.Vars(r)
	distributionId, _ := strconv.Atoi(vars["distribution"])
	retailerCartId, _ := strconv.Atoi(vars["retailer"])

	var input dto.UpdateRetailerCartRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	input.RetailerProfileId = claims.ProfileId
	input.DistributionId = uint(distributionId)
	input.RetailerCartId = uint(retailerCartId)
	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.retailerUC.UpdateRetailerCart(&input); err != nil {
		switch err {
		case helper.ErrInvalidTime:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		case gorm.ErrInvalidData:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		default:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)
}

// Delete Retailer Cart godoc
// @Summary Delete retailer cart.
// @Description This endpoint for retailer delete the cart of distribution form distribution.
// @Tags Retailer
// @Accept json
// @Produce json
// @param retailer path integer true "retailer id"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /retail/{retailer} [delete]
// @Security BearerAuth
func (h *RetailerHandler) DeleteRetailerCart(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Retailer) {
		helper.HttpError(w, http.StatusForbidden, "you are not allowed")
		return
	}

	vars := mux.Vars(r)
	retailerCartId, _ := strconv.Atoi(vars["retailer"])
	if err := h.retailerUC.DeleteRetailerCart(uint(retailerCartId), claims.ProfileId); err != nil {
		switch err {
		case helper.ErrInvalidTime:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		default:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)
}

// Search Retailer Cart godoc
// @Summary Search retailer cart.
// @Description This endpoint for search the  cart of distribution form retailer cart.
// @Tags Retailer
// @Accept json
// @Produce json
// @param search query string true "search"
// @Success 200 {object} []dto.GetRetailerCart
// @Success 204 {object} []dto.GetRetailerCart
// @Failure 401 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /retail/search [get]
// @Security BearerAuth
func (h *RetailerHandler) SearchRetailerCart(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("search")
	if input == "" {
		helper.HttpError(w, http.StatusBadRequest, "search is required")
		return
	}

	result, err := h.retailerUC.SearchRetailerCart(input)
	if err != nil {
		switch err {
		case gorm.ErrEmptySlice:
			helper.HttpWriter(w, http.StatusNoContent, result)
		default:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, result)
}

// Get Retailer Cart By Retailer Id godoc
// @Summary Get retailer cart by retailer id.
// @Description This endpoint for get the retailer cart from retailer id.
// @Tags Retailer
// @Accept json
// @Produce json
// @Success 200 {object} []dto.GetRetailerCart
// @Success 204 {object} []dto.GetRetailerCart
// @Failure 401 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /retail [get]
// @Security BearerAuth
func (h *RetailerHandler) GetRetailerCarts(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Retailer) {
		helper.HttpError(w, http.StatusForbidden, "you are not allowed")
		return
	}

	result, err := h.retailerUC.GetRetailerCarts(claims.ProfileId)
	if err != nil {
		switch err {
		case gorm.ErrEmptySlice:
			helper.HttpWriter(w, http.StatusNoContent, result)
		default:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, result)
}

// Get Retailer Cart By Retailer Id godoc
// @Summary Get retailer cart by retailer id.
// @Description This endpoint for get the retailer cart from retailer id.
// @Tags Retailer
// @Accept json
// @Produce json
// @Param retailer query integer true "retailer id"
// @Success 200 {object} []dto.GetRetailerCart
// @Success 204 {object} []dto.GetRetailerCart
// @Failure 401 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /retail/{retailer} [get]
// @Security BearerAuth
func (h *RetailerHandler) GetRetailerCartById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	retailerCartId, _ := strconv.Atoi(vars["retailer"])
	result, err := h.retailerUC.GetRetailerCartById(uint(retailerCartId))
	if err != nil {
		switch err {
		case gorm.ErrEmptySlice:
			helper.HttpWriter(w, http.StatusNoContent, result)
		default:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, result)
}
