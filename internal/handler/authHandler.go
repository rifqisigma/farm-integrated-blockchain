package handler

import (
	"encoding/json"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/internal/usecase"
	"farm-integrated-web3/utils/helper"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authUC    usecase.AuthUsecase
	validator *validator.Validate
}

func NewAuthHandler(authUC usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUC, validator.New()}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input dto.Register
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.authUC.Register(&input); err != nil {
		helper.HttpError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.HttpWriter(w, http.StatusCreated, map[string]string{"message": "user registered successfully, please check your email to verify your account"})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input dto.Login
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	shortToken, longToken, err := h.authUC.Login(&input)
	if err != nil {
		if err == helper.ErrLoginNotSuccess {
			helper.HttpError(w, http.StatusUnauthorized, err.Error())
			return
		}
		helper.HttpError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]string{
		"access_token":  longToken,
		"refresh_token": shortToken,
	}

	helper.HttpWriter(w, http.StatusOK, response)
}

func (h *AuthHandler) ValidateUser(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		helper.HttpError(w, http.StatusBadRequest, "token is required")
		return
	}

	if err := h.authUC.ValidateUser(token); err != nil {
		switch err {
		case helper.ErrUserNotFound:
			helper.HttpError(w, http.StatusNotFound, "token is required")
		case helper.ErrInvalidToken:
			helper.HttpError(w, http.StatusBadRequest, "token is expired")
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, map[string]string{"message": "user validated successfully"})
}

func (h *AuthHandler) RefreshLongToken(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		helper.HttpError(w, http.StatusBadRequest, "token is required")
		return
	}

	newToken, err := h.authUC.RefreshLongToken(token)
	if err != nil {
		switch err {
		case helper.ErrInvalidToken:
			helper.HttpError(w, http.StatusBadRequest, "token is expired")
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	response := map[string]string{
		"access_token": newToken,
	}
	helper.HttpWriter(w, http.StatusOK, response)
}

func (h *AuthHandler) RequestResetPassword(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		helper.HttpError(w, http.StatusBadRequest, "email is required")
		return
	}

	if err := h.authUC.RequestResetPassword(email); err != nil {
		switch err {
		case helper.ErrUserNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, map[string]string{"message": "please check your email to reset your password"})
}

func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var input dto.UserNewPassword
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.authUC.ResetPassword(&input); err != nil {
		switch err {
		case helper.ErrInvalidToken:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		case helper.ErrBadRequest:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		case helper.ErrUserNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, map[string]string{"message": "password has been reset successfully"})
}
