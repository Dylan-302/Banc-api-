package usecase

import (
	"banc-api/src/modules/user/domain/entities"
	"banc-api/src/modules/user/domain/repositories"
)

// UpdateUserUsecase actualiza un usuario existente en la base de datos.
func UpdateUserUsecase(user entities.User, repo repositories.UserRepository) (*entities.User, error) {
	return repo.Update(user)
}
