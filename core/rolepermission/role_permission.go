package rolepermission

import "time"

//RolePermission product item that available to rent or sell
type RolePermission struct {
	RoleID       string
	PermissionID string
	CreatedAt    time.Time
	CreatedBy    string
	ModifiedAt   time.Time
	ModifiedBy   string
	Version      int
}

//NewRolePermission create new item
func NewRolePermission(
	roleID string,
	permissionID string,
	creator string,
	createdAt time.Time) RolePermission {

	return RolePermission{
		RoleID:       roleID,
		PermissionID: permissionID,
		CreatedAt:    createdAt,
		CreatedBy:    creator,
		ModifiedAt:   createdAt,
		ModifiedBy:   creator,
		Version:      1,
	}
}

//ModifyRolePermission update existing item data
func (oldRolePermission *RolePermission) ModifyRolePermission(newRoleID string, newPermissionID string, updater string, modifiedAt time.Time) RolePermission {
	return RolePermission{
		RoleID:       newRoleID,
		PermissionID: newPermissionID,
		CreatedAt:    oldRolePermission.CreatedAt,
		CreatedBy:    oldRolePermission.CreatedBy,
		ModifiedAt:   modifiedAt,
		ModifiedBy:   updater,
		Version:      oldRolePermission.Version + 1,
	}
}
