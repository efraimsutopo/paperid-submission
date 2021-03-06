package session

import (
	"github.com/efraimsutopo/paperid-submission/model"
	"gorm.io/gorm"
)

type Repository interface {
	CreateSession(data model.Session) (*model.Session, error)
	GetSessionByToken(tokenString string) (*model.Session, error)
	DeleteSessionByToken(tokenString string) error
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

func (r *repository) CreateSession(data model.Session) (*model.Session, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repository) GetSessionByToken(tokenString string) (*model.Session, error) {
	var res model.Session
	err := r.db.
		Where("token = ?", tokenString).
		First(&res).Error
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *repository) DeleteSessionByToken(tokenString string) error {
	err := r.db.
		Where("token = ?", tokenString).
		Delete(&model.Session{}).Error
	if err != nil {
		return err
	}

	return nil
}
