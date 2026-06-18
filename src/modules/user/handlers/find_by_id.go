package handlers

import (
	"strconv"

	"banc-api/src/modules/user/application/usecase"
	"banc-api/src/modules/user/domain/repositories"

	"github.com/gofiber/fiber/v3"
)

type FindByIdUserHandler struct {
	repo repositories.UserRepository
}

func NewFindByIdUserHandler(repo repositories.UserRepository) *FindByIdUserHandler {
	return &FindByIdUserHandler{repo: repo}
}

func (h *FindByIdUserHandler) Handle(c fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El ID proporcionado no es válido",
		})
	}

	response, err := usecase.UserFindByIdUsecase(uint(id), h.repo)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
