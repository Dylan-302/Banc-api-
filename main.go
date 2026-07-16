package main

import (
	"errors"
	"log"

	config "banc-api/src/common/config"
	adapter "banc-api/src/infraestructure/db/adapter"
	dao "banc-api/src/infraestructure/db/dao"
	userrepo "banc-api/src/modules/user/domain/repositories"
	userhandlers "banc-api/src/modules/user/handlers"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func errorHandler(c fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		code = fiberErr.Code
	}
	return c.Status(code).JSON(fiber.Map{"isError": true, "message": err.Error()})
}

func main() {
	cfg := config.NewConfig()
	db := adapter.NewDBConnection(cfg)

	var repo userrepo.UserRepository
	if db.DB != nil {
		repo = dao.NewUserDao(db.DB)
	} else {
		repo = dao.NewInMemoryUserRepository()
	}

	createUserHandler := userhandlers.NewCreateUserHandler(repo)
	findByIDHandler := userhandlers.NewFindByIdUserHandler(repo)
	getAllUsersHandler := userhandlers.NewGetAllUsersHandler(repo)
	updateUserHandler := userhandlers.NewUpdateUserHandler(repo)
	deleteUserHandler := userhandlers.NewDeleteUserHandler(repo)

	app := fiber.New(fiber.Config{ErrorHandler: errorHandler})

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	api := app.Group("/api/v1/users")
	api.Post("/", createUserHandler.Handle)
	api.Get("/", getAllUsersHandler.Handle)
	api.Get("/:id", findByIDHandler.Handle)
	api.Put("/:id", updateUserHandler.Handle)
	api.Delete("/:id", deleteUserHandler.Handle)

	log.Printf("Starting Fiber server on port %s", cfg.App.Port)
	log.Fatal(app.Listen(":" + cfg.App.Port))
}
