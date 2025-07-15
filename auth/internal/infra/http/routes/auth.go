package routes

import (
	"encoding/json"
	"net/http"

	"github.com/dawit_hopes/saas/auth/internal/domain/model"
	"github.com/dawit_hopes/saas/auth/internal/domain/port/inbound"
	"github.com/dawit_hopes/saas/auth/internal/infra/http/dto"
	"github.com/dawit_hopes/saas/auth/internal/infra/http/utils"
	"github.com/gin-gonic/gin"

	appErr "github.com/dawit_hopes/saas/auth/internal/common/errors"
)

// AuthHandler handles HTTP requests for authentication
type AuthHandler struct {
	authService inbound.AuthService
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(authService inbound.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func RegisterAuthRoutes(rg *gin.RouterGroup, authService inbound.AuthService) {
	handler := NewAuthHandler(authService)

	rg.POST("/register", wrap(handler.Register))
	rg.POST("/login", wrap(handler.Login))
	rg.POST("/refresh-token", wrap(handler.RefreshToken))
	rg.POST("/me", wrap(handler.Me))
}

func wrap(f func(http.ResponseWriter, *http.Request)) gin.HandlerFunc {
	return func(c *gin.Context) {
		f(c.Writer, c.Request)
	}
}

// Register handles user registration
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user dto.SignUp

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendErrorResponse(w, *appErr.ErrInvalidJSONPayload)
		return
	}

	token, err := h.authService.Signup(r.Context(), model.User{
		Email:    user.Email,
		Password: user.Password,
		Location: user.Location,
	})
	if err != nil {
		utils.SendErrorResponse(w, *err)
		return
	}

	utils.SendSuccessResponse(w, dto.AuthResponse{
		StatusCode: http.StatusOK,
		Token:      token,
	})
}

// Login handles user authentication
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user dto.Login

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendErrorResponse(w, *appErr.ErrInvalidJSONPayload)
		return
	}

	if err := user.Validation(); err != nil {
		utils.SendErrorResponse(w, *err)
		return
	}

	token, err := h.authService.Login(r.Context(), user.Email, user.Password)
	if err != nil {
		utils.SendErrorResponse(w, *err)
		return
	}

	utils.SendSuccessResponse(w, dto.AuthResponse{
		StatusCode: http.StatusOK,
		Token:      token,
	})
}

// RefreshToken handles access token refresh
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {

	var body dto.RefreshToken
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendErrorResponse(w, *appErr.ErrInvalidJSONPayload)
		return
	}

	token, err := h.authService.RefreshToken(r.Context(), body.RefreshToken)
	if err != nil {
		utils.SendErrorResponse(w, *err)
		return
	}

	utils.SendSuccessResponse(w, dto.AuthResponse{
		StatusCode: http.StatusOK,
		Token:      token,
	})
}

// Me handles fetching the current user information
func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	var body dto.AccessToken
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendErrorResponse(w, *appErr.ErrInvalidJSONPayload)
		return
	}

	usr, err := h.authService.Me(r.Context(), body.AccessToken)
	if err != nil {
		utils.SendErrorResponse(w, *err)
		return
	}

	utils.SendSuccessResponse(w, dto.UserResponse{
		StatusCode: http.StatusOK,
		User:       usr,
	})
}
