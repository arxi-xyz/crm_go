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

// @Summary		Logout user
// @Description	Invalidate user's refresh token session
// @Tags			Auth
// @Produce		json
// @Param			Authorization	header		string	true	"Refresh Token"
// @Success		200				{object}	map[string]string
// @Failure		400				{object}	httpx.ErrorResponse
// @Failure		401				{object}	httpx.ErrorResponse
// @Router			/auth/logout [get]
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

// @Summary		Login user
// @Description	Authenticate user with phone number and password
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			request	body		authService.LoginRequest	true	"Login credentials"
// @Success		200		{object}	authService.LoginResponse
// @Failure		400		{object}	httpx.ErrorResponse
// @Failure		401		{object}	httpx.ErrorResponse
// @Failure		500		{object}	httpx.ErrorResponse
// @Router			/auth/login [post]
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

// @Summary		Refresh tokens
// @Description	Get new access and refresh tokens using a valid refresh token
// @Tags			Auth
// @Produce		json
// @Param			Authorization	header		string	true	"Refresh Token"
// @Success		200				{object}	authService.RefreshResponse
// @Failure		400				{object}	httpx.ErrorResponse
// @Failure		401				{object}	httpx.ErrorResponse
// @Router			/auth/refresh [post]
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

func (h AuthHandler) SetRoutes(g *echo.Group) {
	auth := g.Group("/auth")

	auth.POST("/login", h.Login)
	auth.POST("/refresh", h.Refresh)
	auth.GET("/logout", h.Logout)
}
