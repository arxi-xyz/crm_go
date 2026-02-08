package httpserver

import (
	"crm_go/pkg/httpx"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Server struct {
	authHandler authHandlerInterface
}

func New(authHandler authHandlerInterface) *Server {
	return &Server{
		authHandler: authHandler,
	}
}

type authHandlerInterface interface {
	Login(c *echo.Context) error
	Logout(c *echo.Context) error
}

func (s *Server) Start() {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.HTTPErrorHandler = httpx.NewErrorHandler()

	api := e.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/login", s.authHandler.Login)
	auth.GET("/logout", s.authHandler.Logout)

	e.GET("/api/ping", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})

	if err := e.Start(":8099"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}

}
