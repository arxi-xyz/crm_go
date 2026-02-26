package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func Authorization(service AuthorizationServiceInterface, perm string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {

			userUuid, exists := c.Get("user_uuid").(string)

			if !exists {
				return c.JSON(http.StatusForbidden, map[string]string{
					"error": "forbidden",
				})
			}

			has, err := service.HasPermission(userUuid, perm)
			if err != nil {
				return c.JSON(http.StatusForbidden, map[string]string{
					"error": "forbidden",
				})
			}

			if !has {
				return c.JSON(http.StatusForbidden, map[string]string{
					"error": "forbidden",
				})
			}

			return next(c)
		}
	}
}
