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

type CollectorHandler struct {
	collectorUC usecase.CollectorUsecase
	validator   *validator.Validate
}

func NewCollectorHandler(collectorUC usecase.CollectorUsecase) *CollectorHandler {
	return &CollectorHandler{collectorUC, validator.New()}
}

// Create Harvest Collector godoc
// @Summary Create Harvest Collector.
// @Description This endpoint for create harvest collector by collector.
// @Tags Collector
// @Accept json
// @Produce json
// @param request body dto.CreateCollector true "request body create"
// @param harvest path integer  true "harvest id"
// @Success 200
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /collector/harvest/{harvest} [post]
// @Security BearerAuth
func (h *CollectorHandler) CreateHarvestCollector(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Collector) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	vars := mux.Vars(r)
	harvestId, _ := strconv.Atoi(vars["harvest"])
	if harvestId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "harvest is required")
		return
	}

	var input dto.CreateCollector
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	input.CollectorProfileId = claims.ProfileId
	input.HarvestId = uint(harvestId)

	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.collectorUC.CreateHarvestCollector(r.Context(), &input); err != nil {
		switch err {
		case helper.ErrQuantityNotEnough:
			helper.HttpError(w, http.StatusBadRequest, "quantity is too much")
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)

}

// Update Harvest Collector godoc
// @Summary Update Harvest Collector.
// @Description This endpoint for Update harvest collector by collector.
// @Tags Collector
// @Accept json
// @Produce json
// @param request body dto.UpdateCollector true "request body update"
// @param collector path integer  true "collector id"
// @Success 200
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /collector/{collector} [patch]
// @Security BearerAuth
func (h *CollectorHandler) UpdateHarvestCollector(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Collector) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	vars := mux.Vars(r)
	collectorId, _ := strconv.Atoi(vars["collector"])
	if collectorId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "harvest is required")
		return
	}

	var input dto.UpdateCollector
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	input.CollectorProfileId = claims.ProfileId
	input.CollectorId = uint(collectorId)

	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.collectorUC.UpdateHarvestCollector(r.Context(), &input); err != nil {
		switch err {
		case helper.ErrQuantityNotEnough:
			helper.HttpError(w, http.StatusBadRequest, "quantity is too much")
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)
}

// Delete Harvest Collector godoc
// @Summary Delete Harvest Collector.
// @Description This endpoint for Delete harvest collector by collector.
// @Tags Collector
// @Accept json
// @Produce json
// @param collector path integer  true "collector id"
// @Success 200
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /collector/{collector} [delete]
// @Security BearerAuth
func (h *CollectorHandler) DeleteHarvestCollector(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Collector) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	vars := mux.Vars(r)
	collectorId, _ := strconv.Atoi(vars["collector"])
	if err := h.collectorUC.DeleteHarvestCollector(r.Context(), uint(collectorId), claims.ProfileId); err != nil {
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

// Accept Harvest Processor for Processor By Collector godoc
// @Summary  Accept harvest processor for processor by collector.
// @Description This endpoint for Accept harvest processor by collector.
// @Tags Collector
// @Accept json
// @Produce json
// @param processor path integer  true "processor id"
// @Success 200
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /collector/accept/processor/{processor} [patch]
// @Security BearerAuth
func (h *CollectorHandler) AcceptHarvestProcessor(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Collector) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	vars := mux.Vars(r)
	processorId, _ := strconv.Atoi(vars["processor"])
	if err := h.collectorUC.AcceptHarvestProcessor(r.Context(), claims.ProfileId, uint(processorId)); err != nil {
		switch err {
		case helper.ErrQuantityNotEnough:
			helper.HttpError(w, http.StatusBadRequest, "quantity is too much")
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)
}

// Accept Distribution for Distributor By Collector godoc
// @Summary  Accept distribution for distributor by collector.
// @Description This endpoint for Accept distribution by collector.
// @Tags Collector
// @Accept json
// @Produce json
// @param distribution path integer  true "distribution id"
// @Success 200
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /collector/accept/distribution/{distribution} [patch]
// @Security BearerAuth
func (h *CollectorHandler) AcceptDistributor(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Collector) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	vars := mux.Vars(r)
	distributorId, _ := strconv.Atoi(vars["distribution"])
	if err := h.collectorUC.AcceptDistributor(r.Context(), claims.ProfileId, uint(distributorId)); err != nil {
		switch err {
		case helper.ErrQuantityNotEnough:
			helper.HttpError(w, http.StatusBadRequest, "quantity is too much")
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)
}

// Get Harvest Collector By Collector Profile Id godoc
// @Summary  Get harvest collector by collector profile Id.
// @Description This endpoint for get harvest collector by collector profile id.
// @Tags Collector
// @Accept json
// @Produce json
// @Success 200 {object} []dto.GetListHarvestCollector
// @Success 204
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /collector [get]
// @Security BearerAuth
func (h *CollectorHandler) ListHarvestCollectorByCollectorId(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Collector) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	response, err := h.collectorUC.ListHarvestCollectorByCollectorId(r.Context(), claims.ProfileId)
	if err != nil {
		switch err {
		case gorm.ErrEmptySlice:
			helper.HttpWriter(w, http.StatusNoContent, response)
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, response)
}

// Get Harvest Collector FYP godoc
// @Summary  Get harvest collector.
// @Description This endpoint for get harvest collector FYP.
// @Tags Collector
// @Accept json
// @Produce json
// @Success 200 {object} []dto.GetListHarvestCollector
// @Success 204
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /collector/fyp [get]
// @Security BearerAuth
func (h *CollectorHandler) ListHarvestCollectorFYP(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Collector) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	response, err := h.collectorUC.ListHarvestCollectorFYP(r.Context())
	if err != nil {
		switch err {
		case gorm.ErrEmptySlice:
			helper.HttpWriter(w, http.StatusNoContent, response)
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, response)
}

// Search Harvest Collector godoc
// @Summary  Search harvest collector.
// @Description This endpoint for search harvest collector.
// @Tags Collector
// @Accept json
// @Produce json
// @param search query string true "search query"
// @Success 200 {object} []dto.GetListHarvestCollector
// @Success 204
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /collector/search [get]
// @Security BearerAuth
func (h *CollectorHandler) SearchHarvestCollector(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Collector) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	search := r.URL.Query().Get("search")
	if search == "" {
		helper.HttpError(w, http.StatusBadRequest, "search is required")
		return
	}

	response, err := h.collectorUC.SearchHarvestCollector(r.Context(), search)
	if err != nil {
		switch err {
		case gorm.ErrEmptySlice:
			helper.HttpWriter(w, http.StatusNoContent, response)
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, response)
}

// Get Harvest Collector By Id godoc
// @Summary   Get harvest collector by id.
// @Description This endpoint for get harvest collector by id.
// @Tags Collector
// @Accept json
// @Produce json
// @param collector path integer true "collector id"
// @Success 200 {object} dto.GetHarvestCollectorById
// @Success 204 {object} dto.GetHarvestCollectorById
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /collector/id/{collector} [get]
// @Security BearerAuth
func (h *CollectorHandler) GetHarvestCollectorById(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Collector) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	vars := mux.Vars(r)
	collectorId, _ := strconv.Atoi(vars["collector"])

	response, err := h.collectorUC.GetHarvestCollectorById(r.Context(), uint(collectorId))
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, response)
}
