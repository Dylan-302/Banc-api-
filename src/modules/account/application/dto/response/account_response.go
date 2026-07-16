package response

type AccountResponse struct {
	IDAccount int    `json:"id_account"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type CreateAccountResponse struct {
	IDAccount int    `json:"id_account"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type UpdateAccountResponse struct {
	IDAccount int    `json:"id_account"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type DeleteAccountResponse struct {
	IDAccount int    `json:"id_account"`
	Email     string `json:"email"`
}
