package request

import "arkan-jaya/business/permission/spec"

//UpdatePermissionRequest update permission request payload
type UpdatePermissionRequest struct {
	Resource   string `json:"name"`
	Permission string `json:"description"`
	Version    int    `json:"version" validate:"required"`
}

//ToUpsertPermissionSpec convert into permission.UpsertPermissionSpec object
func (req *UpdatePermissionRequest) ToUpsertPermissionSpec() *spec.UpsertPermissionSpec {
	var upsertPermissionSpec spec.UpsertPermissionSpec
	upsertPermissionSpec.Resource = req.Resource
	upsertPermissionSpec.Permission = req.Permission

	return &upsertPermissionSpec
}
