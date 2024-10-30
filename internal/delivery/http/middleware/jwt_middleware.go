package middleware

import (
	"net/http"
	"player-wallet-api/pkg/utils"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

type JWTMiddleware struct {
	redisClient *redis.Client
}

func NewJWTMiddleware(rc *redis.Client) *JWTMiddleware {
	return &JWTMiddleware{redisClient: rc}
}

func (m *JWTMiddleware) ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing authorization header"})
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Validate token format and signature
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}

		// Check if token is in Redis (not logged out)
		_, err = m.redisClient.Get(c.Request().Context(), tokenString).Result()
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token expired or invalid"})
		}

		// Set user ID in context
		c.Set("user_id", uint(claims.UserID))
		return next(c)
	}
}
