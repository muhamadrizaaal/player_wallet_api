package usecase

import (
	"context"
	"player-wallet-api/internal/domain"
	"player-wallet-api/internal/repository"
)

type WalletUsecase interface {
	TopUp(ctx context.Context, playerID uint, amount float64) error
	GetByPlayerID(ctx context.Context, playerID uint) (*domain.Wallet, error)
}

type walletUsecase struct {
	walletRepo repository.WalletRepository
}

func NewWalletUsecase(wr repository.WalletRepository) WalletUsecase {
	return &walletUsecase{walletRepo: wr}
}

func (u *walletUsecase) TopUp(ctx context.Context, playerID uint, amount float64) error {
	return u.walletRepo.UpdateBalance(ctx, playerID, amount)
}

func (u *walletUsecase) GetByPlayerID(ctx context.Context, playerID uint) (*domain.Wallet, error) {
	return u.walletRepo.GetByPlayerID(ctx, playerID)
}
