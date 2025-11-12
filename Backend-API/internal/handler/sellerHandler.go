package handler

import (
	"encoding/json"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/entity"
	"farm-integrated-web3/internal/usecase"
	"farm-integrated-web3/utils/helper"
	"farm-integrated-web3/utils/middleware"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type SellerHandler struct {
	retailerUC usecase.SellerUsecase
	validator  *validator.Validate
}

func NewSellerHandler(retailerUC usecase.SellerUsecase) *SellerHandler {
	return &SellerHandler{retailerUC, validator.New()}
}

// Add Seller Box godoc
// @Summary Add seller box.
// @Description This endpoint for seller add the cart of distribution form distribution.
// @Tags Seller
// @Accept json
// @Produce json
// @param distribution path integer true "distribution id"
// @param request body dto.CreateSellerBox true "request body"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /seller/distribution/{distribution} [post]
// @Security BearerAuth
func (h *SellerHandler) AddSellerBox(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Seller) {
		helper.HttpError(w, http.StatusForbidden, "you are not allowed")
		return
	}

	vars := mux.Vars(r)
	distributionId, _ := strconv.Atoi(vars["distribution"])

	var input dto.CreateSellerBox
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	input.SellerProfileId = claims.ProfileId
	input.DistributionId = uint(distributionId)
	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.retailerUC.AddSellerBox(r.Context(), &input); err != nil {
		switch err {
		case helper.ErrQuantityNotEnough:
			helper.HttpError(w, http.StatusBadRequest, "quantity is to much")
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}

		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)

}

// Update Seller  Cart godoc
// @Summary Update seller  cart.
// @Description This endpoint for seller  update the cart of distribution form distribution.
// @Tags Seller
// @Accept json
// @Produce json
// @param seller  path integer true "seller id"
// @param request body dto.UpdateSellerBox true "request body"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /seller/{seller} [patch]
// @Security BearerAuth
func (h *SellerHandler) UpdateSellerBox(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Seller) {
		helper.HttpError(w, http.StatusForbidden, "you are not allowed")
		return
	}

	vars := mux.Vars(r)
	SellerBoxId, _ := strconv.Atoi(vars["seller"])

	var input dto.UpdateSellerBox
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	input.SellerProfileId = claims.ProfileId
	input.SellerBoxId = uint(SellerBoxId)
	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.retailerUC.UpdateSellerBox(r.Context(), &input); err != nil {
		switch err {
		case helper.ErrQuantityNotEnough:
			helper.HttpError(w, http.StatusBadRequest, "quantity is to much")
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		default:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)
}

// Delete Seller  Cart godoc
// @Summary Delete seller  cart.
// @Description This endpoint for seller  delete the cart of distribution form distribution.
// @Tags Seller
// @Accept json
// @Produce json
// @param seller  path integer true "seller id"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /seller/{seller} [delete]
// @Security BearerAuth
func (h *SellerHandler) DeleteSellerBox(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Seller) {
		helper.HttpError(w, http.StatusForbidden, "you are not allowed")
		return
	}

	vars := mux.Vars(r)
	SellerBoxId, _ := strconv.Atoi(vars["seller"])
	if err := h.retailerUC.DeleteSellerBox(r.Context(), uint(SellerBoxId), claims.ProfileId); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		default:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)
}

// Search Seller  Cart godoc
// @Summary Search seller  cart.
// @Description This endpoint for search the  cart of distribution form seller cart.
// @Tags Seller
// @Accept json
// @Produce json
// @param search query string true "search"
// @Success 200 {object} []dto.GetSellerBox
// @Success 204
// @Failure 401 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /seller/search [get]
// @Security BearerAuth
func (h *SellerHandler) SearchSellerBox(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("search")
	if input == "" {
		helper.HttpError(w, http.StatusBadRequest, "search is required")
		return
	}

	result, err := h.retailerUC.SearchSellerBox(r.Context(), input)
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

// Get Seller  Cart By seller  Id godoc
// @Summary Get seller  cart by seller  id.
// @Description This endpoint for get the seller  cart from seller  id.
// @Tags Seller
// @Accept json
// @Produce json
// @Success 200 {object} []dto.GetSellerBox
// @Success 204
// @Failure 401 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /seller [get]
// @Security BearerAuth
func (h *SellerHandler) ListGetSellerBoxsbySellerId(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Seller) {
		helper.HttpError(w, http.StatusForbidden, "you are not allowed")
		return
	}

	result, err := h.retailerUC.ListGetSellerBoxsbySellerId(r.Context(), claims.ProfileId)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpWriter(w, http.StatusNotFound, err)
		case gorm.ErrEmptySlice:
			helper.HttpWriter(w, http.StatusNoContent, result)
		default:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, result)
}

// Get seller  Cart By Seller  Id godoc
// @Summary Get seller  cart by seller  id.
// @Description This endpoint for get the seller  cart from seller  id.
// @Tags Seller
// @Accept json
// @Produce json
// @Param seller path integer true "seller id"
// @Success 200 {object} []dto.GetSellerBox
// @Success 204
// @Failure 401 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /seller/id/{seller} [get]
// @Security BearerAuth
func (h *SellerHandler) GetSellerBoxById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	SellerBoxId, _ := strconv.Atoi(vars["seller"])

	fmt.Println(SellerBoxId)
	result, err := h.retailerUC.GetSellerBoxById(r.Context(), uint(SellerBoxId))
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

// Get FYP seller  Cart  godoc
// @Summary Get FYP seller  cart .
// @Description This endpoint for get FYP the seller  cart.
// @Tags Seller
// @Accept json
// @Produce json
// @Success 200 {object} []dto.GetSellerBox
// @Success 204
// @Failure 401 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /seller/fyp [get]
// @Security BearerAuth
func (h *SellerHandler) ListGetSellerBoxsbySellerIdFYP(w http.ResponseWriter, r *http.Request) {
	result, err := h.retailerUC.ListGetSellerBoxsbySellerIdFYP(r.Context())
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
