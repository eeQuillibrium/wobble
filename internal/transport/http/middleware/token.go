package middleware

import (
	"fmt"
	"github.com/eeQuillibrium/wobble/pkg/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func JWTAuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		tokenString := extractToken(c)
		if tokenString == "" {
			return c.Redirect().To("/account/auth")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(utils.Secret), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Invalid token",
				"details": err.Error(),
			})
		}

		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token is invalid",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token claims",
			})
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User ID not found in token",
			})
		}

		role, ok := claims["role"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Role not found in token",
			})
		}

		c.Locals("userID", uint64(userID))
		c.Locals("role", role)

		return c.Next()
	}
}

func extractToken(c fiber.Ctx) string {
	authHeader := c.Get("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}

	if cookie := c.Cookies("jwt"); cookie != "" {
		return cookie
	}

	if queryToken := c.Query("token"); queryToken != "" {
		return queryToken
	}

	return ""
}
