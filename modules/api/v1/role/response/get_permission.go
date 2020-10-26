package response

import "arkan-jaya/core/permission"

//GetPermissionResponse Get permission by tag response payload
type GetPermissionResponse struct {
	Permissions []*GetPermissionByIDResponse `json:"permissions"`
}

//NewGetPermissionResponse construct GetPermissionResponse
func NewGetPermissionResponse(permissions []permission.Permission) *GetPermissionResponse {
	var permissionResponses []*GetPermissionByIDResponse
	permissionResponses = make([]*GetPermissionByIDResponse, 0)

	for _, permission := range permissions {
		permissionResponses = append(permissionResponses, NewGetPermissionByIDResponse(permission))
	}

	return &GetPermissionResponse{
		permissionResponses,
	}
}
