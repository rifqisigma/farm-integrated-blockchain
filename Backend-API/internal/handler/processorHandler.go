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

type ProcessorHandler struct {
	processorUc usecase.ProcessorUsecase
	validator   *validator.Validate
}

func NewProcessorHandler(processorUc usecase.ProcessorUsecase) *ProcessorHandler {
	return &ProcessorHandler{processorUc, validator.New()}
}

// Create Harvest Processor godoc
// @Summary  Create harvest processor.
// @Description This endpoint for create harvest processor by processor.
// @Tags Processor
// @Accept json
// @Produce json
// @param request body dto.CreateProcessor true "req body create"
// @Success 200
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /processor [post]
// @Security BearerAuth
func (h *ProcessorHandler) CreateProcessor(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Processor) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	var input dto.CreateProcessor
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	input.ProcessorProfileId = claims.ProfileId

	if input.HarvestCollectorId != 0 && input.HarvestId != 0 {
		helper.HttpError(w, http.StatusBadRequest, "id target only 1")
		return
	}
	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.processorUc.CreateProcessor(r.Context(), &input); err != nil {
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

// Update Harvest Processor godoc
// @Summary  Update harvest processor.
// @Description This endpoint for update harvest processor by processor.
// @Tags Processor
// @Accept json
// @Produce json
// @param request body dto.UpdateProcessor true "req body update"
// @param processor path integer true "processor id"
// @Success 200
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /processor/{processor} [patch]
// @Security BearerAuth
func (h *ProcessorHandler) UpdateProcessor(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Processor) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	vars := mux.Vars(r)
	processorId, _ := strconv.Atoi(vars["processor"])
	if processorId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "processor is required")
		return
	}

	var input dto.UpdateProcessor
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	input.ProcessorProfileId = claims.ProfileId
	input.ProcessorHarvestId = uint(processorId)

	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.processorUc.UpdateProcessor(r.Context(), &input); err != nil {
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

// Delete Harvest Processor godoc
// @Summary  Delete harvest processor.
// @Description This endpoint for delete harvest processor by processor.
// @Tags Processor
// @Accept json
// @Produce json
// @param processor path integer true "processor id"
// @Success 200
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /processor/{processor} [delete]
// @Security BearerAuth
func (h *ProcessorHandler) DeleteProcessor(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Processor) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	vars := mux.Vars(r)
	processorId, _ := strconv.Atoi(vars["processor"])
	if processorId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "processor is required")
		return
	}

	if err := h.processorUc.DeleteProcessor(r.Context(), uint(processorId), claims.ProfileId); err != nil {
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

// Accept Distribution for Distributor By Processor godoc
// @Summary Accept distribution for distributor by processor.
// @Description This endpoint for Accept distribution by processor.
// @Tags Processor
// @Accept json
// @Produce json
// @param distribution path integer  true "distribution id"
// @Success 200
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /processor/accept/distribution/{distribution} [patch]
// @Security BearerAuth
func (h *ProcessorHandler) AcceptDistributor(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Processor) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	vars := mux.Vars(r)
	distributorId, _ := strconv.Atoi(vars["distribution"])
	if distributorId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "processor is required")
		return
	}

	if err := h.processorUc.AcceptDistributor(r.Context(), claims.ProfileId, uint(distributorId)); err != nil {
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

// Get Harvest Processor By Processor Profile Id godoc
// @Summary Get harvest processor by processor profile id.
// @Description This endpoint for get harvest processor by processor profile id.
// @Tags Processor
// @Accept json
// @Produce json
// @Success 200 {object} []dto.GetListHarvestProcessor
// @Success 204
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /processor [get]
// @Security BearerAuth
func (h *ProcessorHandler) ListHarvestProcessorByProcessorId(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Processor) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	response, err := h.processorUc.ListHarvestProcessorByProcessorId(r.Context(), claims.ProfileId)
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

// Search Harvest Processor godoc
// @Summary   Search harvest processor.
// @Description This endpoint for search harvest processor.
// @Tags Processor
// @Accept json
// @Produce json
// @param search query string true "search query"
// @Success 200 {object} []dto.GetListHarvestProcessor
// @Success 204
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /processor/search [get]
// @Security BearerAuth
func (h *ProcessorHandler) SearchHarvestProcessor(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Processor) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	search := r.URL.Query().Get("search")
	if search == "" {
		helper.HttpError(w, http.StatusBadRequest, "search is required")
		return
	}

	response, err := h.processorUc.SearchHarvestProcessor(r.Context(), search)
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

// Get Harvest Processor FYP godoc
// @Summary   Get harvest processor FYP.
// @Description This endpoint for get harvest processor FYP.
// @Tags Processor
// @Accept json
// @Produce json
// @Success 200 {object} []dto.GetListHarvestProcessor
// @Success 204
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /processor/fyp [get]
// @Security BearerAuth
func (h *ProcessorHandler) ListHarvestProcessorFYP(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Processor) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	response, err := h.processorUc.ListHarvestProcessorFYP(r.Context())
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

// Get Harvest Processor By Id godoc
// @Summary   Get harvest processor by id.
// @Description This endpoint for get harvest processor by id.
// @Tags Processor
// @Accept json
// @Produce json
// @param processor path integer true "processor id"
// @Success 200 {object} dto.GetHarvestProcessorById
// @Success 204
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /processor/id/{processor} [get]
// @Security BearerAuth
func (h *ProcessorHandler) GetHarvestProcessorById(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if claims.Role != string(entity.Processor) {
		helper.HttpError(w, http.StatusForbidden, "forbidden entry")
		return
	}

	vars := mux.Vars(r)
	processorId, _ := strconv.Atoi(vars["processor"])
	if processorId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "processor is required")
		return
	}

	response, err := h.processorUc.GetHarvestProcessorById(r.Context(), uint(processorId))
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
