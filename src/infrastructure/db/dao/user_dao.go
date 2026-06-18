package dao

import (
	"banc-api/src/infrastructure/db/models"
	"banc-api/src/modules/user/application/mapper"
	"banc-api/src/modules/user/domain/entities"
	"banc-api/src/modules/user/domain/repositories"

	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

// NewUserDAO crea una nueva instancia de UserDAO
func NewUserDAO(db *gorm.DB) repositories.UserRepository {
	return &UserDAO{db: db}
}

// GetAll obtiene todos los usuarios
func (d *UserDAO) GetAll() ([]entities.User, error) {
	var userModels []models.User
	if err := d.db.Find(&userModels).Error; err != nil {
		return nil, err
	}

	var userEntities []entities.User
	for _, m := range userModels {
		entity := mapper.ToEntity(&m)
		if entity != nil {
			userEntities = append(userEntities, *entity)
		}
	}
	return userEntities, nil
}

// GetByID obtiene un usuario por su ID
func (d *UserDAO) GetByID(id uint) (*entities.User, error) {
	var userModel models.User
	if err := d.db.First(&userModel, id).Error; err != nil {
		return nil, err
	}
	return mapper.ToEntity(&userModel), nil
}

// Create crea un nuevo usuario
func (d *UserDAO) Create(user *entities.User) error {
	userModel := mapper.ToModel(user)
	if err := d.db.Create(userModel).Error; err != nil {
		return err
	}
	user.ID = userModel.ID
	user.FechadeCreacion = userModel.FechadeCreacion
	return nil
}

// Update actualiza un usuario existente
func (d *UserDAO) Update(user *entities.User) error {
	userModel := mapper.ToModel(user)
	if err := d.db.Save(userModel).Error; err != nil {
		return err
	}
	return nil
}

// Delete elimina un usuario por su ID
func (d *UserDAO) Delete(id uint) error {
	if err := d.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
