package auth

import (
	"arkan-jaya/business/auth"
	"arkan-jaya/util"
)

//RepositoryFactory Will return business.auth.Repository based on active database connection
func RepositoryFactory(dbCon *util.DatabaseConnection) auth.Repository {
	var authRepo auth.Repository
	if dbCon.Driver == util.MySQL {
		authRepo = NewMySQLRepository(dbCon.MySQLDB)
	} else if dbCon.Driver == util.MongoDB {
		authRepo = NewMongoDBRepository(dbCon.MongoDB)
	}

	return authRepo
}
