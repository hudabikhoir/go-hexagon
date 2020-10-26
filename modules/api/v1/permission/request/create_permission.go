package request

import (
	"arkan-jaya/business/permission/spec"
)

//CreatePermissionRequest create permission request payload
type CreatePermissionRequest struct {
	Resource   string `json:"resource"`
	Permission string `json:"permission"`
}

//ToUpsertPermissionSpec convert into permission.UpsertPermissionSpec object
func (req *CreatePermissionRequest) ToUpsertPermissionSpec() *spec.UpsertPermissionSpec {
	var upsertPermissionSpec spec.UpsertPermissionSpec
	upsertPermissionSpec.Resource = req.Resource
	upsertPermissionSpec.Permission = req.Permission

	return &upsertPermissionSpec
}
