package usecase

import (
	"banc-api/src/modules/user/application/dto/response"
	"banc-api/src/modules/user/domain/repositories"
)

// FindByIDUsecase busca un usuario por su ID.
func FindByIDUsecase(id uint, repo repositories.UserRepository) (*response.UserResponse, error) {
	user, err := repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &response.UserResponse{
		ID:              user.ID,
		Username:        user.Username,
		Email:           user.Email,
		FechadeCreacion: user.FechadeCreacion.Format("2006-01-02 15:04:05"),
	}, nil
}
