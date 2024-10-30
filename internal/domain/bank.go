package domain

import (
	"time"

	"gorm.io/gorm"
)

type Bank struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	PlayerID      uint           `json:"player_id"`
	BankName      string         `gorm:"not null" json:"bank_name"`
	AccountName   string         `gorm:"not null" json:"account_name"`
	AccountNumber string         `gorm:"not null" json:"account_number"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

type BankRequest struct {
	BankName      string `json:"bank_name" validate:"required"`
	AccountName   string `json:"account_name" validate:"required"`
	AccountNumber string `json:"account_number" validate:"required"`
}
