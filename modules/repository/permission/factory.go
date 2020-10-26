package permission

import (
	"arkan-jaya/business/permission"
	"arkan-jaya/util"
)

//RepositoryFactory Will return business.permission.Repository based on active database connection
func RepositoryFactory(dbCon *util.DatabaseConnection) permission.Repository {
	var permissionRepo permission.Repository
	if dbCon.Driver == util.MySQL {
		permissionRepo = NewMySQLRepository(dbCon.MySQLDB)
	} else if dbCon.Driver == util.MongoDB {
		permissionRepo = NewMongoDBRepository(dbCon.MongoDB)
	}

	return permissionRepo
}
