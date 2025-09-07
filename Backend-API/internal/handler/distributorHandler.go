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

type DistributorHandler struct {
	DistributorUC usecase.DistributorUsecase
	Validator     *validator.Validate
}

func NewDistributorHandler(DistributorUC usecase.DistributorUsecase) *DistributorHandler {
	return &DistributorHandler{DistributorUC, validator.New()}
}

// Create Distribution godoc
// @Summary Create Distributions.
// @Description This endpoint for distributor create distribution.
// @Tags Distributor
// @Accept json
// @Produce json
// @Param harvest path integer true "harvest id"
// @Param request body dto.CreateDistributionRequest true "request body create"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /distribution/harvest/{harvest} [patch]
// @Security BearerAuth
func (h *DistributorHandler) CreateDistribution(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Distributor) {
		helper.HttpError(w, http.StatusForbidden, "you are not allowed")
		return
	}

	vars := mux.Vars(r)
	harvestId, _ := strconv.Atoi(vars["harvest"])

	var input dto.CreateDistributionRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	input.HarvestId = uint(harvestId)
	input.FarmerProfileId = claims.ProfileId

	if err := h.Validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.DistributorUC.CreateDistribution(&input); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		case gorm.ErrInvalidData:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)
}

// Update Distribution godoc
// @Summary Update Distributions.
// @Description This endpoint for distributor  update distribution.
// @Tags Distributor
// @Accept json
// @Produce json
// @Param distribution path integer true "distribution id"
// @Param request body dto.UpdateDistributionRequest true "request body update"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /distribution/{distribution} [patch]
// @Security BearerAuth
func (h *DistributorHandler) UpdateDistribution(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Distributor) {
		helper.HttpError(w, http.StatusForbidden, "you are not allowed")
		return
	}

	vars := mux.Vars(r)
	distributionId, _ := strconv.Atoi(vars["distribution"])

	var input dto.UpdateDistributionRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	input.DistributionId = uint(distributionId)
	input.DistributorProfileId = claims.ProfileId

	if err := h.Validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.DistributorUC.UpdateDistribution(&input); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		case gorm.ErrInvalidData:
			helper.HttpError(w, http.StatusBadRequest, err.Error())

		case helper.ErrInvalidTime:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)
}

// Delete Distribution godoc
// @Summary Delete Distribution.
// @Description This endpoint for distributor delete distribution.
// @Tags Distributor
// @Accept json
// @Produce json
// @Param distribution path integer true "distribution id"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /distribution/{distribution} [delete]
// @Security BearerAuth
func (h *DistributorHandler) DeleteDistribution(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Distributor) {
		helper.HttpError(w, http.StatusForbidden, "you are not allowed")
		return
	}

	vars := mux.Vars(r)
	distributionId, _ := strconv.Atoi(vars["distribution"])

	if err := h.DistributorUC.DeleteDistribution(uint(distributionId), claims.ProfileId); err != nil {
		switch err {
		case helper.ErrInvalidTime:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusBadRequest, err.Error())

		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)
}

// Search Distributions godoc
// @Summary Search Distributions.
// @Description This endpoint for search distributions.
// @Tags Distributor
// @Accept json
// @Produce json
// @Param search query string true "query search"
// @Success 200 {object} []dto.GetDistribution
// @Success 204 {object} []dto.GetDistribution
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /distribution/search [get]
// @Security BearerAuth
func (h *DistributorHandler) SearchDistributions(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	if search == "" {
		helper.HttpError(w, http.StatusBadRequest, "search is empty")
	}

	result, err := h.DistributorUC.SearchDistributions(search)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		case gorm.ErrEmptySlice:
			helper.HttpWriter(w, http.StatusNoContent, result)
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, result)
}

// Get Distributions By Distributor Id godoc
// @Summary Get distributions by distributor id.
// @Description This endpoint for get distributions by distributor id.
// @Tags Distributor
// @Accept json
// @Produce json
// @Success 200 {object} []dto.GetDistribution
// @Success 204 {object} []dto.GetDistribution
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /distribution [get]
// @Security BearerAuth
func (h *DistributorHandler) GetDistributionsByDistributorId(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Distributor) {
		helper.HttpError(w, http.StatusForbidden, "you are not allowed")
		return
	}
	result, err := h.DistributorUC.GetDistributionsByDistributorId(claims.ProfileId)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		case gorm.ErrEmptySlice:
			helper.HttpWriter(w, http.StatusNoContent, result)
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, result)
}

// Get Distribution By Id godoc
// @Summary Get distribution by id.
// @Description This endpoint for get detail information of distribution.
// @Tags Distributor
// @Accept json
// @Produce json
// @Success 200 {object} dto.GetHarvestById
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /distribution/{distribution} [get]
// @Security BearerAuth
func (h *DistributorHandler) GetDistributionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	distributionId, _ := strconv.Atoi(vars["distribution"])
	result, err := h.DistributorUC.GetDistributionByid(uint(distributionId))
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())

		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, result)
}

// Update Status Distribution godoc
// @Summary Update Status distribution.
// @Description This endpoint for update Status Distribution.
// @Tags Distributor
// @Accept json
// @Produce json
// @Param distribution path integer true "distribution cart id"
// @Param request body dto.UpdateStatusDistributionRequest true "body request"
// @Success 200
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /distribution/{distribution}/status [patch]
// @Security BearerAuth
func (h *DistributorHandler) UpdateStatusDistribution(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Distributor) {
		helper.HttpError(w, http.StatusForbidden, "you are not allowed")
		return
	}
	vars := mux.Vars(r)
	distributionId, _ := strconv.Atoi(vars["distribution"])

	var input dto.UpdateStatusDistributionRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	input.DistributionId = uint(distributionId)
	input.DistributorProfileId = claims.ProfileId

	if err := h.Validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
	}
	if err := h.DistributorUC.UpdateStatusOfDistribution(&input); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)
}

// Approved Retailer Cart  godoc
// @Summary Approved retailer cart.
// @Description This endpoint for distributor approve retailer cart.
// @Tags Distributor
// @Accept json
// @Produce json
// @Param retailerCart path integer true "retailer cart id"
// @Param request body dto.ApprovedRetailerCart true "body request"
// @Success 200
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /distribution/retailer-cart/{retailerCart} [patch]
// @Security BearerAuth
func (h *DistributorHandler) ApprovedRetailerCartForRetailer(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Distributor) {
		helper.HttpError(w, http.StatusForbidden, "you are not allowed")
		return
	}
	vars := mux.Vars(r)
	retailerCartId, _ := strconv.Atoi(vars["retailerCart"])

	var input dto.ApprovedRetailerCart
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	input.RetailerCartId = uint(retailerCartId)
	input.DistributorProfileId = claims.ProfileId

	if err := h.Validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
	}
	if err := h.DistributorUC.ApprovedRetailerCartForRetailer(&input); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)
}
