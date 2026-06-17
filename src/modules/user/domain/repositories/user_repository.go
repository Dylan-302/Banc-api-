package repositories

import "banc-api/src/modules/user/domain/entities"

type UserRepository interface {
	GetAll() ([]entities.User, error)
	Create(user entities.User) (*entities.User, error)
	FindByID(id uint) (*entities.User, error)
	Update(user entities.User) (*entities.User, error)
	Delete(id uint) error
}
