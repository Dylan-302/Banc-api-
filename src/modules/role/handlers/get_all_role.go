package handlers

import (
	"banc-api/src/modules/role/application/usecase"
	"banc-api/src/modules/role/domain/repositories"

	"github.com/gofiber/fiber/v3"
)

type GetAllRolesHandler struct {
	repo repositories.RoleRepository
}

func NewGetAllRolesHandler(repo repositories.RoleRepository) *GetAllRolesHandler {
	return &GetAllRolesHandler{repo: repo}
}

func (h *GetAllRolesHandler) Handle(c fiber.Ctx) error {
	responseList, err := usecase.GetAllRolesUsecase(h.repo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al recuperar la lista de roles",
		})
	}

	return c.Status(fiber.StatusOK).JSON(responseList)
}
