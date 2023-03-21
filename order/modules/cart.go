package modules

import (
	"github.com/dodysat/gue-order/database"
	"github.com/dodysat/gue-order/models"
	"github.com/gofiber/fiber/v2"
)

type Cart struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint   `json:"user_id"`
	ProductID   uint   `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       uint   `json:"price"`
	Amount      uint   `json:"amount"`
}

func CreateResponseCart(cart models.Cart) Cart {
	return Cart{
		ID:          cart.ID,
		UserID:      cart.UserID,
		ProductID:   cart.ProductID,
		ProductName: cart.ProductName,
		Price:       cart.Price,
		Amount:      cart.Amount,
	}
}

func CreateCart(c *fiber.Ctx) error {
	var cart models.Cart
	userId := c.Locals("userId")
	err := c.BodyParser(&cart)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error parsing cart"})
	}

	cart.UserID = uint(userId.(float64))

	// if product_id exist in cart, add amount
	var productInCart models.Cart
	database.Database.Db.Find(&productInCart, "user_id = ? AND product_id = ?", userId, cart.ProductID)
	if productInCart.ID != 0 {
		productInCart.Amount += cart.Amount
		database.Database.Db.Save(&productInCart)
		return c.Status(201).JSON(CreateResponseCart(productInCart))
	}

	database.Database.Db.Create(&cart)
	return c.Status(201).JSON(CreateResponseCart(cart))
}

func GetCarts(c *fiber.Ctx) error {
	var cart []models.Cart
	userId := c.Locals("userId")
	database.Database.Db.Find(&cart, "user_id = ?", userId)
	if len(cart) == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No product found"})
	}

	return c.Status(200).JSON(cart)
}

func GetCart(c *fiber.Ctx) error {
	var cart models.Cart
	id := c.Params("id")
	userId := c.Locals("userId")
	database.Database.Db.Find(&cart, "user_id = ? AND id = ?", userId, id)
	if cart.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No cart found with given id"})
	}
	return c.Status(200).JSON(cart)
}

func UpdateCart(c *fiber.Ctx) error {
	var cart models.Cart
	id := c.Params("id")
	userId := c.Locals("userId")
	database.Database.Db.Find(&cart, "user_id = ? AND id = ?", userId, id)
	if cart.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No cart found with given id"})
	}

	err := c.BodyParser(&cart)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error parsing cart"})
	}
	database.Database.Db.Model(&cart).Updates(cart)
	return c.Status(200).JSON(CreateResponseCart(cart))
}

func DeleteCart(c *fiber.Ctx) error {
	var cart models.Cart
	id := c.Params("id")
	userId := c.Locals("userId")
	database.Database.Db.Find(&cart, "user_id = ? AND id = ?", userId, id)
	if cart.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No cart found with given id"})
	}

	database.Database.Db.Delete(&cart)
	return c.Status(200).JSON(fiber.Map{
		"message": "Cart deleted",
	})
}
