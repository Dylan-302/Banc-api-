package handlers

import (
	"net/http"

	"banc-api/src/common/response"
	"banc-api/src/modules/user/application/dto/request"
	"banc-api/src/modules/user/application/mapper"
	"banc-api/src/modules/user/application/usecase"

	"github.com/gin-gonic/gin"
)

type CreateUserHandler struct {
	useCase *usecase.CreateUserUseCase
}

// NewCreateUserHandler crea una nueva instancia del handler de creación de usuario
func NewCreateUserHandler(useCase *usecase.CreateUserUseCase) *CreateUserHandler {
	return &CreateUserHandler{useCase: useCase}
}

// Handle procesa la petición de creación de usuario
func (h *CreateUserHandler) Handle(c *gin.Context) {
	var req request.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusUnprocessableEntity, "Datos de solicitud inválidos", err.Error())
		return
	}

	user, err := h.useCase.Execute(req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Error al crear el usuario", err.Error())
		return
	}

	response.Created(c, mapper.ToResponse(user), "Usuario creado con éxito")
}
