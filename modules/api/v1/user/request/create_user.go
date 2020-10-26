package request

import (
	"arkan-jaya/business/user/spec"
)

//CreateUserRequest create user request payload
type CreateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
	IsActive []string   `json:"is_active"`
}

//ToUpsertUserSpec convert into user.UpsertUserSpec object
func (req *CreateUserRequest) ToUpsertUserSpec() *spec.UpsertUserSpec {
	var upsertUserSpec spec.UpsertUserSpec
	upsertUserSpec.Name = req.Name
	upsertUserSpec.Username = req.Username
	upsertUserSpec.Email = req.Email
	upsertUserSpec.Password = req.Password
	upsertUserSpec.RoleID = req.RoleID
	upsertUserSpec.IsActive = req.IsActive

	return &upsertUserSpec
}
