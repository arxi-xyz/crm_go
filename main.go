package main

import (
	"crm_go/db/postgres"
	"crm_go/delivery/httpserver"
	"crm_go/delivery/httpserver/handlers/authHandler"
	"crm_go/pkg/validation"
	"crm_go/repositories/userRepository"
	"crm_go/services/authService"
	"log"
)

func main() {
	validation.Init()

	db, err := postgres.New(postgres.Config{
		Host:     "localhost",
		Username: "crm_user",
		Password: "crm_password",
		Database: "crm",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := userRepository.New(db)
	svc := authService.New(repo)
	h := authHandler.New(svc)

	srv := httpserver.New(h)
	srv.Start()
}
