package response

import (
	"arkan-jaya/core/rolepermission"
	"time"
)

//GetRolePermissionByRoleIDResponse Get permission by ID response payload
type GetRolePermissionByRoleIDResponse struct {
	ID           string    `json:"id"`
	RoleID       string    `json:"role_id"`
	PermissionID string    `json:"permission_id"`
	ModifiedAt   time.Time `json:"modifiedAt"`
	Version      int       `json:"version"`
}

//NewGetRolePermissionByRoleIDResponse construct GetRolePermissionRoleByIDResponse
func NewGetRolePermissionByRoleIDResponse(permission rolepermission.RolePermission) *GetRolePermissionByRoleIDResponse {
	var permissionResponse GetRolePermissionByRoleIDResponse
	permissionResponse.PermissionID = permission.PermissionID
	permissionResponse.ModifiedAt = permission.ModifiedAt
	permissionResponse.Version = permission.Version

	return &permissionResponse
}
