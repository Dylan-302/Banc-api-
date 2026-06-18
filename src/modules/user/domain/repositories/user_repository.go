package repositories

import "banc-api/src/modules/user/domain/entities"

type UserRepository interface {
	GetAll() ([]entities.User, error)
	GetByID(id uint) (*entities.User, error)
	Create(user *entities.User) error
	Update(user *entities.User) error
	Delete(id uint) error
}
