package handlers

import (
	"banc-api/src/modules/account/application/usecase"
	"banc-api/src/modules/account/domain/repositories"

	"github.com/gofiber/fiber/v3"
)

type GetAllAccountsHandler struct {
	repo repositories.AccountRepository
}

func NewGetAllAccountsHandler(repo repositories.AccountRepository) *GetAllAccountsHandler {
	return &GetAllAccountsHandler{repo: repo}
}

func (h *GetAllAccountsHandler) Handle(c fiber.Ctx) error {
	responseList, err := usecase.GetAllAccountsUsecase(h.repo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al recuperar la lista de cuentas",
		})
	}

	return c.Status(fiber.StatusOK).JSON(responseList)
}
