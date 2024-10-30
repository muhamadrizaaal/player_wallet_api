package usecase

import (
	"context"
	"errors"
	"player-wallet-api/internal/domain"
	"player-wallet-api/internal/repository"
	"player-wallet-api/pkg/utils"

	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
)

type PlayerUsecase interface {
	Register(ctx context.Context, req *domain.RegisterRequest) error
	Login(ctx context.Context, req *domain.LoginRequest) (string, error)
	Logout(ctx context.Context, token string) error
	GetByID(ctx context.Context, id uint) (*domain.Player, error)
	GetAll(ctx context.Context, filters map[string]interface{}) ([]domain.Player, error)
}

type playerUsecase struct {
	playerRepo  repository.PlayerRepository
	walletRepo  repository.WalletRepository
	redisClient *redis.Client
}

func NewPlayerUsecase(pr repository.PlayerRepository, wr repository.WalletRepository, rc *redis.Client) PlayerUsecase {
	return &playerUsecase{
		playerRepo:  pr,
		walletRepo:  wr,
		redisClient: rc,
	}
}

func (u *playerUsecase) Register(ctx context.Context, req *domain.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	player := &domain.Player{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	err = u.playerRepo.Create(ctx, player)
	if err != nil {
		return err
	}

	// Create wallet for new player
	wallet := &domain.Wallet{
		PlayerID: player.ID,
		Balance:  0,
	}
	return u.walletRepo.Create(ctx, wallet)
}

func (u *playerUsecase) Login(ctx context.Context, req *domain.LoginRequest) (string, error) {
	player, err := u.playerRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(player.ID)
	if err != nil {
		return "", err
	}

	// Store token in Redis
	err = u.redisClient.Set(ctx, token, player.ID, 0).Err()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *playerUsecase) Logout(ctx context.Context, token string) error {
	return u.redisClient.Del(ctx, token).Err()
}

func (u *playerUsecase) GetByID(ctx context.Context, id uint) (*domain.Player, error) {
	return u.playerRepo.GetByID(ctx, id)
}

func (u *playerUsecase) GetAll(ctx context.Context, filters map[string]interface{}) ([]domain.Player, error) {
	return u.playerRepo.GetAll(ctx, filters)
}
