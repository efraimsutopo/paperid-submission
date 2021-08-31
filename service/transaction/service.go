package transaction

import (
	"errors"
	"net/http"

	"github.com/efraimsutopo/paperid-submission/model"
	transactionRepository "github.com/efraimsutopo/paperid-submission/repository/transaction"
	"github.com/efraimsutopo/paperid-submission/structs"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Service interface {
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

func (s *service) GetTransactionByID(ec echo.Context, transactionID uint64) (
	*model.Transaction, *structs.ErrorResponse,
) {
	res, err := s.transactionRepository.GetTransactionByID(1, transactionID) // TODO: Change Real User ID from token
	if err != nil {
		return nil, s.handleError(err)
	}
	return res, nil
}

func (s *service) CreateTransaction(ec echo.Context, req structs.CreateTransactionRequest) (
	*model.Transaction, *structs.ErrorResponse,
) {
	var toInsert = model.Transaction{
		UserID: 1, // TODO: Change Real User ID from token
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

	err := s.transactionRepository.UpdateTransactionByID(1, req.ID, mapUpdates) // TODO: Change Real User ID from token
	if err != nil {
		return nil, s.handleError(err)
	}

	res, err := s.transactionRepository.GetTransactionByID(1, req.ID) // TODO: Change Real User ID from token
	if err != nil {
		return nil, s.handleError(err)
	}

	return res, nil
}

func (s *service) DeleteTransactionByID(ec echo.Context, transactionID uint64) *structs.ErrorResponse {
	err := s.transactionRepository.DeleteTransactionByID(1, transactionID) // TODO: Change Real User ID from token
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
