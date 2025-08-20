package main

import (
	"github.com/IvanDrf/units/internal/handlers"
)

func main() {
	server := handlers.InitServer()

	server.RegisterRoutes()
	server.Start()
}
