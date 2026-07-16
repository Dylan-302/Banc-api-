package usecase

import (
	response "banc-api/src/modules/account/application/dto/response"
	"banc-api/src/modules/account/domain/repositories"
)

type UpdateAccountDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UpdateAccountUsecase(id uint, data UpdateAccountDTO, repo repositories.AccountRepository) (*response.AccountResponse, error) {
	account, err := repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, nil
	}

	if data.Username != "" {
		account.Username = data.Username
	}
	if data.Email != "" {
		account.Email = data.Email
	}
	if data.Password != "" {
		account.Password = data.Password
	}

	updated, err := repo.Update(*account)
	if err != nil {
		return nil, err
	}
	if updated == nil {
		return nil, nil
	}

	res := &response.AccountResponse{
		IDAccount: updated.IDAccount,
		Username:  updated.Username,
		Email:     updated.Email,
	}

	return res, nil
}
