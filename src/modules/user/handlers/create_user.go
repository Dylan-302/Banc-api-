package handlers

import (
	"banc-api/src/modules/user/application/dto/request"
	"banc-api/src/modules/user/application/mapper"
	"banc-api/src/modules/user/application/usecase"
	"banc-api/src/modules/user/domain/repositories"
	"time"

	"github.com/gofiber/fiber/v3"
)

// CreateUserHandler es un handler para crear nuevos usuarios.
type CreateUserHandler struct {
	repo repositories.UserRepository
}

// NewCreateUserHandler crea una nueva instancia del handler.
func NewCreateUserHandler(repo repositories.UserRepository) *CreateUserHandler {
	return &CreateUserHandler{repo}
}

// Handle maneja la solicitud de creación de usuario.
func (h *CreateUserHandler) Handle(c fiber.Ctx) error {
	var req request.CreateUserRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Mapear el request a la entidad
	user := mapper.CreateUserRequestToEntity(req)
	user.FechadeCreacion = time.Now()

	// Ejecutar el caso de uso
	createdUser, err := usecase.CreateUserUsecase(user, h.repo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	// Mapear la respuesta
	response := mapper.UserEntityToResponse(*createdUser)

	return c.Status(fiber.StatusCreated).JSON(response)
}
