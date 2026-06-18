package usecase

import (
	"time"

	"banc-api/src/modules/user/application/dto/request"
	"banc-api/src/modules/user/domain/entities"
	"banc-api/src/modules/user/domain/repositories"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserUseCase struct {
	repo repositories.UserRepository
}

// NewCreateUserUseCase crea un caso de uso para crear usuario
func NewCreateUserUseCase(repo repositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo}
}

// Execute ejecuta la creación del usuario con contraseña encriptada
func (uc *CreateUserUseCase) Execute(req request.CreateUserRequest) (*entities.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		Username:        req.Username,
		Email:           req.Email,
		Password:        string(hashedPassword),
		FechadeCreacion: time.Now(),
	}

	if err := uc.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
