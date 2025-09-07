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
// @param crop path integer true "crop id"
// @param request body dto.HarvestRequest true "create harvest"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /farm/harvest/crop/{crop} [post]
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

	vars := mux.Vars(r)
	cropId, _ := strconv.Atoi(vars["crop"])
	if cropId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "crop is required")
		return
	}

	var input dto.HarvestRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	input.FarmerProfileId = uint(claims.ProfileId)
	input.CropID = uint(cropId)

	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := h.farmerUC.CreateHarvest(&input)
	if err != nil {
		switch err {
		case helper.ErrInvalidTime:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
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
// @param crop path integer true "crop id"
// @param request body dto.HarvestRequest true "update harvest"
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

	cropId, _ := strconv.Atoi(vars["crop"])
	if cropId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "crop is required")
		return
	}

	var input dto.HarvestUpdate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	input.FarmerProfileId = claims.ProfileId
	input.HarvestId = uint(claims.UserID)

	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.farmerUC.UpdateHarvest(&input); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		case helper.ErrInvalidTime:
			helper.HttpError(w, http.StatusBadRequest, err.Error())

		case gorm.ErrInvalidData:
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

	if err := h.farmerUC.DeleteHarvest(claims.ProfileId, uint(harvestId)); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		case helper.ErrInvalidTime:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
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
// @Param request body dto.AcceptFarmerForDistributor true "request body accept"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /farm/distribution/{distribution} [patch]
// @Security BearerAuth
func (h *Farmerhandler) AcceptedFarmerForDistributor(w http.ResponseWriter, r *http.Request) {
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

	var input dto.AcceptFarmerForDistributor
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	input.FarmerProfieId = claims.ProfileId
	input.DistributionId = uint(distributionId)

	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.farmerUC.AcceptedFarmerForDistributor(&input); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
}

// Get Harvests By Farmer Id godoc
// @Summary Get Harvest by Farmer Id.
// @Description This endpoint for farmer get their own harvests.
// @Tags Farmer
// @Accept json
// @Produce json
// @Success 200 {object} []dto.GetListHarvest
// @Success 204 {object} []dto.GetListHarvest
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /farm/harvest [get]
// @Security BearerAuth
func (h *Farmerhandler) ListHarvestByFarmerId(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	result, err := h.farmerUC.ListHarvestByFarmerId(claims.ProfileId)
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

	result, err := h.farmerUC.HarvestById(uint(harvestId))
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
// @Success 200 {object} dto.GetHarvestById
// @Success 204 {object} []dto.GetListHarvest
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

	result, err := h.farmerUC.SearchHarvest(search)
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
