package handlers

import (
	"strconv"

	"banc-api/src/modules/user/application/usecase"
	"banc-api/src/modules/user/domain/repositories"

	"github.com/gofiber/fiber/v3"
)

type UpdateUserHandler struct {
	repo repositories.UserRepository
}

func NewUpdateUserHandler(repo repositories.UserRepository) *UpdateUserHandler {
	return &UpdateUserHandler{repo: repo}
}

func (h *UpdateUserHandler) Handle(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El ID proporcionado no es válido",
		})
	}

	var body usecase.UpdateUserDTO
	if err := c.Bind().Body(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos de actualización inválidos",
		})
	}

	res, err := usecase.UpdateUserUsecase(uint(id), body, h.repo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al actualizar el usuario",
		})
	}
	if res == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Usuario no encontrado",
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}