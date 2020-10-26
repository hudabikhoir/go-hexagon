package rolepermission

import (
	"arkan-jaya/business/rolepermission"
	"arkan-jaya/util"
)

//RepositoryFactory Will return business.rolePermission.Repository based on active database connection
func RepositoryFactory(dbCon *util.DatabaseConnection) rolepermission.Repository {
	var rolePermissionRepo rolepermission.Repository
	if dbCon.Driver == util.MySQL {
		rolePermissionRepo = NewMySQLRepository(dbCon.MySQLDB)
	} else if dbCon.Driver == util.MongoDB {
		rolePermissionRepo = NewMongoDBRepository(dbCon.MongoDB)
	}

	return rolePermissionRepo
}
