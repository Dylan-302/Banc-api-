package response

import "time"

type NewUserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponseDelete struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	FechadeCreacion string `json:"fecha_creacion"`
}

type UserResponse struct {
	ID              int       `json:"id"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	FechadeCreacion time.Time `json:"fecha_creacion"`
}
