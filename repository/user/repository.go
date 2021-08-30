package user

import (
	"gorm.io/gorm"
)

type Repository interface {
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}
