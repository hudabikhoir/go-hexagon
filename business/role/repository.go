package role

import core "arkan-jaya/core/role"

//Repository ingoing port for role
type Repository interface {
	//FindRoleByID If data not found will return nil without error
	FindRoleByID(ID string) (*core.Role, error)

	//FindAllByTag If no data match with the given tag, will return empty slice instead of nil
	GetAll() ([]core.Role, error)

	//InsertRole Insert new role into storage
	InsertRole(role core.Role) error

	//UpdateRole if data not found will return core.ErrZeroAffected
	UpdateRole(role core.Role, currentVersion int) error

	//UpdateRole if data not found will return core.ErrZeroAffected
	DeleteRole(ID string) error
}
