package spec

//UpsertRoleSpec create and update item spec
type UpsertRolePermissionSpec struct {
	RoleID       string `validate:"required"`
	PermissionID string `validate:"required"`
}
