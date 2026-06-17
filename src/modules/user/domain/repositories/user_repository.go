package repositories

import "banc-api/src/modules/user/domain/entities"

// UserRepository define los métodos que necesita la capa de dominio para persistencia.
type UserRepository interface {
	Create(user entities.User) (*entities.User, error)
	FindByID(id uint) (*entities.User, error)
	Update(user entities.User) (*entities.User, error)
	Delete(id uint) error
}
