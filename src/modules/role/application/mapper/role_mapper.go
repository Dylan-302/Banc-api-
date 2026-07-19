package role

import "banc-api/src/modules/role/domain/entities"

// RoleFindByID retrieves a role by its ID from the database or any other data source.
func RoleFindByID(id uint) (*entities.Role, error) {
	return &entities.Role{
		ID:          int(id),
		Name:        "",
		Description: "",
	}, nil
}
