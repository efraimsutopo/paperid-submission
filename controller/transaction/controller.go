package transaction

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/efraimsutopo/paperid-submission/modules/customValidator"
	transactionService "github.com/efraimsutopo/paperid-submission/service/transaction"
	"github.com/efraimsutopo/paperid-submission/structs"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	GetAllInPagination(ec echo.Context) error
	GetTransactionSummary(ec echo.Context) error
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

func (c *controller) GetTransactionSummary(ec echo.Context) error {
	var req structs.GetTransactionSummary

	if err := ec.Bind(&req); err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}

	if req.StartDate != "" && req.EndDate != "" {
		var (
			errValidation = structs.ErrorResponse{
				Code: http.StatusBadRequest,
			}
			formatErrorValidation = "invalid format date for field '%s', must be YYYY-MM-DD"
		)
		startDate, err := time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			return ec.JSON(
				errValidation.Code,
				fmt.Sprintf(formatErrorValidation, "startDate"),
			)
		}
		endDate, err := time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			return ec.JSON(
				errValidation.Code,
				fmt.Sprintf(formatErrorValidation, "endDate"),
			)
		}
		if startDate.After(endDate) {
			return ec.JSON(
				errValidation.Code,
				"endDate must be greated than startDate",
			)
		}
	} else {
		req.StartDate = time.Now().AddDate(0, 0, -7).Format("2006-01-02")
		req.EndDate = time.Now().Format("2006-01-02")
	}

	res, errSvc := c.service.GetTransactionSummary(ec, req)
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
