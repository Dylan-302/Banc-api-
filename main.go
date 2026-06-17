package main

import (
	"log"

	config "banc-api/src/common/config"
	adapter "banc-api/src/infrastructure/db/adapter"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Inicializar configuración
	cfg := config.NewConfig()

	// Conectar a la base de datos
	_ = adapter.NewDBConnection(cfg)

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("hello world")
	})

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
