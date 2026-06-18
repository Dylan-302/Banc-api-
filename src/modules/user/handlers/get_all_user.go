package handlers

import (
	"banc-api/src/common/response"
	"banc-api/src/modules/user/application/mapper"
	"banc-api/src/modules/user/application/usecase"

	"github.com/gin-gonic/gin"
)

type GetAllUserHandler struct {
	useCase *usecase.GetAllUserUseCase
}

// NewGetAllUserHandler crea una nueva instancia del handler de listado de usuarios
func NewGetAllUserHandler(useCase *usecase.GetAllUserUseCase) *GetAllUserHandler {
	return &GetAllUserHandler{useCase: useCase}
}

// Handle procesa la petición de listado de todos los usuarios
func (h *GetAllUserHandler) Handle(c *gin.Context) {
	users, err := h.useCase.Execute()
	if err != nil {
		response.InternalError(c, "Error al obtener usuarios: "+err.Error())
		return
	}

	response.Success(c, mapper.ToResponseList(users), "Usuarios obtenidos con éxito")
}
