package rolepermission

import (
	"arkan-jaya/business"
	rolepermissionBusiness "arkan-jaya/business/rolepermission"
	"arkan-jaya/modules/api/common"
	"arkan-jaya/modules/api/v1/rolepermission/request"
	"arkan-jaya/modules/api/v1/rolepermission/response"
	"net/http"

	v10 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

//Controller Get permission API controller
type Controller struct {
	service   rolepermissionBusiness.Service
	validator *v10.Validate
}

//NewController Construct permission API controller
func NewController(service rolepermissionBusiness.Service) *Controller {
	return &Controller{
		service,
		v10.New(),
	}
}

//FindRolePermission Find all role permission echo handler
func (controller *Controller) FindRolePermission(c echo.Context) error {
	rolePermissions, err := controller.service.GetRolePermissions()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	response := response.NewGetPermissionResponse(rolePermissions)
	return c.JSON(http.StatusOK, response)
}

//CreateNewRolePermission Create new permission echo handler
func (controller *Controller) CreateNewRolePermission(c echo.Context) error {
	createRolePermissionRequest := new(request.CreateRolePermissionRequest)

	if err := c.Bind(createRolePermissionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	ID, err := controller.service.CreateRolePermission(*createRolePermissionRequest.ToUpsertRolePermissionSpec(), "creator")

	if err != nil {
		if err == business.ErrInvalidSpec {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	response := response.NewCreateNewRolePermissionResponse(ID)
	return c.JSON(http.StatusCreated, response)
}

//UpdateRolePermission update permission echo handler
func (controller *Controller) UpdateRolePermission(c echo.Context) error {
	updateRolePermissionRequest := new(request.UpdateRolePermissionRequest)

	if err := c.Bind(updateRolePermissionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	err := controller.validator.Struct(updateRolePermissionRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	err = controller.service.UpdatePermission(
		c.Param("id"),
		*updateRolePermissionRequest.ToUpsertPermissionSpec(),
		updateRolePermissionRequest.Version,
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

//DeleteRolePermission update permission echo handler
func (controller *Controller) DeleteRolePermission(c echo.Context) error {
	err := controller.service.DeleteRolePermission(
		c.Param("id"),
		"deleter")

	if err != nil {
		if err == business.ErrNotFound {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
	}

	return c.NoContent(http.StatusNoContent)
}
