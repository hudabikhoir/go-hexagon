package request

import "arkan-jaya/business/rolepermission/spec"

//UpdateRolePermissionRequest update permission request payload
type UpdateRolePermissionRequest struct {
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
	Version      int    `json:"version" validate:"required"`
}

//ToUpsertRolePermissionSpec convert into permission.UpsertRolePermissionSpec object
func (req *UpdateRolePermissionRequest) ToUpsertRolePermissionSpec() *spec.UpsertRolePermissionSpec {
	var upsertRolePermissionSpec spec.UpsertRolePermissionSpec
	upsertRolePermissionSpec.RoleID = req.RoleID
	upsertRolePermissionSpec.PermissionID = req.PermissionID

	return &upsertRolePermissionSpec
}
