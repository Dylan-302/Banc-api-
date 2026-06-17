package dao

import (
	"banc-api/src/infrastructure/db/models"
	"banc-api/src/modules/user/domain/entities"

	"gorm.io/gorm"
)

// UserDao implementa UserRepository usando GORM.
type UserDao struct {
	db *gorm.DB
}

// NewUserDao crea un nuevo DAO de usuarios.
func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

// GetAll devuelve todos los usuarios almacenados.
func (u *UserDao) GetAll() ([]entities.User, error) {
	var users []models.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	result := make([]entities.User, len(users))
	for i, item := range users {
		result[i] = entities.User{
			ID:              item.ID,
			Username:        item.Username,
			Email:           item.Email,
			Password:        item.Password,
			FechadeCreacion: item.FechadeCreacion,
		}
	}

	return result, nil
}

// Create inserta un nuevo usuario en la base de datos.
func (u *UserDao) Create(user entities.User) (*entities.User, error) {
	model := models.User{
		Username:        user.Username,
		Email:           user.Email,
		Password:        user.Password,
		FechadeCreacion: user.FechadeCreacion,
	}

	if err := u.db.Create(&model).Error; err != nil {
		return nil, err
	}

	user.ID = model.ID
	return &user, nil
}

// FindByID busca un usuario por su ID.
func (u *UserDao) FindByID(id uint) (*entities.User, error) {
	var model models.User
	if err := u.db.First(&model, id).Error; err != nil {
		return nil, err
	}

	return &entities.User{
		ID:              model.ID,
		Username:        model.Username,
		Email:           model.Email,
		Password:        model.Password,
		FechadeCreacion: model.FechadeCreacion,
	}, nil
}

// Update actualiza un usuario existente.
func (u *UserDao) Update(user entities.User) (*entities.User, error) {
	model := models.User{
		ID:              user.ID,
		Username:        user.Username,
		Email:           user.Email,
		Password:        user.Password,
		FechadeCreacion: user.FechadeCreacion,
	}

	if err := u.db.Save(&model).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Delete elimina un usuario por su ID.
func (u *UserDao) Delete(id uint) error {
	return u.db.Delete(&models.User{}, id).Error
}
