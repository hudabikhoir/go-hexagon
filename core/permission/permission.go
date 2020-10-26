package permission

import "time"

//Permission product item that available to rent or sell
type Permission struct {
	ID         string
	Resource   string
	Permission string
	CreatedAt  time.Time
	CreatedBy  string
	ModifiedAt time.Time
	ModifiedBy string
	Version    int
}

//NewPermission create new item
func NewPermission(
	id string,
	resource string,
	permission string,
	creator string,
	createdAt time.Time) Permission {

	return Permission{
		ID:         id,
		Resource:   resource,
		Permission: permission,
		CreatedAt:  createdAt,
		CreatedBy:  creator,
		ModifiedAt: createdAt,
		ModifiedBy: creator,
		Version:    1,
	}
}

//ModifyPermission update existing item data
func (oldPermission *Permission) ModifyPermission(newResource string, newPermission string, updater string, modifiedAt time.Time) Permission {
	return Permission{
		ID:         oldPermission.ID,
		Resource:   newResource,
		Permission: newPermission,
		CreatedAt:  oldPermission.CreatedAt,
		CreatedBy:  oldPermission.CreatedBy,
		ModifiedAt: modifiedAt,
		ModifiedBy: updater,
		Version:    oldPermission.Version + 1,
	}
}
