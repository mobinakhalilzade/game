package httpserver

import (
	"fmt"
	"gameapp/config"
	"gameapp/service/authservice"
	"gameapp/service/userservice"
	"gameapp/validator/uservalidator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config        config.Config
	authSvc       authservice.Service
	userSvc       userservice.Service
	uservalidator uservalidator.Validator
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service, uservalidator uservalidator.Validator) Server {
	return Server{
		config:        config,
		authSvc:       authSvc,
		userSvc:       userSvc,
		uservalidator: uservalidator,
	}
}

func (s Server) Serve() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/health-check", s.healthCheck)

	// Users
	userGroup := e.Group("/users")

	userGroup.POST("/register", s.userRegister)
	userGroup.POST("/login", s.userLogin)
	userGroup.GET("/profile", s.userProfile)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))
}
