package handlers

import (
	"banc-api/src/modules/user/application/dto/request"
	"banc-api/src/modules/user/application/usecase"
	"banc-api/src/modules/user/domain/repositories"

	"github.com/gofiber/fiber/v3"
)

type CreateUserHandler struct {
	repo repositories.UserRepository
}

func NewCreateUserHandler(repo repositories.UserRepository) *CreateUserHandler {
	if repo == nil {
		return nil
	}
	return &CreateUserHandler{repo: repo}
}

func (h *CreateUserHandler) Handle(c fiber.Ctx) error {
	var req request.CreateUserRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error al procesar los datos de entrada",
		})
	}
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Todos los campos son requeridos",
		})
	}

	newUser, err := usecase.CreateUserUsecase(h.repo, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newUser)
}
