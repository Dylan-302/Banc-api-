package usecase

import (
	"banc-api/src/modules/user/domain/entities"
	"banc-api/src/modules/user/domain/repositories"
)

type FindByIDUseCase struct {
	repo repositories.UserRepository
}

// NewFindByIDUseCase crea un caso de uso para buscar usuario por ID
func NewFindByIDUseCase(repo repositories.UserRepository) *FindByIDUseCase {
	return &FindByIDUseCase{repo: repo}
}

// Execute busca y retorna el usuario correspondiente al ID
func (uc *FindByIDUseCase) Execute(id uint) (*entities.User, error) {
	return uc.repo.GetByID(id)
}
