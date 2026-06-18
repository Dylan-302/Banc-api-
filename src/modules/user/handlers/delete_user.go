package handlers

import (
	"net/http"
	"strconv"

	"banc-api/src/common/response"
	"banc-api/src/modules/user/application/usecase"

	"github.com/gin-gonic/gin"
)

type DeleteUserHandler struct {
	useCase *usecase.DeleteUserUseCase
}

// NewDeleteUserHandler crea una nueva instancia del handler de eliminación de usuario
func NewDeleteUserHandler(useCase *usecase.DeleteUserUseCase) *DeleteUserHandler {
	return &DeleteUserHandler{useCase: useCase}
}

// Handle procesa la petición de eliminación de usuario
func (h *DeleteUserHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "ID de usuario inválido")
		return
	}

	err = h.useCase.Execute(uint(id))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Error al eliminar el usuario", err.Error())
		return
	}

	response.Success(c, nil, "Usuario eliminado con éxito")
}
