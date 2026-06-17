package handlers

import (
	"banc-api/src/modules/user/application/dto/request"
	"banc-api/src/modules/user/application/mapper"
	"banc-api/src/modules/user/application/usecase"
	"banc-api/src/modules/user/domain/entities"
	"banc-api/src/modules/user/domain/repositories"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type UpdateUserHandler struct {
	repo repositories.UserRepository
}

func NewUpdateUserHandler(repo repositories.UserRepository) *UpdateUserHandler {
	return &UpdateUserHandler{repo}
}

func (u *UpdateUserHandler) Handle(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var req request.UpdateUserRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Crear entidad de usuario a actualizar
	userToUpdate := entities.User{
		ID:       uint(id),
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	// Ejecutar el caso de uso
	updatedUser, err := usecase.UpdateUserUsecase(userToUpdate, u.repo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	// Mapear la respuesta
	response := mapper.UserEntityToResponse(*updatedUser)

	return c.Status(fiber.StatusOK).JSON(response)
}
