package account

import handlers "banc-api/src/modules/account/handlers"

func ConfigureModuleRoutes(
	handlerCreateAccount *handlers.CreateAccountHandlers,
	handlerGetAccountByID *handlers.FindByIdAccountHandler,
	handlerUpdateAccount *handlers.UpdateAccountHandler,
	handlerDeleteAccount *handlers.DeleteAccountHandler,
	handlerGetAllAccounts *handlers.GetAllAccountsHandler,
) {
	_ = handlerCreateAccount
	_ = handlerGetAccountByID
	_ = handlerUpdateAccount
	_ = handlerDeleteAccount
	_ = handlerGetAllAccounts
}
