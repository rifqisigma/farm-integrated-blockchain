package handler

import (
	"encoding/json"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/internal/usecase"
	"farm-integrated-web3/utils/helper"
	"farm-integrated-web3/utils/middleware"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authUC    usecase.AuthUsecase
	validator *validator.Validate
}

func NewAuthHandler(authUC usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUC, validator.New()}
}

// Register godoc
// @Summary Register for new user
// @Description This endpoint for registery new user.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Register Data"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 409 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /auth/gmail/register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.authUC.Register(&input); err != nil {
		switch err {
		case gorm.ErrDuplicatedKey:
			helper.HttpError(w, http.StatusConflict, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}

		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)
}

// Login godoc
// @Summary Login user
// @Description This endpoint for user login, get the access and refresh token.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login Data"
// @Success 200 {object} dto.ResponseLogin
// @Failure 400 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /auth/gmail/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	shortToken, longToken, err := h.authUC.Login(&input)
	if shortToken == "" {
		helper.HttpError(w, http.StatusInternalServerError, "short err")
		return
	}

	if longToken == "" {
		helper.HttpError(w, http.StatusInternalServerError, "long err")
		return
	}
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, "user not found or not verified")
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, dto.ResponseLogin{
		AccessToken:  shortToken,
		RefreshToken: longToken,
	})
}

// Validate User godoc
// @Summary Validation User.
// @Description This endpoint will send on your email and for validating user.
// @Tags Auth
// @Accept json
// @Produce json
// @Param token query string true "Token validasi user"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /auth/gmail/verification [get]
func (h *AuthHandler) ValidateUser(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		helper.HttpError(w, http.StatusBadRequest, "token is required")
		return
	}

	if err := h.authUC.ValidateUser(token); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, "token is required")
		case helper.ErrInvalidToken:
			helper.HttpError(w, http.StatusBadRequest, "token is expired")
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, dto.ResponseMessage{
		Message: "validate the new success",
	})
}

// New Refresh Token godoc
// @Summary New Refresh Token.
// @Description This endpoint for get new refresh token, for get new Refresh token you must have  a valid refresh token .
// @Tags Auth
// @Accept json
// @Produce json
// @Param token query string true "Refresh token"
// @Success 200 {object} dto.ResponseRefreshToken
// @Failure 400 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /auth/gmail/refresh-token [post]
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

	helper.HttpWriter(w, http.StatusOK, dto.ResponseRefreshToken{
		RefreshToken: newToken,
	})
}

// Request reset password godoc
// @Summary Requset reset password.
// @Description This endpoint for user request send link reset password at email.
// @Tags Auth
// @Accept js/ @Produce json
// @Param email query string true "query email"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /auth/gmail/forgot-password [post]
func (h *AuthHandler) RequestResetPassword(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		helper.HttpError(w, http.StatusBadRequest, "email is required")
		return
	}

	if err := h.authUC.RequestResetPassword(email); err != nil {
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

// Reset password godoc
// @Summary Reset Password.
// @Description This endpoint for user request send link reset password at email.
// @Tags Auth
// @Accept js/ @Produce json
// @Param request body dto.UserResetPasswordRequest true "Reset password"
// @Param  token query string true "token"
// @Param email query string true "email"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 401 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /auth/gmail/reset-password [post]
func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		helper.HttpError(w, http.StatusBadRequest, "email is required")
		return
	}

	token := r.URL.Query().Get("email")
	if token == "" {
		helper.HttpError(w, http.StatusBadRequest, "email is required")
		return
	}

	var input dto.UserResetPasswordRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	input.Email = email
	input.Token = token
	if err := h.validator.Struct(&input); err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.authUC.ResetPassword(&input); err != nil {
		switch err {
		case helper.ErrInvalidToken:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		case helper.ErrBadRequest:
			helper.HttpError(w, http.StatusBadRequest, err.Error())
		case gorm.ErrRecordNotFound:
			helper.HttpError(w, http.StatusNotFound, err.Error())
		default:
			helper.HttpError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	helper.HttpWriter(w, http.StatusOK, nil)
}

// Resend verification user godoc
// @Summary Reset Password.
// @Description This endpoint for user request resend link reset password in email if the user is late in verifying the account.
// @Tags Auth
// @Accept js/ @Produce json
// @Param email query string true "query email"
// @Success 200
// @Failure 400 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /auth/gmail/resend-verification [post]
func (h *AuthHandler) ResendVerificationEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		helper.HttpError(w, http.StatusBadRequest, "email is required")
		return
	}

	if err := h.authUC.ResendVerificationEmail(email); err != nil {
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

// Create new access token godoc
// @Summary Create new access token.
// @Description This endpoint to get new access token, you must have a valid refresh token.
// @Tags Auth
// @Accept json
// @Produce json
// @Failure 401 {object} dto.ResponseError
// @Success 200 {object} dto.ResponseAccessToken
// @Failure 400 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /auth/access-token [post]
// @Security BearerAuth
func (h *AuthHandler) CreateAccessToken(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaimsLongExp)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	token, err := h.authUC.CreateAccessToken(claims.UserID)
	if err != nil {
		helper.HttpError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.HttpWriter(w, http.StatusOK, dto.ResponseAccessToken{
		AccessToken: token,
	})
}

// Logout Token godoc
// @Summary Logout.
// @Description This endpoint for Logout.
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /auth/logout [post]
// @Security BearerAuth
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := h.authUC.Logout(claims.UserID); err != nil {
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

// Delete Account Token godoc
// @Summary Delete Account.
// @Description This endpoint for delete user.
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200
// @Failure 401 {object} dto.ResponseError
// @Failure 404 {object} dto.ResponseError
// @Failure 403 {object} dto.ResponseError
// @Failure 500 {object} dto.ResponseError
// @Router /auth/delete-account [delete]
// @Security BearerAuth
func (h *AuthHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middleware.UserContextKey).(*helper.JWTclaims)
	if !ok {
		helper.HttpError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	if err := h.authUC.DeleteAccount(claims.UserID); err != nil {
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
