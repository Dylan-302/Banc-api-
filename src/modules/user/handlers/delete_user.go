package handlers

import (
	"banc-api/src/modules/user/application/usecase"
	"banc-api/src/modules/user/domain/repositories"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// DeleteUserHandler es un handler para eliminar usuarios.
type DeleteUserHandler struct {
	repo repositories.UserRepository
}

// NewDeleteUserHandler crea una nueva instancia del handler.
func NewDeleteUserHandler(repo repositories.UserRepository) *DeleteUserHandler {
	if repo == nil {
		return nil
	}

	return &DeleteUserHandler{repo}
}

// Handle maneja la solicitud de eliminación de usuario.
func (h *DeleteUserHandler) Handle(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	// Ejecutar el caso de uso
	if err := usecase.DeleteUserUsecase(uint(id), h.repo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
