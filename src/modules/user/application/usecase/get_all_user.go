package usecase

import (
	"banc-api/src/modules/user/application/dto/response"
	"banc-api/src/modules/user/domain/repositories"
)

func GetAllUsersUsecase(repo repositories.UserRepository) ([]*response.UserResponse, error) {
	usersEntities, err := repo.FindAll()
	if err != nil {
		return nil, err
	}

	var responseList []*response.UserResponse

	for i := range usersEntities {
		user := usersEntities[i]
		dto := &response.UserResponse{
			ID:       int(user.ID),
			Username: user.Username,
			Email:    user.Email,
		}
		responseList = append(responseList, dto)
	}

	return responseList, nil
}
