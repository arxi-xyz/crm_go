package http

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.GET("/ping", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	if err := e.Start(":8099"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
