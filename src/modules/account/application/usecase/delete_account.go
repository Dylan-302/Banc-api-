package usecase

import "banc-api/src/modules/account/domain/repositories"

// DeleteAccountUsecase elimina un account de la base de datos.
func DeleteAccountUsecase(id uint, repo repositories.AccountRepository) error {
	return repo.Delete(id)
}
