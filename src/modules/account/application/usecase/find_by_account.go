package usecase

import (
	"errors"

	response "banc-api/src/modules/account/application/dto/response"
	"banc-api/src/modules/account/domain/repositories"
)

func AccountFindByIdUsecase(id uint, repo repositories.AccountRepository) (*response.AccountResponse, error) {
	account, err := repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, errors.New("cuenta no encontrada")
	}

	res := &response.AccountResponse{
		IDAccount: account.IDAccount,
		Username:  account.Username,
		Email:     account.Email,
	}

	return res, nil
}
