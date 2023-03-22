package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Body struct {
	Service string      `json:"service"`
	Path    string      `json:"path"`
	Method  string      `json:"method"`
	Body    interface{} `json:"body"`
}

func Auth(c *fiber.Ctx) error {
	// urlAuth := "http://localhost:3001/verify"
	urlAuth := "http://auth:3001/verify"

	token := c.Get("Authorization")

	body := &Body{
		Service: "order",
		Path:    c.Path(),
		Method:  c.Method(),
		Body:    c.Body(),
	}

	payloadBuff := new(bytes.Buffer)
	json.NewEncoder(payloadBuff).Encode(body)

	req, err := http.NewRequest("POST", urlAuth, payloadBuff)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Internal Server Error"})
	}

	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Internal Server Error"})
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	userId := result["userId"].(float64)

	c.Locals("userId", userId)

	return c.Next()
}
