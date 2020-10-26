package response

import "arkan-jaya/core/rolepermission"

//GetRolePermissionResponse Get permission by tag response payload
type GetRolePermissionResponse struct {
	Permissions []*GetRolePermissionByRoleIDResponse `json:"permissions"`
}

//NewGetPermissionResponse construct GetPermissionResponse
func NewGetPermissionResponse(rolePermissions []rolepermission.RolePermission) *GetRolePermissionResponse {
	var rolePermissionResponses []*GetRolePermissionByRoleIDResponse
	rolePermissionResponses = make([]*GetRolePermissionByRoleIDResponse, 0)

	for _, rolePermission := range rolePermissions {
		rolePermissionResponses = append(rolePermissionResponses, NewGetRolePermissionByRoleIDResponse(rolePermission))
	}

	return &GetRolePermissionResponse{
		rolePermissionResponses,
	}
}
