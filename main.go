package main

import (
	"crm_go/cache/redis"
	"crm_go/db/postgres"
	"crm_go/delivery/httpserver"
	"crm_go/delivery/httpserver/handlers/authHandler"
	"crm_go/pkg/validation"
	"crm_go/repositories/userRepository"
	"crm_go/services/authService"
	"log"
	"time"

	_ "crm_go/docs"
)

//	@title			CRM Go API
//	@version		1.0
//	@description	A CRM backend API built with Go and Echo framework

//	@host		localhost:8099
//	@BasePath	/api

func main() {
	validation.Init()

	db, err := postgres.New(postgres.Config{
		Host:     "localhost",
		Username: "crm_user",
		Password: "crm_password",
		Database: "crm",
		SSLMode:  "disable",
	})

	cache := redis.New(redis.Config{
		Host:     "localhost",
		Port:     "6379",
		Password: "",
		Db:       0,
	})

	authConfig := authService.Config{
		JWTSecret:  []byte("sharif_secret"),
		AccessTTL:  15 * time.Minute,
		RefreshTTL: 30 * 24 * time.Hour,
		Issuer:     "crm_go",
	}
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	repo := userRepository.New(db)
	svc := authService.New(repo, *cache, authConfig)
	h := authHandler.New(svc)

	srv := httpserver.New(h)
	srv.Start()
}
