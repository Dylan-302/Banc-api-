package usecase

import (
	"banc-api/src/modules/user/application/dto/request"
	"banc-api/src/modules/user/domain/entities"
	"banc-api/src/modules/user/domain/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UpdateUserUseCase struct {
	repo repositories.UserRepository
}

// NewUpdateUserUseCase crea un caso de uso para actualizar usuario
func NewUpdateUserUseCase(repo repositories.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{repo: repo}
}

// Execute actualiza un usuario existente de forma parcial
func (uc *UpdateUserUseCase) Execute(id uint, req request.UpdateUserRequest) (*entities.User, error) {
	user, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	if err := uc.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}
