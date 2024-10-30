package handler

import (
	"net/http"
	"player-wallet-api/internal/domain"
	"player-wallet-api/internal/usecase"

	"github.com/labstack/echo/v4"
)

type BankHandler struct {
	bankUsecase usecase.BankUsecase
}

func NewBankHandler(bu usecase.BankUsecase) *BankHandler {
	return &BankHandler{bankUsecase: bu}
}

func (h *BankHandler) RegisterBank(c echo.Context) error {
	playerID := c.Get("user_id").(uint)
	req := new(domain.BankRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := h.bankUsecase.Create(c.Request().Context(), playerID, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Bank account registered successfully"})
}

func (h *BankHandler) GetPlayerBanks(c echo.Context) error {
	playerID := c.Get("user_id").(uint)

	banks, err := h.bankUsecase.GetByPlayerID(c.Request().Context(), playerID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, banks)
}
