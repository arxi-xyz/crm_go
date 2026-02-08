package authHandler

import (
	"crm_go/services/authService"
	"net/http"

	"github.com/labstack/echo/v5"
)

type AuthHandler struct {
	AuthService authServiceInterface
}

type authServiceInterface interface {
	Login(request authService.LoginRequest) (authService.LoginResponse, error)
}

func New(authService authServiceInterface) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Logout(c *echo.Context) error {

	return c.String(http.StatusOK, "Logout")
}

func (h *AuthHandler) Login(c *echo.Context) error {
	var req authService.LoginRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid_body")
	}

	resp, err := h.AuthService.Login(req)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
