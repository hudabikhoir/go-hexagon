package response

import (
	"arkan-jaya/core/user"
	"time"
)

//GetUserByIDResponse Get user by ID response payload
type GetUserByIDResponse struct {
	ID         string `json:"_id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	CreatedAt  time.Time
	ModifiedAt time.Time
	RoleID     int  `json:"role_id"`
	IsActive   []string `json:"is_active"`
}

//NewGetUserByIDResponse construct GetUserByIDResponse
func NewGetUserByIDResponse(user user.User) *GetUserByIDResponse {
	var userResponse GetUserByIDResponse
	userResponse.ID = user.ID
	userResponse.Name = user.Name
	userResponse.Username = user.Username
	userResponse.Password = user.Password
	userResponse.Email = user.Email
	userResponse.ModifiedAt = user.ModifiedAt
	userResponse.CreatedAt = user.CreatedAt
	userResponse.RoleID = user.RoleID
	userResponse.IsActive = user.IsActive

	return &userResponse
}
