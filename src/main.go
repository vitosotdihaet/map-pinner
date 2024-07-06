package main

import (
	"log"

	"github.com/vitosotdihaet/map-pinner/package/handlers"
	"github.com/vitosotdihaet/map-pinner/package/server"
)

func main() {
	handler := new(handlers.Handler)
	server := new(server.Server)

	if err := server.Run("8080", handler.InitEndpoints()); err != nil {
		log.Fatalf("error while running the server: %s", err.Error())
	}
}
