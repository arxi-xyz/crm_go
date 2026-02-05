package main

import "crm_go/delivery/httpserver"

func main() {
	httpServer := httpserver.New()
	httpServer.Start()
}
