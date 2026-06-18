package main

import (
	"banc-api/src/infrastructure/server"
	"log"
)

func main() {
	app := server.NewApp()
	if err := app.Run(); err != nil {
		log.Fatalf("Error al arrancar el servidor: %v", err)
	}
}
