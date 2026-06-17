package infraestructure

import "github.com/gofiber/fiber/v2"

func JWTMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func AuthMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func AdminMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func UserMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func GuestMiddleware(c *fiber.Ctx) error {
	return c.Next()
}
