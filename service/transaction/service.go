package transaction

import (
	"errors"
	"net/http"

	"github.com/efraimsutopo/paperid-submission/helper"
	"github.com/efraimsutopo/paperid-submission/model"
	transactionRepository "github.com/efraimsutopo/paperid-submission/repository/transaction"
	"github.com/efraimsutopo/paperid-submission/structs"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Service interface {
	GetAllInPagination(ec echo.Context, req structs.GetAllInPaginationRequest) (*structs.Pagination, *structs.ErrorResponse)
	GetTransactionSummary(ec echo.Context, req structs.GetTransactionSummary) (*model.TransactionSumary, *structs.ErrorResponse)
	GetTransactionByID(ec echo.Context, transactionID uint64) (*model.Transaction, *structs.ErrorResponse)
	CreateTransaction(ec echo.Context, req structs.CreateTransactionRequest) (*model.Transaction, *structs.ErrorResponse)
	UpdateTransactionByID(ec echo.Context, req structs.UpdateTransactionRequest) (*model.Transaction, *structs.ErrorResponse)
	DeleteTransactionByID(ec echo.Context, transactionID uint64) *structs.ErrorResponse
}

type service struct {
	transactionRepository transactionRepository.Repository
}

func New(
	transactionRepository transactionRepository.Repository,
) Service {
	return &service{
		transactionRepository,
	}
}

func (s *service) GetAllInPagination(
	ec echo.Context,
	req structs.GetAllInPaginationRequest,
) (*structs.Pagination, *structs.ErrorResponse) {
	token, err := helper.GetTokenFromContext(ec)
	if err != nil {
		return nil, s.handleError(err)
	}

	res, err := s.transactionRepository.GetAllInPagination(token.UserID, req)
	if err != nil {
		return nil, s.handleError(err)
	}

	return res, nil
}

func (s *service) GetTransactionSummary(
	ec echo.Context,
	req structs.GetTransactionSummary,
) (*model.TransactionSumary, *structs.ErrorResponse) {
	token, err := helper.GetTokenFromContext(ec)
	if err != nil {
		return nil, s.handleError(err)
	}

	res, err := s.transactionRepository.GetTransactionSummary(token.UserID, req)
	if err != nil {
		s.handleError(err)
	}

	return res, nil
}

func (s *service) GetTransactionByID(ec echo.Context, transactionID uint64) (
	*model.Transaction, *structs.ErrorResponse,
) {
	token, err := helper.GetTokenFromContext(ec)
	if err != nil {
		return nil, s.handleError(err)
	}

	res, err := s.transactionRepository.GetTransactionByID(token.UserID, transactionID)
	if err != nil {
		return nil, s.handleError(err)
	}
	return res, nil
}

func (s *service) CreateTransaction(ec echo.Context, req structs.CreateTransactionRequest) (
	*model.Transaction, *structs.ErrorResponse,
) {
	token, err := helper.GetTokenFromContext(ec)
	if err != nil {
		return nil, s.handleError(err)
	}

	toInsert := model.Transaction{
		UserID: token.UserID,
		Type:   req.Type,
		Amount: req.Amount,
		Note:   req.Note,
	}

	res, err := s.transactionRepository.CreateTransaction(toInsert)
	if err != nil {
		return nil, s.handleError(err)
	}

	return res, nil
}

func (s *service) UpdateTransactionByID(ec echo.Context, req structs.UpdateTransactionRequest) (
	*model.Transaction, *structs.ErrorResponse,
) {
	token, err := helper.GetTokenFromContext(ec)
	if err != nil {
		return nil, s.handleError(err)
	}

	mapUpdates := make(map[string]interface{})

	if req.Type != nil {
		mapUpdates["type"] = req.Type
	}
	if req.Amount != nil {
		mapUpdates["amount"] = req.Amount
	}
	if req.Note != nil {
		mapUpdates["note"] = req.Note
	}

	err = s.transactionRepository.UpdateTransactionByID(token.UserID, req.ID, mapUpdates)
	if err != nil {
		return nil, s.handleError(err)
	}

	res, err := s.transactionRepository.GetTransactionByID(token.UserID, req.ID)
	if err != nil {
		return nil, s.handleError(err)
	}

	return res, nil
}

func (s *service) DeleteTransactionByID(ec echo.Context, transactionID uint64) *structs.ErrorResponse {
	token, err := helper.GetTokenFromContext(ec)
	if err != nil {
		return s.handleError(err)
	}

	err = s.transactionRepository.DeleteTransactionByID(token.UserID, transactionID)
	if err != nil {
		return s.handleError(err)
	}
	return nil
}

func (s *service) handleError(err error) *structs.ErrorResponse {
	if err != nil {
		code := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
		}

		return &structs.ErrorResponse{
			Code:    code,
			Message: err.Error(),
		}
	}

	return nil
}
