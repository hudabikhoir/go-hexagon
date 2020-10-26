package permission

import (
	"arkan-jaya/business"
	"arkan-jaya/business/permission/spec"
	core "arkan-jaya/core/permission"
	"arkan-jaya/util"
	"fmt"
	"time"

	validator "github.com/go-playground/validator/v10"
)

//Service outgoing port for permission
type Service interface {
	GetPermissions() ([]core.Permission, error)

	CreatePermission(upsertpermissionSpec spec.UpsertPermissionSpec, createdBy string) (string, error)

	UpdatePermission(ID string, upsertpermissionSpec spec.UpsertPermissionSpec, currentVersion int, modifiedBy string) error

	DeletePermission(ID string, modifiedBy string) error
}

//=============== The implementation of those interface put below =======================

type service struct {
	repository Repository
	validate   *validator.Validate
}

//NewService Construct permission service object
func NewService(repository Repository) Service {
	return &service{
		repository,
		validator.New(),
	}
}

//GetPermissions Get all permissions by given tag, return zero array if not match
func (s *service) GetPermissions() ([]core.Permission, error) {

	permissions, err := s.repository.GetAll()
	if err != nil || permissions == nil {
		return []core.Permission{}, err
	}

	return permissions, err
}

//CreatePermission Create new permission and store into database
func (s *service) CreatePermission(upsertpermissionSpec spec.UpsertPermissionSpec, createdBy string) (string, error) {
	err := s.validate.Struct(upsertpermissionSpec)

	if err != nil {
		return "", business.ErrInvalidSpec
	}

	ID := util.GenerateID()
	permission := core.NewPermission(
		ID,
		upsertpermissionSpec.Resource,
		upsertpermissionSpec.Permission,
		createdBy,
		time.Now(),
	)

	err = s.repository.InsertPermission(permission)
	fmt.Println(err)
	if err != nil {
		return "", err
	}

	return ID, nil
}

//UpdatePermission Update existing permission in the database.
//Will return ErrNotFound when permission is not exists or ErrConflict if data version is not match
func (s *service) UpdatePermission(ID string, upsertpermissionSpec spec.UpsertPermissionSpec, currentVersion int, modifiedBy string) error {
	err := s.validate.Struct(upsertpermissionSpec)

	if err != nil || len(ID) == 0 {
		return business.ErrInvalidSpec
	}

	//get the permission first to make sure data is exist
	permission, err := s.repository.FindPermissionByID(ID)

	if err != nil {
		return err
	} else if permission == nil {
		return business.ErrNotFound
	} else if permission.Version != currentVersion {
		return business.ErrHasBeenModified
	}

	newPermission := permission.ModifyPermission(upsertpermissionSpec.Resource, upsertpermissionSpec.Permission, modifiedBy, time.Now())

	return s.repository.UpdatePermission(newPermission, currentVersion)
}

//DeletePermission Delete existing permission in the database.
//Will return ErrNotFound when permission is not exists or ErrConflict if data version is not match
func (s *service) DeletePermission(ID string, modifiedBy string) error {
	if len(ID) == 0 {
		return business.ErrInvalidSpec
	}

	//get the permission first to make sure data is exist
	permission, err := s.repository.FindPermissionByID(ID)

	if err != nil {
		return err
	} else if permission == nil {
		return business.ErrNotFound
	}

	return s.repository.DeletePermission(ID)
}
