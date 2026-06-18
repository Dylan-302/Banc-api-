package handlers

import (
	"net/http"
	"strconv"

	"banc-api/src/common/response"
	"banc-api/src/modules/user/application/mapper"
	"banc-api/src/modules/user/application/usecase"

	"github.com/gin-gonic/gin"
)

type FindByIDHandler struct {
	useCase *usecase.FindByIDUseCase
}

// NewFindByIDHandler crea una nueva instancia del handler para buscar usuario por ID
func NewFindByIDHandler(useCase *usecase.FindByIDUseCase) *FindByIDHandler {
	return &FindByIDHandler{useCase: useCase}
}

// Handle procesa la petición de búsqueda de usuario por ID
func (h *FindByIDHandler) Handle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "ID de usuario inválido")
		return
	}

	user, err := h.useCase.Execute(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "Usuario no encontrado", err.Error())
		return
	}

	response.Success(c, mapper.ToResponse(user), "Usuario obtenido con éxito")
}
