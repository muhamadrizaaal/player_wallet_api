package domain

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	PlayerID  uint           `json:"player_id"`
	Balance   float64        `gorm:"not null;default:0" json:"balance"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type TopUpRequest struct {
	Amount float64 `json:"amount" validate:"required,gt=0"`
}
