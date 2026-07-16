package account

import (
	"banc-api/src/modules/account/application/dto/request"
	response "banc-api/src/modules/account/application/dto/response"
	"banc-api/src/modules/account/domain/entities"
)

// CreateAccountRequestToEntity mapea una solicitud de creación de cuenta a una entidad.
func CreateAccountRequestToEntity(req request.CreateAccountRequest) entities.Account {
	return entities.Account{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

// AccountEntityToResponse mapea una entidad de cuenta a una respuesta.
func AccountEntityToResponse(account entities.Account) response.CreateAccountResponse {
	return response.CreateAccountResponse{
		IDAccount: account.IDAccount,
		Username:  account.Username,
		Email:     account.Email,
	}
}

// AccountsEntitiesToResponses mapea múltiples entidades de cuenta a respuestas.
func AccountsEntitiesToResponses(accounts []entities.Account) []response.CreateAccountResponse {
	result := make([]response.CreateAccountResponse, len(accounts))
	for i, account := range accounts {
		result[i] = AccountEntityToResponse(account)
	}
	return result
}
