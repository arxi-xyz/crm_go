package main

import (
	"crm_go/adaptor/redis"
	"crm_go/db/postgres"
	"crm_go/delivery/httpserver"
	"crm_go/delivery/httpserver/handlers/authHandler"
	"crm_go/delivery/httpserver/handlers/userHandler"
	"crm_go/delivery/httpserver/middlewares"
	"crm_go/pkg/validation"
	"crm_go/repositories/userRepository"
	"crm_go/services/authService"
	"crm_go/services/userService"
	"log"
	"time"

	_ "crm_go/docs"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
)

var k = koanf.New(".")

//	@title			CRM Go API
//	@version		1.0
//	@description	A CRM backend API built with Go and Echo framework

//	@host		localhost:8099
//	@BasePath	/api

func main() {
	validation.Init()

	if err := k.Load(file.Provider("./.env"), dotenv.Parser()); err != nil {
		log.Println("no .env file found, reading from environment variables")
	}

	k.Load(env.Provider("", ".", func(s string) string {
		return s
	}), nil)

	db, err := postgres.New(postgres.Config{
		Host:     k.String("DB_HOST"),
		Port:     k.Int("DB_PORT"),
		Username: k.String("DB_USER"),
		Password: k.String("DB_PASSWORD"),
		Database: k.String("DB_NAME"),
		SSLMode:  k.String("DB_SSL_MODE"),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cache := redis.New(redis.Config{
		Host:     k.String("REDIS_HOST"),
		Port:     k.String("REDIS_PORT"),
		Password: k.String("REDIS_PASSWORD"),
		Db:       k.Int("REDIS_DB"),
	})

	authConfig := authService.Config{
		JWTSecret:  []byte(k.String("JWT_SECRET")),
		AccessTTL:  time.Duration(k.Int("JWT_ACCESS_TTL_MINUTES")) * time.Minute,
		RefreshTTL: time.Duration(k.Int("JWT_REFRESH_TTL_DAYS")) * 24 * time.Hour,
		Issuer:     k.String("JWT_ISSUER"),
	}

	repo := userRepository.New(db)

	authSvc := authService.New(repo, cache, authConfig)
	userSvc := userService.New(repo)

	authH := authHandler.New(authSvc)
	userH := userHandler.New(userSvc)
	authMw := middlewares.Auth(authSvc)

	srv := httpserver.New(authH, userH, authMw)
	srv.Start(":" + k.String("SERVER_PORT"))
}
