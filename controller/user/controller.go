package user

import (
	"net/http"

	"github.com/efraimsutopo/paperid-submission/modules/customValidator"
	userService "github.com/efraimsutopo/paperid-submission/service/user"
	"github.com/efraimsutopo/paperid-submission/structs"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	Register(ec echo.Context) error
	Login(ec echo.Context) error
	Logout(ec echo.Context) error
	Profile(ec echo.Context) error
	CheckValidSession(ec echo.Context) error
}

type controller struct {
	service userService.Service
}

func New(service userService.Service) Controller {
	return &controller{
		service,
	}
}

func (c *controller) Register(ec echo.Context) error {
	var req structs.RegisterUserRequest

	if err := ec.Bind(&req); err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}

	if err := ec.Validate(&req); err != nil {
		errValidator := customValidator.TransformValidatorError(err)
		return ec.JSON(errValidator.Code, errValidator)
	}

	res, errSvc := c.service.Register(ec, req)
	if errSvc != nil {
		return ec.JSON(errSvc.Code, errSvc)
	}

	return ec.JSON(http.StatusOK, res)
}

func (c *controller) Login(ec echo.Context) error {
	var req structs.LoginRequest

	if err := ec.Bind(&req); err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}

	if err := ec.Validate(&req); err != nil {
		errValidator := customValidator.TransformValidatorError(err)
		return ec.JSON(errValidator.Code, errValidator)
	}

	res, errSvc := c.service.Login(ec, req)
	if errSvc != nil {
		return ec.JSON(errSvc.Code, errSvc)
	}

	return ec.JSON(http.StatusOK, res)
}

func (c *controller) Logout(ec echo.Context) error {
	errSvc := c.service.Logout(ec)
	if errSvc != nil {
		return ec.JSON(errSvc.Code, errSvc)
	}

	return ec.JSON(http.StatusOK, structs.MessageResponse{
		Message: "successfully logout",
	})
}

func (c *controller) Profile(ec echo.Context) error {
	res, errSvc := c.service.GetUser(ec)
	if errSvc != nil {
		return ec.JSON(errSvc.Code, errSvc)
	}

	return ec.JSON(http.StatusOK, res)
}

func (c *controller) CheckValidSession(ec echo.Context) error {
	return c.service.CheckValidSession(ec)
}
