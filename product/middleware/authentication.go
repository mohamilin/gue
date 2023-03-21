package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	urlAuth := "http://localhost:3001/user/verify"

	token := c.Get("Authorization")

	req, err := http.NewRequest("GET", urlAuth, nil)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}
