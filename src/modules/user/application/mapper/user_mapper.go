package mapper

import (
	"banc-api/src/modules/user/application/dto/request"
	"banc-api/src/modules/user/application/dto/response"
	"banc-api/src/modules/user/domain/entities"
)

// CreateUserRequestToEntity mapea una solicitud de creación de usuario a una entidad.
func CreateUserRequestToEntity(req request.CreateUserRequest) entities.User {
	return entities.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

// UserEntityToResponse mapea una entidad de usuario a una respuesta.
func UserEntityToResponse(user entities.User) response.UserResponse {
	return response.UserResponse{
		ID:              int(user.ID),
		Username:        user.Username,
		Email:           user.Email,
		FechadeCreacion: user.FechadeCreacion,
	}
}
