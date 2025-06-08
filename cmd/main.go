package main

import (
	"log"

	"github.com/IvanDrf/units/internal/database"
	"github.com/IvanDrf/units/internal/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()
	database.InitDB()

	server.POST("/", handlers.PostHandler)
	server.GET("/", handlers.GetHandler)
	server.PATCH("/conversions/:id", handlers.PatchHandler)
	server.DELETE("/conversions/:id", handlers.DeleteHandler)

	if err := server.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
