package middleware

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/dodysat/gue-auth/database"
	"github.com/dodysat/gue-auth/models"

	"github.com/gofiber/fiber/v2"
)

func Activity(c *fiber.Ctx) error {
	header := c.Get("Authorization")

	token := strings.Split(header, " ")[1]
	secret := "secret"
	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { return []byte(secret), nil })
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid Token"})
	}

	claims := decoded.Claims.(jwt.MapClaims)
	userID := claims["userId"]

	database.Database.Db.Create(&models.Activity{
		UserID:    uint(userID.(float64)),
		Service:   "auth",
		Path:      c.Path(),
		Method:    c.Method(),
		Timestamp: time.Now(),
	})

	return c.Next()
}
