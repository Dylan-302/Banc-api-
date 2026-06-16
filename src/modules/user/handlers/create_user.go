package handlers

import (
	"github.com/gofiber/fiber/v2"

	"banc-api/src/modules/user/application/usecase"
	"banc-api/src/modules/user/domain/repositories"
)

type UserHandler struct {
	Repo repositories.UserRepository
}

func NewUserHandler(repo repositories.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	createdUser, err := usecase.CreateUser(req.Username, req.Email, req.Password, h.Repo)
	if err != nil {
		return err
	}
	// Construir respuesta mínima (sin exponer password)
	resp := struct {
		ID              uint   `json:"id"`
		Username        string `json:"username"`
		Email           string `json:"email"`
		FechadeCreacion string `json:"fecha_creacion"`
	}{
		ID:              createdUser.ID,
		Username:        createdUser.Username,
		Email:           createdUser.Email,
		FechadeCreacion: createdUser.FechadeCreacion.Format("2006-01-02T15:04:05Z07:00"),
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}
