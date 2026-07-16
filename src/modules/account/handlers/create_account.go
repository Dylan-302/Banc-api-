package handlers

import (
	"banc-api/src/modules/account/application/dto/request"
	mapper "banc-api/src/modules/account/application/mapper"
	"banc-api/src/modules/account/application/usecase"
	"banc-api/src/modules/account/domain/repositories"

	"github.com/gofiber/fiber/v3"
)

type CreateAccountHandlers struct {
	repo repositories.AccountRepository
}

func NewCreateAccountHandlers(repo repositories.AccountRepository) *CreateAccountHandlers {
	return &CreateAccountHandlers{repo}

}

func (h *CreateAccountHandlers) Handler(c fiber.Ctx) error {

	var req request.CreateAccountRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	account := mapper.CreateAccountRequestToEntity(req)

	createAccount, err := usecase.CreateAccountUsecase(account, h.repo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create account",
		})
	}

	response := mapper.AccountEntityToResponse(*createAccount)
	return c.Status(fiber.StatusCreated).JSON(response)
}
