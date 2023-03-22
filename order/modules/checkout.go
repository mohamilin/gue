package modules

import (
	"fmt"
	"net/http"

	"github.com/dodysat/gue-order/database"
	"github.com/dodysat/gue-order/models"
	"github.com/gofiber/fiber/v2"
)

type Checkout struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint   `json:"user_id"`
	ProductID   uint   `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       uint   `json:"price"`
	TotalPrice  uint   `json:"total_price"`
	Amount      uint   `json:"amount"`
}

type CartPayload struct {
	CartId uint `json:"cart_id"`
}

func CreateResponseCheckout(checkout models.Checkout) Checkout {
	return Checkout{
		ID:          checkout.ID,
		UserID:      checkout.UserID,
		ProductID:   checkout.ProductID,
		ProductName: checkout.ProductName,
		Price:       checkout.Price,
		TotalPrice:  checkout.TotalPrice,
		Amount:      checkout.Amount,
	}
}

func GetCheckouts(c *fiber.Ctx) error {
	var checkouts []models.Checkout
	userId := c.Locals("userId")

	database.Database.Db.Find(&checkouts, "user_id = ?", userId)

	return c.Status(200).JSON(checkouts)
}

func GetCheckout(c *fiber.Ctx) error {
	var checkout models.Checkout
	userId := c.Locals("userId")
	checkoutId := c.Params("id")

	database.Database.Db.Find(&checkout, "id = ? AND user_id = ?", checkoutId, userId)
	if checkout.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Checkout not found"})
	}

	return c.Status(200).JSON(checkout)
}

func CreateCheckout(c *fiber.Ctx) error {
	var checkout models.Checkout
	var cart models.Cart
	userId := c.Locals("userId")

	cartPayload := new(CartPayload)

	if err := c.BodyParser(cartPayload); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error parsing cart"})
	}

	cartId := cartPayload.CartId

	database.Database.Db.Find(&cart, "id = ? AND user_id = ?", cartId, userId)
	if cart.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Cart not found"})
	}

	checkout.UserID = uint(userId.(float64))
	checkout.ProductID = cart.ProductID
	checkout.ProductName = cart.ProductName
	checkout.Price = cart.Price
	checkout.Amount = cart.Amount
	checkout.TotalPrice = cart.Price * cart.Amount

	productId := cart.ProductID
	amount := -int(cart.Amount)
	urlProductService := fmt.Sprintf("http://product:3002/product/%d/modifystock?amount=%d", productId, amount)
	token := c.Get("Authorization")

	req, err := http.NewRequest("PUT", urlProductService, nil)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error update stock 1"})
	}

	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error update stock 2"})
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return c.Status(500).JSON(fiber.Map{"message": "Error update stock 3"})
	}

	database.Database.Db.Create(&checkout)
	database.Database.Db.Delete(&cart)

	return c.Status(201).JSON(CreateResponseCheckout(checkout))
}

func DeleteCheckout(c *fiber.Ctx) error {
	var checkout models.Checkout
	userId := c.Locals("userId")
	checkoutId := c.Params("id")

	database.Database.Db.Find(&checkout, "id = ? AND user_id = ?", checkoutId, userId)
	if checkout.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "Checkout not found"})
	}

	productId := checkout.ProductID
	amount := int(checkout.Amount)
	urlProductService := fmt.Sprintf("http://product:3002/product/%d/modifystock?amount=%d", productId, amount)
	token := c.Get("Authorization")

	req, err := http.NewRequest("PUT", urlProductService, nil)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error update stock 1"})
	}

	req.Header.Set("Authorization", token)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error update stock 2"})
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return c.Status(500).JSON(fiber.Map{"message": "Error update stock 3"})
	}

	database.Database.Db.Delete(&checkout)

	return c.Status(200).JSON(fiber.Map{"message": "Checkout deleted"})
}
