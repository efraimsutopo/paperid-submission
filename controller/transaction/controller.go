package transaction

import (
	"net/http"
	"strconv"

	"github.com/efraimsutopo/paperid-submission/modules/customValidator"
	transactionService "github.com/efraimsutopo/paperid-submission/service/transaction"
	"github.com/efraimsutopo/paperid-submission/structs"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	GetAllInPagination(ec echo.Context) error
	GetTransactionByID(ec echo.Context) error
	CreateTransaction(ec echo.Context) error
	UpdateTransactionByID(ec echo.Context) error
	DeleteTransactionByID(ec echo.Context) error
}

type controller struct {
	service transactionService.Service
}

func New(service transactionService.Service) Controller {
	return &controller{
		service,
	}
}

func (c *controller) GetAllInPagination(ec echo.Context) error {

	var req structs.GetAllInPaginationRequest

	if err := ec.Bind(&req); err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}

	if err := ec.Validate(&req); err != nil {
		errValidator := customValidator.TransformValidatorError(err)
		return ec.JSON(errValidator.Code, errValidator)
	}

	res, errSvc := c.service.GetAllInPagination(ec, req)
	if errSvc != nil {
		return ec.JSON(errSvc.Code, errSvc)
	}

	return ec.JSON(http.StatusOK, res)
}

func (c *controller) GetTransactionByID(ec echo.Context) error {

	transactionID, err := strconv.ParseUint(ec.Param("transactionID"), 10, 64)
	if err != nil || transactionID == 0 {
		return ec.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "param transactionID is required",
		})
	}

	res, errSvc := c.service.GetTransactionByID(ec, transactionID)
	if errSvc != nil {
		return ec.JSON(errSvc.Code, errSvc)
	}

	return ec.JSON(http.StatusOK, res)
}

func (c *controller) CreateTransaction(ec echo.Context) error {
	var req structs.CreateTransactionRequest

	if err := ec.Bind(&req); err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}

	if err := ec.Validate(&req); err != nil {
		errValidator := customValidator.TransformValidatorError(err)
		return ec.JSON(errValidator.Code, errValidator)
	}

	res, err := c.service.CreateTransaction(ec, req)
	if err != nil {
		return ec.JSON(err.Code, err)
	}

	return ec.JSON(http.StatusCreated, res)
}

func (c *controller) UpdateTransactionByID(ec echo.Context) error {
	var req structs.UpdateTransactionRequest

	transactionID, err := strconv.ParseUint(ec.Param("transactionID"), 10, 64)
	if err != nil || transactionID == 0 {
		return ec.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "param transactionID is required",
		})
	}

	if err := ec.Bind(&req); err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}

	if err := ec.Validate(&req); err != nil {
		errValidator := customValidator.TransformValidatorError(err)
		return ec.JSON(errValidator.Code, errValidator)
	}

	req.ID = transactionID
	res, errSvc := c.service.UpdateTransactionByID(ec, req)
	if errSvc != nil {
		return ec.JSON(errSvc.Code, errSvc)
	}

	return ec.JSON(http.StatusCreated, res)
}

func (c *controller) DeleteTransactionByID(ec echo.Context) error {

	transactionID, err := strconv.ParseUint(ec.Param("transactionID"), 10, 64)
	if err != nil || transactionID == 0 {
		return ec.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "param transactionID is required",
		})
	}

	errSvc := c.service.DeleteTransactionByID(ec, transactionID)
	if errSvc != nil {
		return ec.JSON(errSvc.Code, errSvc)
	}

	return ec.JSON(http.StatusNoContent, nil)
}
