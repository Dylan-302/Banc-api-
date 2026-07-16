package repositories

import entities "banc-api/src/modules/account/domain/entities"

// AccountRepository define los métodos que necesita la capa de dominio para persistencia.
type AccountRepository interface {
	Save(account entities.Account) (*entities.Account, error)
	FindAll() ([]entities.Account, error)
	FindByID(id uint) (*entities.Account, error)
	Update(account entities.Account) (*entities.Account, error)
	Delete(id uint) error
}
