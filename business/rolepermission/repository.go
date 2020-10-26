package rolepermission

import core "arkan-jaya/core/rolepermission"

//Repository ingoing port for role permission
type Repository interface {
	//FindRolePermissionByID If data not found will return nil without error
	FindRolePermissionByRoleID(ID string) (*core.RolePermission, error)

	//FindAllByTag If no data match with the given tag, will return empty slice instead of nil
	GetAll() ([]core.RolePermission, error)

	//InsertRolePermission Insert new role permission into storage
	InsertRolePermission(rolePermission core.RolePermission) error

	//UpdateRolePermission if data not found will return core.ErrZeroAffected
	UpdateRolePermission(rolePermission core.RolePermission, currentVersion int) error

	//UpdateRolePermission if data not found will return core.ErrZeroAffected
	DeleteRolePermission(ID string) error
}
