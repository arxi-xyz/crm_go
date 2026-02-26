package httpserver

import (
	"crm_go/delivery/httpserver/handlers/authHandler"
	"crm_go/delivery/httpserver/handlers/userHandler"
	"crm_go/delivery/httpserver/middlewares"
	"crm_go/pkg/httpx"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Server struct {
	authHandler          *authHandler.AuthHandler
	userHandler          *userHandler.UserHandler
	authMiddleware       echo.MiddlewareFunc
	authorizationService middlewares.AuthorizationServiceInterface
}

func New(
	authHandler *authHandler.AuthHandler,
	userHandler *userHandler.UserHandler,
	authMiddleware echo.MiddlewareFunc,
	authorizationService middlewares.AuthorizationServiceInterface,
) *Server {
	return &Server{
		authHandler:          authHandler,
		userHandler:          userHandler,
		authMiddleware:       authMiddleware,
		authorizationService: authorizationService,
	}
}

func (s *Server) Start(addr string) {
	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
		},
		MaxAge: 86400,
	}))

	e.HTTPErrorHandler = httpx.NewErrorHandler()

	// Swagger
	swaggerHandler := httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	)
	e.GET("/swagger/*", func(c *echo.Context) error {
		swaggerHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	api := e.Group("/api")

	s.authHandler.SetRoutes(api)

	protected := api.Group("", s.authMiddleware)
	s.userHandler.SetRoutes(protected, s.authorizationService)

	api.GET("/ping", s.Ping)

	if err := e.Start(addr); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

// Ping @Summary		Health check
// @Description	Returns pong to verify the API is running
// @Tags			Health
// @Produce		json
// @Success		200	{object}	map[string]string
// @Router			/ping [get]
func (s *Server) Ping(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "pong"})
}
