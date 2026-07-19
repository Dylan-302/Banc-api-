package usecase

import (
	response "banc-api/src/modules/role/application/dto/response"
	repositories "banc-api/src/modules/role/domain/repositories"
)

func GetAllRolesUsecase(repo repositories.RoleRepository) ([]*response.RoleResponse, error) {
	rolesEntities, err := repo.FindAll()
	if err != nil {
		return nil, err
	}

	responseList := make([]*response.RoleResponse, 0, len(rolesEntities))
	for _, role := range rolesEntities {
		responseList = append(responseList, &response.RoleResponse{
			ID:   role.ID,
			Name: role.Name,
		})

	}

	return responseList, nil
}
