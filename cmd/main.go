package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	if err := server.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
