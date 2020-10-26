package role

import (
	"arkan-jaya/business/role"
	"arkan-jaya/util"
)

//RepositoryFactory Will return business.role.Repository based on active database connection
func RepositoryFactory(dbCon *util.DatabaseConnection) role.Repository {
	var roleRepo role.Repository
	if dbCon.Driver == util.MySQL {
		roleRepo = NewMySQLRepository(dbCon.MySQLDB)
	} else if dbCon.Driver == util.MongoDB {
		roleRepo = NewMongoDBRepository(dbCon.MongoDB)
	}

	return roleRepo
}
