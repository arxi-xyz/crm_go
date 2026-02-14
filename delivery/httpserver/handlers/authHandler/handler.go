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
	Refresh(request authService.RefreshRequest) (authService.RefreshResponse, error)
	Logout(request authService.LogoutRequest) error
}

func New(authService authServiceInterface) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Logout(c *echo.Context) error {
	var req authService.LogoutRequest

	req.RefreshToken = c.Request().Header.Get("Authorization")

	if req.RefreshToken == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing_authorization")
	}

	err := h.AuthService.Logout(req)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{})
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

func (h *AuthHandler) Refresh(c *echo.Context) error {
	// todo: duplicate code
	var req authService.RefreshRequest

	req.RefreshToken = c.Request().Header.Get("Authorization")

	if req.RefreshToken == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing_authorization")
	}

	resp, err := h.AuthService.Refresh(req)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
