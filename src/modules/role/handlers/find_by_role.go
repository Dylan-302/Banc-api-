package handlers

import (
	"strconv"

	"banc-api/src/modules/role/domain/repositories"
	roleusecase "banc-api/src/modules/role/domain/usecases"

	"github.com/gofiber/fiber/v3"
)

type UpdateRoleHandler struct {
	repo repositories.RoleRepository
}

func NewUpdateRoleHandler(repo repositories.RoleRepository) *UpdateRoleHandler {
	return &UpdateRoleHandler{repo: repo}
}

func (h *UpdateRoleHandler) Handle(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El ID proporcionado no es válido",
		})
	}

	var body roleusecase.UpdateRoleDTO
	if err := c.Bind().Body(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos de actualización inválidos",
		})
	}

	res, err := roleusecase.UpdateRoleUsecase(uint(id), body, h.repo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al actualizar el rol",
		})
	}
	if res == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Rol no encontrado",
		})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
