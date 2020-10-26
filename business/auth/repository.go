package auth

import core "arkan-jaya/core/user"

//Repository ingoing port for item
type Repository interface {
	//FindUserByID If data not found will return nil without error
	FindUserByUsername(username string) (*core.User, error)
}
