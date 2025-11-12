package handler

import (
	"farm-integrated-web3/internal/usecase"
	"farm-integrated-web3/utils/helper"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BlockchainHandler struct {
	blockchainUC usecase.BlockchainUsecase
}

func NewBlockchainHandler(blockchainUC usecase.BlockchainUsecase) *BlockchainHandler {
	return &BlockchainHandler{blockchainUC}
}

// Get All Harvest godoc
// @Summary Get all harvest.
// @Description This endpoint for get all harvest from blockchain API.
// @Tags Blockchain
// @Accept json
// @Produce json
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /blockchain/harvest [get]
// @Security BearerAuth
func (h *BlockchainHandler) GetAllHarvestBc(w http.ResponseWriter, r *http.Request) {
	result, err := h.blockchainUC.GetAllHarvest()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.HttpWriter(w, http.StatusOK, result)
}

// Get All Harvest Collector godoc
// @Summary Get all harvest collector.
// @Description This endpoint for get all harvest collector from blockchain API.
// @Tags Blockchain
// @Accept json
// @Produce json
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /blockchain/harvest-collector [get]
// @Security BearerAuth
func (h *BlockchainHandler) GetAllHarvestCollectorBc(w http.ResponseWriter, r *http.Request) {
	result, err := h.blockchainUC.GetAllHarvestCollector()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.HttpWriter(w, http.StatusOK, result)
}

// Get All Harvest Processor godoc
// @Summary Get all harvest processor.
// @Description This endpoint for get all harvest processor from blockchain API.
// @Tags Blockchain
// @Accept json
// @Produce json
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /blockchain/harvest-processor [get]
// @Security BearerAuth
func (h *BlockchainHandler) GetAllHarvestProcessorBc(w http.ResponseWriter, r *http.Request) {
	result, err := h.blockchainUC.GetAllHarvestProcessor()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.HttpWriter(w, http.StatusOK, result)
}

// Get All Distribution godoc
// @Summary Get all distribution.
// @Description This endpoint for get all distribution from blockchain API.
// @Tags Blockchain
// @Accept json
// @Produce json
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /blockchain/distribution [get]
// @Security BearerAuth
func (h *BlockchainHandler) GetAllDsitributionBc(w http.ResponseWriter, r *http.Request) {
	result, err := h.blockchainUC.GetAllDistribution()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.HttpWriter(w, http.StatusOK, result)
}

// Get All Seller-Box godoc
// @Summary Get all seller-box .
// @Description This endpoint for get all seller-box from blockchain API.
// @Tags Blockchain
// @Accept json
// @Produce json
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /blockchain/seller-box [get]
// @Security BearerAuth
func (h *BlockchainHandler) GetAllSellerBoxBc(w http.ResponseWriter, r *http.Request) {
	result, err := h.blockchainUC.GetAllSellerBox()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.HttpWriter(w, http.StatusOK, result)
}

// Get Harvest By Id godoc
// @Summary Get harvest by id.
// @Description This endpoint for get harvest from blockchain API by Id.
// @Tags Blockchain
// @Accept json
// @Produce json
// @Param harvest path integer true "harvest id"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /blockchain/harvest/{harvest} [get]
// @Security BearerAuth
func (h *BlockchainHandler) GetHarvestBcByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	harvestId, _ := strconv.Atoi(vars["harvest"])
	if harvestId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "harvest is required")
		return
	}

	result, err := h.blockchainUC.GetHarvestById(int64(harvestId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.HttpWriter(w, http.StatusOK, result)
}

// Get Harvest Collector By Id godoc
// @Summary Get harvest collector by id.
// @Description This endpoint for get harvest collector from blockchain API by Id.
// @Tags Blockchain
// @Accept json
// @Produce json
// @Param collector path integer true "collector id"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /blockchain/harvest-collector/{collector} [get]
// @Security BearerAuth
func (h *BlockchainHandler) GetHarvestCollectorBcByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	processorId, _ := strconv.Atoi(vars["collector"])
	if processorId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "collector is required")
		return
	}

	result, err := h.blockchainUC.GetHarvestCollectorById(int64(processorId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.HttpWriter(w, http.StatusOK, result)
}

// Get Harvest Processor By Id godoc
// @Summary Get harvest processor by id.
// @Description This endpoint for get harvest processor from blockchain API by Id.
// @Tags Blockchain
// @Accept json
// @Produce json
// @Param processor path integer true "processor id"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /blockchain/harvest-processor/{processor} [get]
// @Security BearerAuth
func (h *BlockchainHandler) GetHarvestProcessorBcByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	processorId, _ := strconv.Atoi(vars["processor"])
	if processorId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "processor is required")
		return
	}

	result, err := h.blockchainUC.GetHarvestProcessorById(int64(processorId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.HttpWriter(w, http.StatusOK, result)
}

// Get Distribution By Id godoc
// @Summary Get harvest processor by id.
// @Description This endpoint for get distribution from blockchain API by Id.
// @Tags Blockchain
// @Accept json
// @Produce json
// @Param distribution path integer true "distribution id"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /blockchain/distribution/{distribution} [get]
// @Security BearerAuth
func (h *BlockchainHandler) GetDistributionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	distributionId, _ := strconv.Atoi(vars["distribution"])
	if distributionId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "distrbution is required")
		return
	}

	result, err := h.blockchainUC.GetDistributionById(int64(distributionId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.HttpWriter(w, http.StatusOK, result)
}

// Get Seller-Box By Id godoc
// @Summary Get seller-box by id.
// @Description This endpoint for get seller-box from blockchain API by Id.
// @Tags Blockchain
// @Accept json
// @Produce json
// @Param seller-box path integer true "seller-box id"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /blockchain/seller-box/{seller-box} [get]
// @Security BearerAuth
func (h *BlockchainHandler) GetSellerBoxBcByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sellerBoxId, _ := strconv.Atoi(vars["seller-box"])
	if sellerBoxId == 0 {
		helper.HttpError(w, http.StatusBadRequest, "seller-box is required")
		return
	}

	result, err := h.blockchainUC.GetSellerBoxById(int64(sellerBoxId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.HttpWriter(w, http.StatusOK, result)
}
