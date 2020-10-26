package response

import (
	"arkan-jaya/core/permission"
	"time"
)

//GetPermissionByIDResponse Get permission by ID response payload
type GetPermissionByIDResponse struct {
	ID         string    `json:"id"`
	Resource   string    `json:"resource"`
	Permission string    `json:"permission"`
	ModifiedAt time.Time `json:"modifiedAt"`
	Version    int       `json:"version"`
}

//NewGetPermissionByIDResponse construct GetPermissionByIDResponse
func NewGetPermissionByIDResponse(permission permission.Permission) *GetPermissionByIDResponse {
	var permissionResponse GetPermissionByIDResponse
	permissionResponse.ID = permission.ID
	permissionResponse.Resource = permission.Resource
	permissionResponse.Permission = permission.Permission
	permissionResponse.ModifiedAt = permission.ModifiedAt
	permissionResponse.Version = permission.Version

	return &permissionResponse
}
