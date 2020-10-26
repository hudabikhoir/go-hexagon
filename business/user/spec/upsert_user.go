package spec

//UpsertUserSpec create and update item spec
type UpsertUserSpec struct {
	Name     string `validate:"required"`
	Username string `validate:"required,min=3"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
	RoleID   int    `validate:"required"`
	IsActive bool   `validate:"required"`
}
