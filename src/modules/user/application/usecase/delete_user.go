package usecase

import (
	"banc-api/src/modules/user/domain/repositories"
)

type DeleteUserUseCase struct {
	repo repositories.UserRepository
}

// NewDeleteUserUseCase crea un caso de uso para eliminar usuario
func NewDeleteUserUseCase(repo repositories.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{repo: repo}
}

// Execute elimina el usuario después de verificar su existencia
func (uc *DeleteUserUseCase) Execute(id uint) error {
	_, err := uc.repo.GetByID(id)
	if err != nil {
		return err
	}
	return uc.repo.Delete(id)
}
