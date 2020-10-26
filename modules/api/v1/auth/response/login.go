package response

//CreateNewUserResponse Create item response payload
type CreateNewUserResponse struct {
	Token string `json:"token"`
}

//NewCreateNewUserResponse construct CreateNewUserResponse
func NewCreateNewUserResponse(token string) *CreateNewUserResponse {
	return &CreateNewUserResponse{
		token,
	}
}
