package response

//CreateNewRolePermissionResponse Create item response payload
type CreateNewRolePermissionResponse struct {
	ID string `json:"id"`
}

//NewCreateNewRolePermissionResponse construct CreateNewRolePermissionResponse
func NewCreateNewRolePermissionResponse(id string) *CreateNewRolePermissionResponse {
	return &CreateNewRolePermissionResponse{
		id,
	}
}
