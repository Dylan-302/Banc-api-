package handlers

import (
	"banc-api/src/modules/user/application/usecase"
	"banc-api/src/modules/user/domain/repositories"

	"github.com/gofiber/fiber/v3"
)

type GetAllUsersHandler struct {
	repo repositories.UserRepository
}

func NewGetAllUsersHandler(repo repositories.UserRepository) *GetAllUsersHandler {
	return &GetAllUsersHandler{repo: repo}
}

func (h *GetAllUsersHandler) Handle(c fiber.Ctx) error {
	responseList, err := usecase.GetAllUsersUsecase(h.repo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al recuperar la lista de usuarios",
		})
	}

	return c.Status(fiber.StatusOK).JSON(responseList)
}
