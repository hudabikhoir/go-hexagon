package role

import "time"

//Role product item that available to rent or sell
type Role struct {
	ID           string
	Name         string
	PermissionID []string
	CreatedAt    time.Time
	CreatedBy    string
	ModifiedAt   time.Time
	ModifiedBy   string
	Version      int
}

//NewRole create new item
func NewRole(
	id string,
	name string,
	permissionID []string,
	creator string,
	createdAt time.Time) Role {

	return Role{
		ID:           id,
		Name:         name,
		PermissionID: permissionID,
		CreatedAt:    createdAt,
		CreatedBy:    creator,
		ModifiedAt:   createdAt,
		ModifiedBy:   creator,
		Version:      1,
	}
}

//ModifyRole update existing item data
func (oldRole *Role) ModifyRole(newName string, newPermissionID []string, updater string, modifiedAt time.Time) Role {
	return Role{
		ID:           oldRole.ID,
		Name:         newName,
		PermissionID: newPermissionID,
		CreatedAt:    oldRole.CreatedAt,
		CreatedBy:    oldRole.CreatedBy,
		ModifiedAt:   modifiedAt,
		ModifiedBy:   updater,
		Version:      oldRole.Version + 1,
	}
}
