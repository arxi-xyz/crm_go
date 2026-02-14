package httpx

import (
	"crm_go/pkg/appError"
	"net/http"

	"github.com/labstack/echo/v5"
)

type ErrorResponse struct {
	Code    string         `json:"code" example:"validation_error"`
	Message string         `json:"message" example:"invalid input"`
	Meta    map[string]any `json:"meta,omitempty"`
}

func NewErrorHandler() echo.HTTPErrorHandler {
	return func(c *echo.Context, err error) {
		if ae, ok := appError.AsAppError(err); ok {
			write(c, ae)
			return
		}

		if ae, ok := appError.FromValidator(err); ok {
			write(c, ae)
			return
		}

		if he, ok := err.(*echo.HTTPError); ok {

			ae := appError.New(he.Code, "http_error", http.StatusText(he.Code), err, map[string]any{
				"detail": he.Message,
			})
			write(c, ae)
			return
		}

		write(c, appError.Internal(err))
	}
}

func write(c *echo.Context, ae *appError.AppError) {
	if r, _ := echo.UnwrapResponse(c.Response()); r != nil && r.Committed {
		return
	}
	_ = c.JSON(ae.Status, map[string]any{
		"code":    ae.Code,
		"message": ae.Message,
		"meta":    ae.Meta,
	})
}
