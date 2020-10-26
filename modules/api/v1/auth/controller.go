package auth

import (
	"arkan-jaya/business"
	authBusiness "arkan-jaya/business/auth"
	"arkan-jaya/modules/api/common"
	"arkan-jaya/modules/api/v1/auth/request"
	"arkan-jaya/modules/api/v1/auth/response"
	"net/http"

	v10 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

//Controller Get auth API controller
type Controller struct {
	service   authBusiness.Service
	validator *v10.Validate
}

//NewController Construct auth API controller
func NewController(service authBusiness.Service) *Controller {
	return &Controller{
		service,
		v10.New(),
	}
}

//login Create new auth echo handler
func (controller *Controller) Login(c echo.Context) error {
	createAuthRequest := new(request.CreateAuthRequest)

	if err := c.Bind(createAuthRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	token, err := controller.service.LoginHandler(*createAuthRequest.ToUpsertAuthSpec())

	if err != nil {
		if err == business.ErrInvalidSpec {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	response := response.NewCreateNewUserResponse(token)
	return c.JSON(http.StatusCreated, response)
}
