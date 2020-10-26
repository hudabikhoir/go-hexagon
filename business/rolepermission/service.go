package rolepermission

import (
	"arkan-jaya/business"
	"arkan-jaya/business/rolepermission/spec"
	core "arkan-jaya/core/rolepermission"
	"arkan-jaya/util"
	"time"

	validator "github.com/go-playground/validator/v10"
)

//Service outgoing port for role
type Service interface {
	GetRolePermissions() ([]core.RolePermission, error)

	CreateRolePermission(upsertroleSpec spec.UpsertRolePermissionSpec, createdBy string) (string, error)

	UpdateRolePermission(ID string, upsertroleSpec spec.UpsertRolePermissionSpec, currentVersion int, modifiedBy string) error

	DeleteRolePermission(ID string, modifiedBy string) error
}

//=============== The implementation of those interface put below =======================

type service struct {
	repository Repository
	validate   *validator.Validate
}

//NewService Construct rolePermission service object
func NewService(repository Repository) Service {
	return &service{
		repository,
		validator.New(),
	}
}

//GetRolePermissions Get all rolePermissions by given tag, return zero array if not match
func (s *service) GetRolePermissions() ([]core.RolePermission, error) {

	rolePermissions, err := s.repository.GetAll()
	if err != nil || rolePermissions == nil {
		return []core.RolePermission{}, err
	}

	return rolePermissions, err
}

//CreateRolePermission Create new rolePermission and store into database
func (s *service) CreateRolePermission(upsertrolePermissionSpec spec.UpsertRolePermissionSpec, createdBy string) (string, error) {
	err := s.validate.Struct(upsertrolePermissionSpec)

	if err != nil {
		return "", business.ErrInvalidSpec
	}

	ID := util.GenerateID()
	rolePermission := core.NewRolePermission(
		upsertrolePermissionSpec.RoleID,
		upsertrolePermissionSpec.PermissionID,
		createdBy,
		time.Now(),
	)

	err = s.repository.InsertRolePermission(rolePermission)
	if err != nil {
		return "", err
	}

	return ID, nil
}

//UpdateRolePermission Update existing rolePermission in the database.
//Will return ErrNotFound when rolePermission is not exists or ErrConflict if data version is not match
func (s *service) UpdateRolePermission(ID string, upsertrolePermissionSpec spec.UpsertRolePermissionSpec, currentVersion int, modifiedBy string) error {
	err := s.validate.Struct(upsertrolePermissionSpec)

	if err != nil || len(ID) == 0 {
		return business.ErrInvalidSpec
	}

	//get the rolePermission first to make sure data is exist
	rolePermission, err := s.repository.FindRolePermissionByRoleID(ID)

	if err != nil {
		return err
	} else if rolePermission == nil {
		return business.ErrNotFound
	} else if rolePermission.Version != currentVersion {
		return business.ErrHasBeenModified
	}

	newRolesPermission := rolePermission.ModifyRolePermission(upsertrolePermissionSpec.RoleID, upsertrolePermissionSpec.PermissionID, modifiedBy, time.Now())

	return s.repository.UpdateRolePermission(newRolesPermission, currentVersion)
}

//DeleteRolePermission Delete existing rolePermission in the database.
//Will return ErrNotFound when rolePermission is not exists or ErrConflict if data version is not match
func (s *service) DeleteRolePermission(ID string, modifiedBy string) error {
	if len(ID) == 0 {
		return business.ErrInvalidSpec
	}

	//get the rolePermission first to make sure data is exist
	rolePermission, err := s.repository.FindRolePermissionByRoleID(ID)

	if err != nil {
		return err
	} else if rolePermission == nil {
		return business.ErrNotFound
	}

	return s.repository.DeleteRolePermission(ID)
}
