package model

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	ID        uint64         `json:"-" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	UserID    uint64         `json:"-"`
	Token     string         `json:"-"`
}
