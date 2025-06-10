package main

import (
	"sync"

	"github.com/IvanDrf/units/internal/database"
	"github.com/IvanDrf/units/internal/handlers"
)

func main() {
	server := new(handlers.Server)
	wg := new(sync.WaitGroup)

	wg.Add(2)
	go func() {
		defer wg.Done()
		server = handlers.InitServer()
	}()

	go func() {
		defer wg.Done()
		database.InitDB()
	}()

	wg.Wait()

	server.RegisterRoutes()
	server.Start()
}
