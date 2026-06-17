package usecase

import "banc-api/src/modules/user/domain/repositories"

// DeleteUserUsecase elimina un usuario de la base de datos.
func DeleteUserUsecase(id uint, repo repositories.UserRepository) error {
	return repo.Delete(id)
}
