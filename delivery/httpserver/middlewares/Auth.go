package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func Auth(service authServiceInterface) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "missing token",
				})
			}

			claims, err := service.ValidateAccessToken(authHeader)
			if err != nil {
				return err
			}

			c.Set("user_uuid", claims.Subject)
			c.Set("claims", claims)

			return next(c)
		}
	}
}
