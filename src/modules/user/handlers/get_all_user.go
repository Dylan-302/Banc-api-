package handlers

import (
	"banc-api/src/modules/user/application/mapper"
	"banc-api/src/modules/user/application/usecase"
	"banc-api/src/modules/user/domain/repositories"

	"github.com/gofiber/fiber/v3"
)

type GetAllUserHandler struct {
	repo repositories.UserRepository
}

func NewGetAllUserHandler(repo repositories.UserRepository) *GetAllUserHandler {
	return &GetAllUserHandler{repo: repo}
}

func (h *GetAllUserHandler) Handle(c fiber.Ctx) error {
	users, err := usecase.GetAllUserUsecase(h.repo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al traer la lista de usuarios",
		})
	}

	return c.Status(fiber.StatusOK).JSON(mapper.UsersEntitiesToResponses(users))
}
