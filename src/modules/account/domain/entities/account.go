package entities

import "time"

// Account es la entidad principal de cuentas en el sistema
type Account struct {
	IDAccount int       `json:"id_account"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
