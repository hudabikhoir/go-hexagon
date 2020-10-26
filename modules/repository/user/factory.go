package user

import (
	"arkan-jaya/business/user"
	"arkan-jaya/util"
)

//RepositoryFactory Will return business.user.Repository based on active database connection
func RepositoryFactory(dbCon *util.DatabaseConnection) user.Repository {
	var userRepo user.Repository
	if dbCon.Driver == util.MySQL {
		userRepo = NewMySQLRepository(dbCon.MySQLDB)
	} else if dbCon.Driver == util.MongoDB {
		userRepo = NewMongoDBRepository(dbCon.MongoDB)
	}

	return userRepo
}
