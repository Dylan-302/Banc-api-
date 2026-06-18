package usecase

import (
	"banc-api/src/modules/user/domain/entities"
	"banc-api/src/modules/user/domain/repositories"
)

type GetAllUserUseCase struct {
	repo repositories.UserRepository
}

// NewGetAllUserUseCase crea un caso de uso para obtener todos los usuarios
func NewGetAllUserUseCase(repo repositories.UserRepository) *GetAllUserUseCase {
	return &GetAllUserUseCase{repo: repo}
}

// Execute obtiene todos los usuarios del repositorio
func (uc *GetAllUserUseCase) Execute() ([]entities.User, error) {
	return uc.repo.GetAll()
}
