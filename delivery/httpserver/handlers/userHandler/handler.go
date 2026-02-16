package userHandler

import (
	"crm_go/services/userService"
	"net/http"

	"github.com/labstack/echo/v5"
)

type UserHandler struct {
	UserService userServiceInterface
}

type userServiceInterface interface {
	GetMe(uuid string) (userService.MeResponse, error)
}

func New(userService userServiceInterface) *UserHandler {
	return &UserHandler{UserService: userService}
}

// @Summary		Get current user
// @Description	Returns the authenticated user's profile information
// @Tags			User
// @Produce		json
// @Param			Authorization	header		string	true	"Access Token"
// @Success		200				{object}	userService.MeResponse
// @Failure		401				{object}	httpx.ErrorResponse
// @Failure		404				{object}	httpx.ErrorResponse
// @Failure		500				{object}	httpx.ErrorResponse
// @Router			/me [get]
func (h *UserHandler) Me(c *echo.Context) error {
	userUUID, ok := c.Get("user_uuid").(string)
	if !ok || userUUID == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "missing user identity")
	}

	resp, err := h.UserService.GetMe(userUUID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
