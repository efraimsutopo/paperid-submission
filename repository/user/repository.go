package user

import (
	"github.com/efraimsutopo/paperid-submission/model"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(data model.User) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

func (r *repository) CreateUser(data model.User) (*model.User, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) GetUserByEmail(email string) (*model.User, error) {
	var res model.User

	err := r.db.
		Where("email = ?", email).
		First(&res).Error
	if err != nil {
		return nil, err
	}

	return &res, nil
}
