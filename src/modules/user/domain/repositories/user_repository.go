package repositories

import "banc-api/src/modules/user/domain/entities"

// UserRepository define los métodos que necesita la capa de dominio para persistencia.
type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	GetAll() ([]entities.User, error)
}
