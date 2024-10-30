package handler

import (
	"net/http"
	"player-wallet-api/internal/domain"
	"player-wallet-api/internal/usecase"

	"github.com/labstack/echo/v4"
)

type WalletHandler struct {
	walletUsecase usecase.WalletUsecase
}

func NewWalletHandler(wu usecase.WalletUsecase) *WalletHandler {
	return &WalletHandler{walletUsecase: wu}
}

func (h *WalletHandler) TopUp(c echo.Context) error {
	playerID := c.Get("user_id").(uint)
	req := new(domain.TopUpRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := h.walletUsecase.TopUp(c.Request().Context(), playerID, req.Amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Top up successful"})
}

func (h *WalletHandler) GetWallet(c echo.Context) error {
	playerID := c.Get("user_id").(uint)

	wallet, err := h.walletUsecase.GetByPlayerID(c.Request().Context(), playerID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, wallet)
}
