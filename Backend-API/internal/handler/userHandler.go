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

type UserHandler struct {
	userUC    usecase.UserUsecase
	validator *validator.Validate
}

func NewUserHandler(userUC usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUC, validator.New()}
}

// Create Profile godoc
// @Summary Create profile.
// @Description This endpoint for new user create profile for first if yser skip it, then consumer.
// @Param id path integer true "User ID"
// @Tags User
// @Accept json
// @Produce json
// @param request body dto.CreateProfileRequest true "create profile"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 409 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /user/{id}/profile [post]
func (h *UserHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid user id")
		return
	}
	if id == 0 {
		helper.HttpError(w, http.StatusBadRequest, "email is required")
		return
	}

	var input dto.CreateProfileRequest
	input.UserId = uint(id)
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userUC.CreateProfile(&input); err != nil {
		switch err {
		case gorm.ErrDuplicatedKey:
			helper.HttpError(w, http.StatusConflict, err.Error())
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
			return
		}
		helper.HttpWriter(w, http.StatusOK, nil)
	}
}

// Change Password godoc
// @Summary Change password by email.
// @Description This endpoint for user want ti change password in app.
// @Param id path integer true "User ID"
// @Tags User
// @Accept json
// @Produce json
// @param request body dto.UserChangePasswordRequest true "change password"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /user/change-password [post]
// @Security BearerAuth
func (h *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var input dto.UserChangePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	input.Email = claims.Email
	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if input.NewPassword != input.ConfirmNewPassword {
		helper.HttpError(w, http.StatusBadRequest, "new password and confirm password do not match")
		return
	}

	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userUC.ChangePassword(&input); err != nil {
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

// Update Profile godoc
// @Summary Update profil.
// @Description This endpoint for user update their profile.
// @Tags User
// @Accept json
// @Produce json
// @param request body dto.UpdateProfileRequest true "update profile"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /user/update [patch]
// @Security BearerAuth
func (h *UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateProfileRequest

	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	input.UserId = claims.UserID
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userUC.UpdateProfile(&input); err != nil {
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

// Update Role godoc
// @Summary Update Role.
// @Description This endpoint for user update role, the old role must a consumer and can update only one chance, the update role is a farmer, distributor, retailer.
// @Tags User
// @Accept json
// @Produce json
// @param request body dto.UpdateRoleRequest true "update role"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /user/role [patch]
// @Security BearerAuth
func (h *UserHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateRoleRequest

	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	input.UserId = claims.UserID
	input.OldRole = entity.Status(claims.Role)
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userUC.UpdateRole(&input); err != nil {
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
