package repository

import (
	"context"
	"player-wallet-api/internal/domain"

	"gorm.io/gorm"
)

type BankRepository interface {
	Create(ctx context.Context, bank *domain.Bank) error
	GetByPlayerID(ctx context.Context, playerID uint) ([]domain.Bank, error)
}

type bankRepository struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB) BankRepository {
	return &bankRepository{db}
}

func (r *bankRepository) Create(ctx context.Context, bank *domain.Bank) error {
	return r.db.WithContext(ctx).Create(bank).Error
}

func (r *bankRepository) GetByPlayerID(ctx context.Context, playerID uint) ([]domain.Bank, error) {
	var banks []domain.Bank
	err := r.db.WithContext(ctx).Where("player_id = ?", playerID).Find(&banks).Error
	return banks, err
}
