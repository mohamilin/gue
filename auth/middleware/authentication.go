package middleware

import (
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"github.com/dodysat/gue-auth/database"
	"github.com/dodysat/gue-auth/models"

	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
	var session models.Session

	header := c.Get("Authorization")
	if header == "" {
		return c.Status(401).JSON(fiber.Map{"message": "Missing Authorization Header"})
	}
	token := strings.Split(header, " ")[1]
	if token == "" {
		return c.Status(401).JSON(fiber.Map{"message": "Missing Token"})
	}

	secret := os.Getenv("JWT_SECRET")
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { return []byte(secret), nil })
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid Token"})
	}

	database.Database.Db.Where("access_token = ?", token).First(&session)
	if session.ID == 0 {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid Token"})
	}

	return c.Next()
}

func RefreshToken(c *fiber.Ctx) error {
	var session models.Session

	header := c.Get("Authorization")
	if header == "" {
		return c.Status(401).JSON(fiber.Map{"message": "Missing Authorization Header"})
	}
	token := strings.Split(header, " ")[1]
	if token == "" {
		return c.Status(401).JSON(fiber.Map{"message": "Missing Token"})
	}

	secret := os.Getenv("JWT_SECRET")
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { return []byte(secret), nil })
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid Token"})
	}

	database.Database.Db.Where("refresh_token = ?", token).First(&session)
	if session.ID == 0 {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid Token"})
	}

	return c.Next()
}
