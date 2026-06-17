package usecase

import (
	"banc-api/src/modules/user/application/dto/request"
	"banc-api/src/modules/user/application/dto/response"
	"banc-api/src/modules/user/domain/repositories"
)

// UpdateUserUsecase actualiza los datos de un usuario.
func UpdateUserUsecase(id uint, req request.UpdateUserRequest, repo repositories.UserRepository) (*response.UserResponse, error) {
	user, err := repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		user.Password = req.Password
	}

	updatedUser, err := repo.Update(*user)
	if err != nil {
		return nil, err
	}

	return &response.UserResponse{
		ID:              updatedUser.ID,
		Username:        updatedUser.Username,
		Email:           updatedUser.Email,
		FechadeCreacion: updatedUser.FechadeCreacion.Format("2006-01-02 15:04:05"),
	}, nil
}
