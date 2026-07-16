package usecase

import (
	response "banc-api/src/modules/account/application/dto/response"
	repositories "banc-api/src/modules/account/domain/repositories"
)

func GetAllAccountsUsecase(repo repositories.AccountRepository) ([]*response.AccountResponse, error) {
	accountsEntities, err := repo.FindAll()
	if err != nil {
		return nil, err
	}

	responseList := make([]*response.AccountResponse, 0, len(accountsEntities))
	for _, account := range accountsEntities {
		responseList = append(responseList, &response.AccountResponse{
			IDAccount: account.IDAccount,
			Username:  account.Username,
			Email:     account.Email,
		})
	}

	return responseList, nil
}
