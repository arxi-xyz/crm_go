package httpserver

import (
	"crm_go/delivery/httpserver/handlers/authHandler"
	"crm_go/pkg/httpx"
	"crm_go/repositories/userRepository"
	"crm_go/services/authService"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (s *Server) Start() {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.HTTPErrorHandler = httpx.NewErrorHandler()
	
	api := e.Group("/api")

	handler := authHandler.AuthHandler{
		AuthService: authService.New(userRepository.UserRepository{}),
	}
	auth := api.Group("/auth")
	auth.POST("/login", handler.Login)
	auth.GET("/logout", handler.Logout)

	e.GET("/api/ping", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})

	if err := e.Start(":8099"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}

}
