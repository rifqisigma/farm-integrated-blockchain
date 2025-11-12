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

type Farmerhandler struct {
	farmerUC  usecase.FarmerUsecase
	validator *validator.Validate
}

func NewFarmerHandler(farmerUC usecase.FarmerUsecase) *Farmerhandler {
	return &Farmerhandler{farmerUC, validator.New()}
}

// Create Harvest godoc
// @Summary Create harvest for farmer.
// @Description This endpoint for farmer increase the harvest so that it can be distributed by distributors .
// @Tags Farmer
// @Accept json
// @Produce json
// @param request body dto.HarvestCreate true "create harvest"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /farm/harvest [post]
// @Security BearerAuth
func (h *Farmerhandler) CreateHarvest(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Farmer) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	var input dto.HarvestCreate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	input.FarmerProfileId = uint(claims.ProfileId)

	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := h.farmerUC.CreateHarvest(r.Context(), &input)
	if err != nil {
		switch err {
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}

		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)
}

// Update Harvest godoc
// @Summary Update harvest for farmer.
// @Description This endpoint for farmer update data of harvest .
// @Tags Farmer
// @Accept json
// @Produce json
// @param harvest path integer true "harvest id"
// @param request body dto.HarvestUpdate true "update harvest"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /farm/harvest/{harvest} [patch]
// @Security BearerAuth
func (h *Farmerhandler) UpdateHarvest(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Farmer) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	vars := mux.Vars(r)
	harvestId, _ := strconv.Atoi(vars["harvest"])
	if harvestId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "harvest is required")
		return
	}

	var input dto.HarvestUpdate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	input.FarmerProfileId = claims.ProfileId
	input.HarvestId = uint(harvestId)

	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.farmerUC.UpdateHarvest(r.Context(), &input); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())

		case helper.ErrQuantityNotEnough:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)

}

// Delete Harvest godoc
// @Summary Delete harvest for farmer.
// @Description This endpoint for farmer delete the harvest.
// @Tags Farmer
// @Accept json
// @Produce json
// @param harvest path integer true "harvest id"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /farm/harvest/{harvest} [delete]
// @Security BearerAuth
func (h *Farmerhandler) DeleteHarvest(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Farmer) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	vars := mux.Vars(r)
	harvestId, _ := strconv.Atoi(vars["harvest"])
	if harvestId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "harvest is required")
		return
	}

	if err := h.farmerUC.DeleteHarvest(r.Context(), claims.ProfileId, uint(harvestId)); err != nil {
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

// Accepted Distribution godoc
// @Summary Accepted distribution for distributor.
// @Description This endpoint for farmer accept the distribution, the distribution cant proceed  to next step if farmer not yet accept it.
// @Tags Farmer
// @Accept json
// @Produce json
// @Param distribution path integer true "distribution id"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /farm/distribution/{distribution} [patch]
// @Security BearerAuth
func (h *Farmerhandler) AcceptDistribution(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Farmer) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	vars := mux.Vars(r)
	distributionId, _ := strconv.Atoi(vars["distribution"])
	if distributionId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "harvest is required")
		return
	}

	if err := h.farmerUC.AcceptDistributor(r.Context(), claims.ProfileId, uint(distributionId)); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
}

// Accepted Collector godoc
// @Summary Accepted collector-harvest for collector.
// @Description This endpoint for farmer accept the collector, the collector cant proceed  to next step if farmer not yet accept it.
// @Tags Farmer
// @Accept json
// @Produce json
// @Param collector path integer true "collector id"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /farm/collector/{collector} [patch]
// @Security BearerAuth
func (h *Farmerhandler) AcceptHarvestCollector(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Farmer) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	vars := mux.Vars(r)
	collectorId, _ := strconv.Atoi(vars["collector"])
	if collectorId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "harvest is required")
		return
	}

	if err := h.farmerUC.AcceptHarvestCollector(r.Context(), claims.ProfileId, uint(collectorId)); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
}

// Accepted Processor godoc
// @Summary Accepted processor-harvest for processor.
// @Description This endpoint for farmer accept the processor, the processor cant proceed  to next step if farmer not yet accept it.
// @Tags Farmer
// @Accept json
// @Produce json
// @Param processor path integer true "processor id"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /farm/processor/{processor} [patch]
// @Security BearerAuth
func (h *Farmerhandler) AcceptHarvestProcessor(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Farmer) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	vars := mux.Vars(r)
	procesorId, _ := strconv.Atoi(vars["processor"])
	if procesorId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "harvest is required")
		return
	}

	if err := h.farmerUC.AcceptHarvestProcessor(r.Context(), claims.ProfileId, uint(procesorId)); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
}

// Update Status Harvest godoc
// @Summary Update status harvest.
// @Description This endpoint for farmer update status of harvest.
// @Tags Farmer
// @Accept json
// @Produce json
// @param harvest path integer true "harvest id"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /farm/status/{harvest} [patch]
// @Security BearerAuth
func (h *Farmerhandler) UpdateStatusHarvest(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Farmer) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}
	vars := mux.Vars(r)
	harvestId, _ := strconv.Atoi(vars["harvest"])
	if harvestId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "harvest is required")
		return
	}
	if err := h.farmerUC.UpdateStatusHarvest(r.Context(), claims.ProfileId, uint(harvestId)); err != nil {
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

// Get Harvests By Farmer Id godoc
// @Summary Get Harvest by Farmer Id.
// @Description This endpoint for farmer get their own harvests.
// @Tags Farmer
// @Accept json
// @Produce json
// @Success 200 {object} []dto.GetListHarvest
// @Success 204
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /farm [get]
// @Security BearerAuth
func (h *Farmerhandler) ListHarvestByFarmerId(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	result, err := h.farmerUC.ListHarvestByFarmerId(r.Context(), claims.ProfileId)
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

// Get Harvest By Id godoc
// @Summary Get harvest by id.
// @Description This endpoint for a get detail information of a harvest.
// @Tags Farmer
// @Accept json
// @Produce json
// @param harvest path integer true "harvest id"
// @Success 200 {object} dto.GetHarvestById
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /farm/harvest/{harvest} [get]
// @Security BearerAuth
func (h *Farmerhandler) HarvestById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	harvestId, _ := strconv.Atoi(vars["harvest"])
	if harvestId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "harvest is required")
		return
	}

	result, err := h.farmerUC.HarvestById(r.Context(), uint(harvestId))
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

// Search Harvest godoc
// @Summary Search harvest.
// @Description This endpoint for search harvest by crop name and a name farmer.
// @Tags Farmer
// @Accept json
// @Produce json
// @param search query string true "query search"
// @Success 200 {object} []dto.GetListHarvest
// @Success 204
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /farm/search [get]
// @Security BearerAuth
func (h *Farmerhandler) SearchHarvest(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	if search == "" {
		helper.HttpError(w, http.StatusBadRequest, "search is required")
		return
	}

	result, err := h.farmerUC.SearchHarvest(r.Context(), search)
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

// Get FYP Harvest godoc
// @Summary Get FYP harvest.
// @Description This endpoint for get FYP harvest.
// @Tags Farmer
// @Accept json
// @Produce json
// @Success 200 {object} []dto.GetListHarvest
// @Success 204
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /farm/fyp [get]
// @Security BearerAuth
func (h *Farmerhandler) ListHarvestFYP(w http.ResponseWriter, r *http.Request) {

	result, err := h.farmerUC.ListHarvestFYP(r.Context())
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
