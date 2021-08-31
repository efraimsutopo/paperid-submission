package transaction

import (
	"github.com/efraimsutopo/paperid-submission/model"
	"gorm.io/gorm"
)

type Repository interface {
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
