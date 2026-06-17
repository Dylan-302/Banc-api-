package usecase

import (
	"banc-api/src/modules/user/domain/entities"
	"banc-api/src/modules/user/domain/repositories"
)

// GetAllUserUsecase obtiene todos los usuarios desde el repositorio.
func GetAllUserUsecase(repo repositories.UserRepository) ([]entities.User, error) {
	return repo.GetAll()
}