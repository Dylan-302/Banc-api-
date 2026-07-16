package handlers

import (
	"strconv"

	accountusecase "banc-api/src/modules/account/application/usecase"
	"banc-api/src/modules/account/domain/repositories"

	"github.com/gofiber/fiber/v3"
)

type UpdateAccountHandler struct {
	repo repositories.AccountRepository
}

func NewUpdateAccountHandler(repo repositories.AccountRepository) *UpdateAccountHandler {
	return &UpdateAccountHandler{repo: repo}
}

func (h *UpdateAccountHandler) Handle(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El ID proporcionado no es válido",
		})
	}

	var body accountusecase.UpdateAccountDTO
	if err := c.Bind().Body(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos de actualización inválidos",
		})
	}

	res, err := accountusecase.UpdateAccountUsecase(uint(id), body, h.repo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al actualizar la cuenta",
		})
	}
	if res == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Cuenta no encontrada",
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
