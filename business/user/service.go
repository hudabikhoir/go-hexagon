package user

import (
	"arkan-jaya/business"
	"arkan-jaya/business/user/spec"
	core "arkan-jaya/core/user"
	"arkan-jaya/util"
	"time"

	validator "github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

//Service outgoing port for user
type Service interface {
	GetUsers() ([]core.User, error)

	GetUserByID(ID string) (*core.User, error)

	CreateUser(upsertuserSpec spec.UpsertUserSpec, createdBy string) (string, error)

	UpdateUser(ID string, upsertuserSpec spec.UpsertUserSpec, modifiedBy string) error
}

//=============== The implementation of those interface put below =======================

type service struct {
	repository Repository
	validate   *validator.Validate
}

//NewService Construct user service object
func NewService(repository Repository) Service {
	return &service{
		repository,
		validator.New(),
	}
}

//GetUserByID Get user by given ID, return nil if not exist
func (s *service) GetUserByID(ID string) (*core.User, error) {
	return s.repository.FindUserByID(ID)
}

//GetUsers Get all users, return zero array if not match
func (s *service) GetUsers() ([]core.User, error) {

	users, err := s.repository.FindAllUser()
	if err != nil || users == nil {
		return []core.User{}, err
	}

	return users, err
}

//CreateUser Create new user and store into database
func (s *service) CreateUser(upsertuserSpec spec.UpsertUserSpec, createdBy string) (string, error) {
	err := s.validate.Struct(upsertuserSpec)

	if err != nil {
		return "", business.ErrInvalidSpec
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(upsertuserSpec.Password), bcrypt.DefaultCost)

	ID := util.GenerateID()
	user := core.NewUser(
		ID,
		upsertuserSpec.Name,
		upsertuserSpec.Username,
		string(hashedPassword),
		upsertuserSpec.Email,
		upsertuserSpec.RoleID,
		upsertuserSpec.IsActive,
		time.Now(),
	)

	err = s.repository.InsertUser(user)
	if err != nil {
		return "", err
	}

	return ID, nil
}

//UpdateUser Update existing user in the database.
//Will return ErrNotFound when user is not exists or ErrConflict if data version is not match
func (s *service) UpdateUser(ID string, upsertuserSpec spec.UpsertUserSpec, modifiedBy string) error {
	err := s.validate.Struct(upsertuserSpec)

	if err != nil || len(ID) == 0 {
		return business.ErrInvalidSpec
	}

	//get the user first to make sure data is exist
	user, err := s.repository.FindUserByID(ID)

	if err != nil {
		return err
	} else if user == nil {
		return business.ErrNotFound
	}

	newUser := user.ModifyUser(upsertuserSpec.Name, upsertuserSpec.Username, upsertuserSpec.Password, upsertuserSpec.Email, upsertuserSpec.RoleID, upsertuserSpec.IsActive, time.Now())

	return s.repository.UpdateUser(newUser)
}
