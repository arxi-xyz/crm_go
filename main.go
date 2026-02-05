package main

import "crm_go/delivery/http"

func main() {
	httpServer := http.NewServer()
	httpServer.Start()
}
