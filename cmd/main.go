package main

import (
	"github.com/IvanDrf/units/internal/database"
	"github.com/IvanDrf/units/internal/handlers"
)

func main() {
	database.InitDB()
	server := handlers.InitServer()

	server.RegisterRoutes()
	server.Start()
}
