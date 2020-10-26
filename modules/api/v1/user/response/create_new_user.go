package response

//CreateNewUserResponse Create item response payload
type CreateNewUserResponse struct {
	ID string `json:"id"`
}

//NewCreateNewUserResponse construct CreateNewUserResponse
func NewCreateNewUserResponse(id string) *CreateNewUserResponse {
	return &CreateNewUserResponse{
		id,
	}
}
