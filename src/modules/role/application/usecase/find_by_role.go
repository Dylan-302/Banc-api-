package usecase

import (
	"errors"

	response "banc-api/src/modules/role/application/dto/response"
	"banc-api/src/modules/role/domain/repositories"
)

func FindByRoleUsecase(id uint, repo repositories.RoleRepository) (*response.RoleResponse, error) {
	role, err := repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if role == nil {
		return nil, errors.New("rol no encontrado")
	}

	res := &response.RoleResponse{
		ID:   role.ID,
		Name: role.Name,
	}

	return res, nil
}
