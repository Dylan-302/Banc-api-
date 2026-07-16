package handlers

import (
	"strconv"

	"banc-api/src/modules/account/application/usecase"
	"banc-api/src/modules/account/domain/repositories"

	"github.com/gofiber/fiber/v3"
)

type FindByIdAccountHandler struct {
	repo repositories.AccountRepository
}

func NewFindByIdAccountHandler(repo repositories.AccountRepository) *FindByIdAccountHandler {
	return &FindByIdAccountHandler{repo: repo}
}

func (h *FindByIdAccountHandler) Handle(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El ID proporcionado no es válido",
		})
	}

	response, err := usecase.AccountFindByIdUsecase(uint(id), h.repo)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

type SearchAccountHandler struct {
	repo repositories.AccountRepository
}

func NewSearchAccountHandler(repo repositories.AccountRepository) *SearchAccountHandler {
	return &SearchAccountHandler{repo: repo}
}

func (h *SearchAccountHandler) Handle(c fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El ID proporcionado no es válido",
		})
	}

	response, err := usecase.AccountFindByIdUsecase(uint(id), h.repo)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
