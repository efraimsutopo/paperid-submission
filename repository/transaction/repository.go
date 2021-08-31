package transaction

import (
	"math"

	"github.com/efraimsutopo/paperid-submission/model"
	"github.com/efraimsutopo/paperid-submission/structs"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllInPagination(userID uint64, req structs.GetAllInPaginationRequest) (*structs.Pagination, error)
	GetTransactionByID(userID, transactionID uint64) (*model.Transaction, error)
	CreateTransaction(data model.Transaction) (*model.Transaction, error)
	UpdateTransactionByID(userID, transactionID uint64, updates map[string]interface{}) error
	DeleteTransactionByID(userID, transactionID uint64) error
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

func (r *repository) GetAllInPagination(userID uint64, req structs.GetAllInPaginationRequest) (*structs.Pagination, error) {
	var (
		transactions []*model.Transaction
		pagination   = structs.Pagination{
			Limit: req.Limit,
			Page:  req.Page,
			Sort:  req.Sort,
		}
	)

	query := r.db.Where("user_id = ?", userID)
	if req.Type != nil {
		query = query.Where("type = ?", req.Type)
	}
	if req.Amount != nil {
		query = query.Where("amount = ?", req.Amount)
	}
	if req.Note != nil {
		query = query.Where("note = ?", req.Note)
	}

	err := query.
		Scopes(paginate(transactions, &pagination, query)).
		Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	pagination.Rows = transactions

	return &pagination, nil
}

func (r *repository) GetTransactionByID(userID, transactionID uint64) (*model.Transaction, error) {
	var res model.Transaction

	err := r.db.
		Where("id = ?", transactionID).
		Where("user_id = ?", userID).
		First(&res).Error
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *repository) CreateTransaction(data model.Transaction) (*model.Transaction, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) UpdateTransactionByID(
	userID, transactionID uint64,
	updates map[string]interface{},
) error {
	err := r.db.Model(&model.Transaction{}).
		Where("user_id = ?", userID).
		Where("id = ?", transactionID).
		Updates(updates).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteTransactionByID(userID, transactionID uint64) error {
	var data model.Transaction

	err := r.db.
		Where("id = ?", transactionID).
		Where("user_id = ?", userID).
		First(&data).Error
	if err != nil {
		return err
	}

	err = r.db.Delete(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func paginate(value interface{}, pagination *structs.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
