package spec

//UpsertAuthSpec Credential to login
type UpsertAuthSpec struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}
