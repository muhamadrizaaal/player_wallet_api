package repository

import (
	"context"
	"player-wallet-api/internal/domain"

	"gorm.io/gorm"
)

type WalletRepository interface {
	Create(ctx context.Context, wallet *domain.Wallet) error
	GetByPlayerID(ctx context.Context, playerID uint) (*domain.Wallet, error)
	UpdateBalance(ctx context.Context, playerID uint, amount float64) error
}

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepository{db}
}

func (r *walletRepository) Create(ctx context.Context, wallet *domain.Wallet) error {
	return r.db.WithContext(ctx).Create(wallet).Error
}

func (r *walletRepository) GetByPlayerID(ctx context.Context, playerID uint) (*domain.Wallet, error) {
	var wallet domain.Wallet
	err := r.db.WithContext(ctx).Where("player_id = ?", playerID).First(&wallet).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *walletRepository) UpdateBalance(ctx context.Context, playerID uint, amount float64) error {
	return r.db.WithContext(ctx).Model(&domain.Wallet{}).
		Where("player_id = ?", playerID).
		UpdateColumn("balance", gorm.Expr("balance + ?", amount)).Error
}
