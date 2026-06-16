package usecase

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	"banc-api/src/modules/user/domain/entities"
	"banc-api/src/modules/user/domain/repositories"
)

// CreateUser ejecuta la lógica de negocio para crear un usuario y devuelve la entidad creada.
func CreateUser(username, email, password string, repo repositories.UserRepository) (*entities.User, error) {
	if username == "" || email == "" || password == "" {
		return nil, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	userEntity := &entities.User{
		Username:        username,
		Email:           email,
		Password:        string(hashedPassword),
		FechadeCreacion: time.Now(),
	}

	createdUser, err := repo.Create(userEntity)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
