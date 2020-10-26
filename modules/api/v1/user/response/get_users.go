package response

import "arkan-jaya/core/user"

//GetUsersResponse Get user by tag response payload
type GetUsersResponse struct {
	Users []*GetUserByIDResponse `json:"users"`
}

//NewGetUsersResponse construct GetUsersResponse
func NewGetUsersResponse(users []user.User) *GetUsersResponse {
	var userResponses []*GetUserByIDResponse
	userResponses = make([]*GetUserByIDResponse, 0)

	for _, user := range users {
		userResponses = append(userResponses, NewGetUserByIDResponse(user))
	}

	return &GetUsersResponse{
		userResponses,
	}
}
