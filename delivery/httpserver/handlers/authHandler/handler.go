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

func New() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Logout(c *echo.Context) error {

	return c.String(http.StatusOK, "Logout")
}

func (h *AuthHandler) Login(c *echo.Context) error {

	req := authService.LoginRequest{
		Phone:    c.FormValue("phone"),
		Password: c.FormValue("password"),
	}

	resp, _ := h.AuthService.Login(req)

	return c.JSON(http.StatusOK, resp)
}
