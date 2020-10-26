package spec

//UpsertPermissionSpec create and update item spec
type UpsertPermissionSpec struct {
	Resource   string `validate:"required"`
	Permission string `validate:"required,min=3"`
}
