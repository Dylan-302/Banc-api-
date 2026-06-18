package usecase

import (
	"banc-api/src/modules/user/application/dto/response"
	"banc-api/src/modules/user/domain/repositories"
	"errors"
)

func UserFindByIdUsecase(id uint, repo repositories.UserRepository) (*response.UserResponse, error) {

	user, err := repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("usuario no encontrado")
	}

	res := &response.UserResponse{
		ID:              int(user.ID),
		Username:        user.Username,
		Email:           user.Email,
		FechadeCreacion: user.FechadeCreacion,
	}

	return res, nil
}
