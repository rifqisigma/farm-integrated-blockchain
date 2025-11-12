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
// @Param request body dto.CreateDistribution true "request body create"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /distribution [post]
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

	var input dto.CreateDistribution
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	input.DistributorProfileId = claims.ProfileId
	if err := h.Validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.DistributorUC.CreateDistribution(r.Context(), &input); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		case helper.ErrQuantityNotEnough:
			helper.HttpError(w, http.StatusBadRequest, "quantity is to much")
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
// @Param request body dto.UpdateDistribution true "request body update"
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

	var input dto.UpdateDistribution
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

	if err := h.DistributorUC.UpdateDistribution(r.Context(), &input); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		case helper.ErrQuantityNotEnough:
			helper.HttpError(w, http.StatusBadRequest, "quantity is to much")
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

	if err := h.DistributorUC.DeleteDistribution(r.Context(), uint(distributionId), claims.ProfileId); err != nil {
		switch err {
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
// @Success 200 {object} []dto.GetListDistribution
// @Success 204
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

	result, err := h.DistributorUC.SearchDistributions(r.Context(), search)
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
// @Success 200 {object} []dto.GetListDistribution
// @Success 204
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /distribution [get]
// @Security BearerAuth
func (h *DistributorHandler) GetListDistributionsByDistributorId(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Distributor) {
		helper.HttpError(w, http.StatusForbidden, "you are not allowed")
		return
	}
	result, err := h.DistributorUC.GetListDistributionsByDistributorId(r.Context(), claims.ProfileId)
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
// @param distribution path integer true "distribution id"
// @Success 200 {object} dto.GetHarvestById
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /distribution/id/{distribution} [get]
// @Security BearerAuth
func (h *DistributorHandler) GetDistributionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	distributionId, _ := strconv.Atoi(vars["distribution"])
	result, err := h.DistributorUC.GetDistributionById(r.Context(), uint(distributionId))
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
// @Param request body dto.UpdateStatusDistribution true "body request"
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

	var input dto.UpdateStatusDistribution
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	input.DistributionId = uint(distributionId)
	input.DistributorProfileId = claims.ProfileId

	if err := h.Validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
	}
	if err := h.DistributorUC.UpdateStatusOfDistribution(r.Context(), &input); err != nil {
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

// Approved seller  Cart  godoc
// @Summary Approved seller  cart.
// @Description This endpoint for distributor approve seller cart.
// @Tags Distributor
// @Accept json
// @Produce json
// @Param sellerbox path integer true "sellerbox id"
// @Success 200
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /distribution/accept/seller/{sellerbox} [patch]
// @Security BearerAuth
func (h *DistributorHandler) AcceptSeller(w http.ResponseWriter, r *http.Request) {
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
	SellerBoxId, _ := strconv.Atoi(vars["sellerbox"])

	if err := h.DistributorUC.AcceptSeller(r.Context(), claims.ProfileId, uint(SellerBoxId)); err != nil {
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

// Get FYP Distributions godoc
// @Summary Get FYP Distributions.
// @Description This endpoint for get FYP distributions.
// @Tags Distributor
// @Accept json
// @Produce json
// @Success 200 {object} []dto.GetListDistribution
// @Success 204
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /distribution/fyp [get]
// @Security BearerAuth
func (h *DistributorHandler) GetListDistributionFYP(w http.ResponseWriter, r *http.Request) {
	result, err := h.DistributorUC.GetListDistributionFYP(r.Context())
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
