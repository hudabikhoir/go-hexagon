package user

import (
	"arkan-jaya/business"
	userBusiness "arkan-jaya/business/user"
	"arkan-jaya/modules/api/common"
	"arkan-jaya/modules/api/v1/user/request"
	"arkan-jaya/modules/api/v1/user/response"
	"net/http"

	v10 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

//Controller Get user API controller
type Controller struct {
	service   userBusiness.Service
	validator *v10.Validate
}

//NewController Construct user API controller
func NewController(service userBusiness.Service) *Controller {
	return &Controller{
		service,
		v10.New(),
	}
}

//CreateNewUser Create new user echo handler
func (controller *Controller) CreateNewUser(c echo.Context) error {
	createUserRequest := new(request.CreateUserRequest)

	if err := c.Bind(createUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	ID, err := controller.service.CreateUser(*createUserRequest.ToUpsertUserSpec(), "creator")

	if err != nil {
		if err == business.ErrInvalidSpec {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	response := response.NewCreateNewUserResponse(ID)
	return c.JSON(http.StatusCreated, response)
}

//GetUserByID Get item by ID echo handler
func (controller *Controller) GetUserByID(c echo.Context) error {
	ID := c.Param("id")
	item, err := controller.service.GetUserByID(ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	} else if item == nil {
		return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
	}

	response := response.NewGetUserByIDResponse(*item)
	return c.JSON(http.StatusOK, response)
}

//GetUsers Find item by tag echo handler
func (controller *Controller) GetUsers(c echo.Context) error {
	items, err := controller.service.GetUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	response := response.NewGetUsersResponse(items)
	return c.JSON(http.StatusOK, response)
}

//UpdateUser update item echo handler
func (controller *Controller) UpdateUser(c echo.Context) error {
	updateUserRequest := new(request.UpdateUserRequest)

	if err := c.Bind(updateUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	err := controller.validator.Struct(updateUserRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	err = controller.service.UpdateUser(
		c.Param("id"),
		*updateUserRequest.ToUpsertUserSpec(),
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
