package request

import (
	"arkan-jaya/business/rolepermission/spec"
)

//CreateRolePermissionRequest create permission request payload
type CreateRolePermissionRequest struct {
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
}

//ToUpsertRolePermissionSpec convert into permission.UpsertRolePermissionSpec object
func (req *CreateRolePermissionRequest) ToUpsertRolePermissionSpec() *spec.UpsertRolePermissionSpec {
	var upsertRolePermissionSpec spec.UpsertRolePermissionSpec
	upsertRolePermissionSpec.RoleID = req.RoleID
	upsertRolePermissionSpec.PermissionID = req.PermissionID

	return &upsertRolePermissionSpec
}
