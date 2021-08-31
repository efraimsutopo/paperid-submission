package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Email     string         `json:"email" gorm:"type:varchar(255);uniqueIndex"`
	Password  string         `json:"-" gorm:"type:varchar(255)"`
	Name      string         `json:"name" gorm:"type:varchar(255)"`
}

func (User) TableName() string {
	return "users"
}
