package usecase

import (
	"banc-api/src/modules/user/domain/entities"
	"banc-api/src/modules/user/domain/repositories"
)

func FindByIDUsecase(id uint, repo repositories.UserRepository) (*entities.User, error) {
	user, err := repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return user, nil
}
