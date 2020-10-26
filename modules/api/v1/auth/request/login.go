package request

import (
	"arkan-jaya/business/auth/spec"
)

//CreateAuthRequest create user request payload
type CreateAuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//ToUpsertAuthSpec convert into user.UpsertAuthSpec object
func (req *CreateAuthRequest) ToUpsertAuthSpec() *spec.UpsertAuthSpec {
	var upsertAuthSpec spec.UpsertAuthSpec
	upsertAuthSpec.Username = req.Username
	upsertAuthSpec.Password = req.Password

	return &upsertAuthSpec
}
