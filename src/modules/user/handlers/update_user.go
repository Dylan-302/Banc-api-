package handlers

import (
	"net/http"
	"strconv"

	"banc-api/src/common/response"
	"banc-api/src/modules/user/application/dto/request"
	"banc-api/src/modules/user/application/mapper"
	"banc-api/src/modules/user/application/usecase"

	"github.com/gin-gonic/gin"
)

type UpdateUserHandler struct {
	useCase *usecase.UpdateUserUseCase
}

// NewUpdateUserHandler crea una nueva instancia del handler de actualización de usuario
func NewUpdateUserHandler(useCase *usecase.UpdateUserUseCase) *UpdateUserHandler {
	return &UpdateUserHandler{useCase: useCase}
}

// Handle procesa la petición de actualización de usuario
func (h *UpdateUserHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "ID de usuario inválido")
		return
	}

	var req request.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusUnprocessableEntity, "Datos de solicitud inválidos", err.Error())
		return
	}

	user, err := h.useCase.Execute(uint(id), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Error al actualizar el usuario", err.Error())
		return
	}

	response.Success(c, mapper.ToResponse(user), "Usuario actualizado con éxito")
}
