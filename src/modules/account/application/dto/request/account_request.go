package request

type CreateAccountRequest struct {
	IDAccount int    `json:"id_account" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

type UpdateAccountRequest struct {
	IDAccount int    `json:"id_account" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

type DeleteAccountRequest struct {
	IDAccount int    `json:"id_account" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}
