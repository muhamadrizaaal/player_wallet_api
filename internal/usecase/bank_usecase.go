package usecase

import (
	"context"
	"player-wallet-api/internal/domain"
	"player-wallet-api/internal/repository"
)

type BankUsecase interface {
	Create(ctx context.Context, playerID uint, req *domain.BankRequest) error
	GetByPlayerID(ctx context.Context, playerID uint) ([]domain.Bank, error)
}

type bankUsecase struct {
	bankRepo repository.BankRepository
}

func NewBankUsecase(br repository.BankRepository) BankUsecase {
	return &bankUsecase{bankRepo: br}
}

func (u *bankUsecase) Create(ctx context.Context, playerID uint, req *domain.BankRequest) error {
	bank := &domain.Bank{
		PlayerID:      playerID,
		BankName:      req.BankName,
		AccountName:   req.AccountName,
		AccountNumber: req.AccountNumber,
	}
	return u.bankRepo.Create(ctx, bank)
}

func (u *bankUsecase) GetByPlayerID(ctx context.Context, playerID uint) ([]domain.Bank, error) {
	return u.bankRepo.GetByPlayerID(ctx, playerID)
}
