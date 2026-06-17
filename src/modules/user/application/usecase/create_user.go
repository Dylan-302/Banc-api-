package usecase

import (
	"banc-api/src/modules/user/domain/entities"
	"banc-api/src/modules/user/domain/repositories"
)

// CreateUserUsecase crea un nuevo usuario en la base de datos.
func CreateUserUsecase(user entities.User, repo repositories.UserRepository) (*entities.User, error) {
	return repo.Create(user)
}
