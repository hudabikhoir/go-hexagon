package spec

//UpsertRoleSpec create and update item spec
type UpsertRoleSpec struct {
	Name         string   `validate:"required"`
	PermissionID []string `validate:"required"`
}
