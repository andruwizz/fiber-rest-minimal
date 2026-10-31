package main

import (
	"go-something/internal/config"
	"log"
)

func main() {
	// Init Database
	db := config.NewDatabase()

	// Init Fiber App
	app := config.NewFiberApp()

	config.Bootstrap(&config.BootstrapConfig{
		DB:  db,
		App: app,
	})

	log.Fatal(app.Listen(":3000"))
}
