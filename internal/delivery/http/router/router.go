package router

import (
	"player-wallet-api/internal/delivery/http/handler"
	"player-wallet-api/internal/delivery/http/middleware"
	"player-wallet-api/internal/repository"
	"player-wallet-api/internal/usecase"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB, redisClient *redis.Client) {
	// Initialize repositories
	playerRepo := repository.NewPlayerRepository(db)
	bankRepo := repository.NewBankRepository(db)
	walletRepo := repository.NewWalletRepository(db)

	// Initialize usecases
	playerUsecase := usecase.NewPlayerUsecase(playerRepo, walletRepo, redisClient)
	bankUsecase := usecase.NewBankUsecase(bankRepo)
	walletUsecase := usecase.NewWalletUsecase(walletRepo)

	// Initialize handlers
	playerHandler := handler.NewPlayerHandler(playerUsecase)
	bankHandler := handler.NewBankHandler(bankUsecase)
	walletHandler := handler.NewWalletHandler(walletUsecase)

	// Initialize middleware
	jwtMiddleware := middleware.NewJWTMiddleware(redisClient)

	// API v1 group
	api := e.Group("/api/v1")

	// Public routes
	api.POST("/players/register", playerHandler.Register)
	api.POST("/players/login", playerHandler.Login)

	// Protected routes
	protected := api.Group("")
	protected.Use(jwtMiddleware.ValidateToken)

	// Player routes
	protected.POST("/players/logout", playerHandler.Logout)
	protected.GET("/players", playerHandler.GetAll)
	protected.GET("/players/:id", playerHandler.GetByID)

	// Bank routes
	protected.POST("/players/banks", bankHandler.RegisterBank)
	protected.GET("/players/banks", bankHandler.GetPlayerBanks)

	// Wallet routes
	protected.POST("/players/wallet/topup", walletHandler.TopUp)
	protected.GET("/players/wallet", walletHandler.GetWallet)
}
