package handlers

import (
	"banc-api/src/modules/account/application/usecase"
	"banc-api/src/modules/account/domain/repositories"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

// DeleteAccountHandler es un handler para eliminar cuentas.
type DeleteAccountHandler struct {
	repo repositories.AccountRepository
}

// NewDeleteAccountHandler crea una nueva instancia del handler.
func NewDeleteAccountHandler(repo repositories.AccountRepository) *DeleteAccountHandler {
	if repo == nil {
		return nil
	}

	return &DeleteAccountHandler{repo}
}

// Handle maneja la solicitud de eliminación de cuenta.
func (h *DeleteAccountHandler) Handle(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid account ID",
		})
	}

	// Ejecutar el caso de uso
	if err := usecase.DeleteAccountUsecase(uint(id), h.repo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete account",
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
