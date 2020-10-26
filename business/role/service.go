package role

import (
	"arkan-jaya/business"
	"arkan-jaya/business/role/spec"
	core "arkan-jaya/core/role"
	"arkan-jaya/util"
	"time"

	validator "github.com/go-playground/validator/v10"
)

//Service outgoing port for role
type Service interface {
	GetRoles() ([]core.Role, error)

	CreateRole(upsertroleSpec spec.UpsertRoleSpec, createdBy string) (string, error)

	UpdateRole(ID string, upsertroleSpec spec.UpsertRoleSpec, currentVersion int, modifiedBy string) error

	DeleteRole(ID string, modifiedBy string) error
}

//=============== The implementation of those interface put below =======================

type service struct {
	repository Repository
	validate   *validator.Validate
}

//NewService Construct role service object
func NewService(repository Repository) Service {
	return &service{
		repository,
		validator.New(),
	}
}

//GetRoles Get all roles by given tag, return zero array if not match
func (s *service) GetRoles() ([]core.Role, error) {

	roles, err := s.repository.GetAll()
	if err != nil || roles == nil {
		return []core.Role{}, err
	}

	return roles, err
}

//CreateRole Create new role and store into database
func (s *service) CreateRole(upsertroleSpec spec.UpsertRoleSpec, createdBy string) (string, error) {
	err := s.validate.Struct(upsertroleSpec)

	if err != nil {
		return "", business.ErrInvalidSpec
	}

	ID := util.GenerateID()
	role := core.NewRole(
		ID,
		upsertroleSpec.Name,
		upsertroleSpec.PermissionID,
		createdBy,
		time.Now(),
	)

	err = s.repository.InsertRole(role)
	if err != nil {
		return "", err
	}

	return ID, nil
}

//UpdateRole Update existing role in the database.
//Will return ErrNotFound when role is not exists or ErrConflict if data version is not match
func (s *service) UpdateRole(ID string, upsertroleSpec spec.UpsertRoleSpec, currentVersion int, modifiedBy string) error {
	err := s.validate.Struct(upsertroleSpec)

	if err != nil || len(ID) == 0 {
		return business.ErrInvalidSpec
	}

	//get the role first to make sure data is exist
	role, err := s.repository.FindRoleByID(ID)

	if err != nil {
		return err
	} else if role == nil {
		return business.ErrNotFound
	} else if role.Version != currentVersion {
		return business.ErrHasBeenModified
	}

	newRole := role.ModifyRole(upsertroleSpec.Name, upsertroleSpec.PermissionID, modifiedBy, time.Now())

	return s.repository.UpdateRole(newRole, currentVersion)
}

//DeleteRole Delete existing role in the database.
//Will return ErrNotFound when role is not exists or ErrConflict if data version is not match
func (s *service) DeleteRole(ID string, modifiedBy string) error {
	if len(ID) == 0 {
		return business.ErrInvalidSpec
	}

	//get the role first to make sure data is exist
	role, err := s.repository.FindRoleByID(ID)

	if err != nil {
		return err
	} else if role == nil {
		return business.ErrNotFound
	}

	return s.repository.DeleteRole(ID)
}
