package repository

import (
	"context"
	"player-wallet-api/internal/domain"

	"gorm.io/gorm"
)

type PlayerRepository interface {
	Create(ctx context.Context, player *domain.Player) error
	GetByID(ctx context.Context, id uint) (*domain.Player, error)
	GetByUsername(ctx context.Context, username string) (*domain.Player, error)
	GetAll(ctx context.Context, filters map[string]interface{}) ([]domain.Player, error)
}

type playerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) PlayerRepository {
	return &playerRepository{db}
}

func (r *playerRepository) Create(ctx context.Context, player *domain.Player) error {
	return r.db.WithContext(ctx).Create(player).Error
}

func (r *playerRepository) GetByID(ctx context.Context, id uint) (*domain.Player, error) {
	var player domain.Player
	err := r.db.WithContext(ctx).Preload("Banks").Preload("Wallet").First(&player, id).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func (r *playerRepository) GetByUsername(ctx context.Context, username string) (*domain.Player, error) {
	var player domain.Player
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&player).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func (r *playerRepository) GetAll(ctx context.Context, filters map[string]interface{}) ([]domain.Player, error) {
	var players []domain.Player
	query := r.db.WithContext(ctx).Preload("Banks").Preload("Wallet")

	for key, value := range filters {
		switch key {
		case "username":
			query = query.Where("username LIKE ?", "%"+value.(string)+"%")
		case "bank_name":
			query = query.Joins("LEFT JOIN banks ON banks.player_id = players.id").
				Where("banks.bank_name LIKE ?", "%"+value.(string)+"%")
		case "min_balance":
			query = query.Joins("LEFT JOIN wallets ON wallets.player_id = players.id").
				Where("wallets.balance >= ?", value)
		}
	}

	err := query.Find(&players).Error
	return players, err
}
