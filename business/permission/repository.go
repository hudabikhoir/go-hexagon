package permission

import core "arkan-jaya/core/permission"

//Repository ingoing port for permission
type Repository interface {
	//FindPermissionByID If data not found will return nil without error
	FindPermissionByID(ID string) (*core.Permission, error)

	//FindAllByTag If no data match with the given tag, will return empty slice instead of nil
	GetAll() ([]core.Permission, error)

	//InsertPermission Insert new permission into storage
	InsertPermission(permission core.Permission) error

	//UpdatePermission if data not found will return core.ErrZeroAffected
	UpdatePermission(permission core.Permission, currentVersion int) error

	//UpdatePermission if data not found will return core.ErrZeroAffected
	DeletePermission(ID string) error
}
