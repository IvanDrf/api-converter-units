package main

import (
	"log"

	"github.com/IvanDrf/units/internal/database"
	"github.com/IvanDrf/units/internal/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	server.POST("/", handlers.PostHandler)
	server.GET("/", handlers.GetHandler)

	database.InitDB()

	if err := server.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
