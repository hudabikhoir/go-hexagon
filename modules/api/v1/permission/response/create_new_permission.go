package response

//CreateNewPermissionResponse Create item response payload
type CreateNewPermissionResponse struct {
	ID string `json:"id"`
}

//NewCreateNewPermissionResponse construct CreateNewPermissionResponse
func NewCreateNewPermissionResponse(id string) *CreateNewPermissionResponse {
	return &CreateNewPermissionResponse{
		id,
	}
}
