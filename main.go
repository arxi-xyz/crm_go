package main

import (
	"crm_go/delivery/httpserver"
	"crm_go/pkg/validation"
)

func main() {
	validation.Init()

	httpServer := httpserver.New()
	httpServer.Start()
}
