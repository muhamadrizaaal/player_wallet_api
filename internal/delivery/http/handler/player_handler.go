package handler

import (
	"net/http"
	"player-wallet-api/internal/domain"
	"player-wallet-api/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PlayerHandler struct {
	playerUsecase usecase.PlayerUsecase
}

func NewPlayerHandler(pu usecase.PlayerUsecase) *PlayerHandler {
	return &PlayerHandler{playerUsecase: pu}
}

func (h *PlayerHandler) Register(c echo.Context) error {
	req := new(domain.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := h.playerUsecase.Register(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Registration successful"})
}

func (h *PlayerHandler) Login(c echo.Context) error {
	req := new(domain.LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	token, err := h.playerUsecase.Login(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token":   token,
		"message": "Login successful",
	})
}

func (h *PlayerHandler) Logout(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing authorization token"})
	}

	// Remove "Bearer " prefix
	token = token[7:]

	err := h.playerUsecase.Logout(c.Request().Context(), token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Logout successful"})
}

func (h *PlayerHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid player ID"})
	}

	player, err := h.playerUsecase.GetByID(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Player not found"})
	}

	return c.JSON(http.StatusOK, player)
}

func (h *PlayerHandler) GetAll(c echo.Context) error {
	// Prepare filters from query parameters
	filters := make(map[string]interface{})

	// Username filter
	if username := c.QueryParam("username"); username != "" {
		filters["username"] = username
	}

	// Bank name filter
	if bankName := c.QueryParam("bank_name"); bankName != "" {
		filters["bank_name"] = bankName
	}

	// Account name filter
	if accountName := c.QueryParam("account_name"); accountName != "" {
		filters["account_name"] = accountName
	}

	// Account number filter
	if accountNumber := c.QueryParam("account_number"); accountNumber != "" {
		filters["account_number"] = accountNumber
	}

	// Minimum balance filter
	if minBalance := c.QueryParam("min_balance"); minBalance != "" {
		balance, err := strconv.ParseFloat(minBalance, 64)
		if err == nil {
			filters["min_balance"] = balance
		}
	}

	// Date range filter for registration
	if registerAt := c.QueryParam("register_at"); registerAt != "" {
		filters["register_at"] = registerAt
	}

	players, err := h.playerUsecase.GetAll(c.Request().Context(), filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"total":   len(players),
		"players": players,
	})
}
