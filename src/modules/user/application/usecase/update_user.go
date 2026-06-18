package usecase

import (
	"banc-api/src/modules/user/application/dto/response"
	"banc-api/src/modules/user/domain/repositories"
)

type UpdateUserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func UpdateUserUsecase(id uint, data UpdateUserDTO, repo repositories.UserRepository) (*response.UserResponse, error) {
	user, err := repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	if data.Username != "" {
		user.Username = data.Username
	}
	if data.Email != "" {
		user.Email = data.Email
	}

	res := &response.UserResponse{
		ID:       int(user.ID),
		Username: user.Username,
		Email:    user.Email,
	}

	return res, nil
}
