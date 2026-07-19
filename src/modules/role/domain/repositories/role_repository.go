package repositories

import entities "banc-api/src/modules/role/domain/entities"

// RoleRepository define los métodos que necesita la capa de dominio para persistencia.
type RoleRepository interface {
	FindAll() ([]entities.Role, error)
	FindByID(id uint) (*entities.Role, error)
}
