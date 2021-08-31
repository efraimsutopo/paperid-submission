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

type TransactionSumary struct {
	UserID               uint64  `json:"userId" gorm:"column:user_id"`
	StartDate            string  `json:"startDate" gorm:"column:start_date"`
	EndDate              string  `json:"endDate" gorm:"column:end_date"`
	MaxExpenseAmount     uint64  `json:"maxExpenseAmount" gorm:"column:max_expense_amount"`
	AverageExpenseAmount float64 `json:"averageExpenseAmount" gorm:"column:average_expense_amount"`
	TotalTransaction     uint64  `json:"totalTransaction" gorm:"column:total_transaction"`
}

func (Transaction) TableName() string {
	return "transactions"
}
