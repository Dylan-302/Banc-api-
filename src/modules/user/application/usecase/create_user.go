package usecase

import (
	"time"

	"banc-api/src/modules/user/application/dto/request"
	"banc-api/src/modules/user/application/dto/response"
	"banc-api/src/modules/user/domain/entities"
	"banc-api/src/modules/user/domain/repositories"
)

// CreateUserUsecase crea un nuevo usuario en el dominio.
func CreateUserUsecase(repo repositories.UserRepository, req request.CreateUserRequest) (*response.UserResponse, error) {
	user := entities.User{
		Username:        req.Username,
		Email:           req.Email,
		Password:        req.Password,
		FechadeCreacion: time.Now(),
	}

	createdUser, err := repo.Create(user)
	if err != nil {
		return nil, err
	}

	return &response.UserResponse{
		ID:              createdUser.ID,
		Username:        createdUser.Username,
		Email:           createdUser.Email,
		FechadeCreacion: createdUser.FechadeCreacion.Format("2006-01-02 15:04:05"),
	}, nil
}
