package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        uint64         `json:"id" gorm:"primarykey"`
	UserID    uint64         `json:"userId" gorm:"index;not null"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Type      string         `json:"type"`
	Amount    uint64         `json:"amount"`
	Note      string         `json:"note"`
}
