package usecase

import (
	"banc-api/src/modules/user/application/dto/response"
	"banc-api/src/modules/user/domain/repositories"
	"errors"
)

func UserFindByIdUsecase(id uint, repo repositories.UserRepository) (*response.UserResponse, error) {
	user, err := FindByIDUsecase(id, repo)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func FindByIDUsecase(id uint, repo repositories.UserRepository) (*response.UserResponse, error) {
	user, err := repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("usuario no encontrado")
	}

	respons := &response.UserResponse{
		ID:              int(user.ID),
		Username:        user.Username,
		Email:           user.Email,
		FechadeCreacion: user.FechadeCreacion,
	}

	return respons, nil
}
