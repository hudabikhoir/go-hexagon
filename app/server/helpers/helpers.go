package helpers

import (
	"arkan-jaya/modules/api"
	"arkan-jaya/modules/api/common"
	"arkan-jaya/util"

	userBussiness "arkan-jaya/business/user"
	userCtrlV1 "arkan-jaya/modules/api/v1/user"
	userRepo "arkan-jaya/modules/repository/user"

	authBussiness "arkan-jaya/business/auth"
	authCtrlV1 "arkan-jaya/modules/api/v1/auth"
	authRepo "arkan-jaya/modules/repository/auth"

	permissionBussiness "arkan-jaya/business/permission"
	permissionCtrlV1 "arkan-jaya/modules/api/v1/permission"
	permissionRepo "arkan-jaya/modules/repository/permission"

	echo "github.com/labstack/echo/v4"
)

func SetErrorHandler(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		// error message must be known RC value
		errResp := common.NewErrorResponse(err.Error(), []string{})
		c.JSON(errResp.HttpStatus, errResp)
	}
}

func RegisterController(dbCon *util.DatabaseConnection) api.Controller {
	//initiate user
	userPermitRepo := userRepo.RepositoryFactory(dbCon)
	userPermitService := userBussiness.NewService(userPermitRepo)
	userPermitControllerV1 := userCtrlV1.NewController(userPermitService)

	//initiate auth
	authPermitRepo := authRepo.RepositoryFactory(dbCon)
	authPermitService := authBussiness.NewService(authPermitRepo)
	authPermitControllerV1 := authCtrlV1.NewController(authPermitService)

	//initiate permission
	permissionPermitRepo := permissionRepo.RepositoryFactory(dbCon)
	permissionPermitService := permissionBussiness.NewService(permissionPermitRepo)
	permissionPermitControllerV1 := permissionCtrlV1.NewController(permissionPermitService)

	//lets put the controller together
	controllers := api.Controller{
		UserController:       userPermitControllerV1,
		AuthController:       authPermitControllerV1,
		PermissionController: permissionPermitControllerV1,
	}

	return controllers
}
