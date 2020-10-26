package user

import core "arkan-jaya/core/user"

//Repository ingoing port for item
type Repository interface {
	//FindAllUser If data not found will return nil without error
	FindAllUser() ([]core.User, error)

	//FindUserByID If data not found will return nil without error
	FindUserByID(ID string) (*core.User, error)

	//InsertUser Register new user
	InsertUser(user core.User) error

	//UpdateUser if data not found will return core.ErrZeroAffected
	UpdateUser(user core.User) error
}
