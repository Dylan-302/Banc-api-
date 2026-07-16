package usecase

import (
	"banc-api/src/modules/account/domain/entities"
	"banc-api/src/modules/account/domain/repositories"
)

// CreateAccountUsecase crea un nuevo account en la base de datos.
func CreateAccountUsecase(account entities.Account, repo repositories.AccountRepository) (*entities.Account, error) {
	return repo.Save(account)
}
