package user

import "time"

//User product user
type User struct {
	ID         string
	Name       string
	Username   string
	Password   string
	Email      string
	CreatedAt  time.Time
	ModifiedAt time.Time
	RoleID     int
	IsActive   []string
}

//NewUser create new user
func NewUser(
	id string,
	name string,
	username string,
	password string,
	email string,
	roleID int,
	permissionID []string,
	createdAt time.Time) User {

	return User{
		ID:         id,
		Name:       name,
		Username:   username,
		Password:   password,
		Email:      email,
		RoleID:     roleID,
		IsActive:   permissionID,
		CreatedAt:  createdAt,
		ModifiedAt: createdAt,
	}
}

//ModifyUser update existing user data
func (oldUser *User) ModifyUser(newName string, newUsername string, newPassword string, newEmail string, newRoleID int, newIsActive []string, modifiedAt time.Time) User {
	return User{
		ID:         oldUser.ID,
		Name:       newName,
		Username:   newUsername,
		Password:   newPassword,
		Email:      newEmail,
		RoleID:     newRoleID,
		IsActive:   newIsActive,
		CreatedAt:  oldUser.CreatedAt,
		ModifiedAt: modifiedAt,
	}
}
