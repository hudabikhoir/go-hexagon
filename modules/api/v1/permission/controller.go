package permission

import (
	"arkan-jaya/business"
	permissionBusiness "arkan-jaya/business/permission"
	"arkan-jaya/modules/api/common"
	"arkan-jaya/modules/api/v1/permission/request"
	"arkan-jaya/modules/api/v1/permission/response"
	"net/http"

	v10 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

//Controller Get permission API controller
type Controller struct {
	service   permissionBusiness.Service
	validator *v10.Validate
}

//NewController Construct permission API controller
func NewController(service permissionBusiness.Service) *Controller {
	return &Controller{
		service,
		v10.New(),
	}
}

//FindPermission Find all permission echo handler
func (controller *Controller) FindPermission(c echo.Context) error {
	permissions, err := controller.service.GetPermissions()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	response := response.NewGetPermissionResponse(permissions)
	return c.JSON(http.StatusOK, response)
}

//CreateNewPermission Create new permission echo handler
func (controller *Controller) CreateNewPermission(c echo.Context) error {
	createPermissionRequest := new(request.CreatePermissionRequest)

	if err := c.Bind(createPermissionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	ID, err := controller.service.CreatePermission(*createPermissionRequest.ToUpsertPermissionSpec(), "creator")

	if err != nil {
		if err == business.ErrInvalidSpec {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	response := response.NewCreateNewPermissionResponse(ID)
	return c.JSON(http.StatusCreated, response)
}

//UpdatePermission update permission echo handler
func (controller *Controller) UpdatePermission(c echo.Context) error {
	updatePermissionRequest := new(request.UpdatePermissionRequest)

	if err := c.Bind(updatePermissionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	err := controller.validator.Struct(updatePermissionRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	err = controller.service.UpdatePermission(
		c.Param("id"),
		*updatePermissionRequest.ToUpsertPermissionSpec(),
		updatePermissionRequest.Version,
		"updater")

	if err != nil {
		if err == business.ErrNotFound {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
		if err == business.ErrHasBeenModified {
			return c.JSON(http.StatusConflict, common.NewConflictResponse())
		}
	}

	return c.NoContent(http.StatusNoContent)
}

//DeletePermission update permission echo handler
func (controller *Controller) DeletePermission(c echo.Context) error {
	err := controller.service.DeletePermission(
		c.Param("id"),
		"deleter")

	if err != nil {
		if err == business.ErrNotFound {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
	}

	return c.NoContent(http.StatusNoContent)
}
