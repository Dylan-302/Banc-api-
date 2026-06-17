package handlers

import (
	"strconv" // Sirve para convertir texto a número

	"banc-api/src/modules/user/application/usecase"
	"banc-api/src/modules/user/domain/repositories"

	"github.com/gofiber/fiber/v3"
)

type SearchUserHandler struct {
	repo repositories.UserRepository
}

func NewSearchUserHandler(repo repositories.UserRepository) *SearchUserHandler {
	return &SearchUserHandler{repo: repo}
}

func (h *SearchUserHandler) Handle(c fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El ID proporcionado no es válido",
		})
	}

	user, err := usecase.FindByIDUsecase(uint(id), h.repo)
	if err != nil {
		// Si el caso de uso devuelve un error (usuario no encontrado)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
