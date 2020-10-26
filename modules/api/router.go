package api

import (
	"arkan-jaya/modules/api/v1/auth"
	"arkan-jaya/modules/api/v1/permission"
	"arkan-jaya/modules/api/v1/user"

	"github.com/labstack/echo"
)

// Controller to define controller that we use
type Controller struct {
	UserController       *user.Controller
	AuthController       *auth.Controller
	PermissionController *permission.Controller
}

//RegisterPath Registera V1 API path
func RegisterPath(e *echo.Echo, ctrl Controller) {
	authV1 := e.Group("v1/auth")
	authV1.POST("/login", ctrl.AuthController.Login)

	// user
	userV1 := e.Group("v1/users")
	userV1.GET("", ctrl.UserController.GetUsers)
	userV1.GET("/:id", ctrl.UserController.GetUserByID)
	userV1.POST("", ctrl.UserController.CreateNewUser)
	userV1.PUT("/:id", ctrl.UserController.UpdateUser)

	// permission
	permissionV1 := e.Group("v1/permissions")
	permissionV1.GET("", ctrl.PermissionController.FindPermission)
	permissionV1.POST("", ctrl.PermissionController.CreateNewPermission)
	permissionV1.PUT("/:id", ctrl.PermissionController.UpdatePermission)
	permissionV1.DELETE("/:id", ctrl.PermissionController.DeletePermission)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
