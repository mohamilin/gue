package modules

import (
	"strconv"

	"github.com/dodysat/gue-product/database"
	"github.com/dodysat/gue-product/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint   `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Stock       uint   `json:"stock"`
}

func CreateResponseProduct(product models.Product) Product {
	return Product{
		ID:          product.ID,
		UserID:      product.UserID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	userId := c.Locals("userId")

	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error parsing product"})
	}

	product.UserID = uint(userId.(float64))

	database.Database.Db.Create(&product)
	return c.Status(201).JSON(CreateResponseProduct(product))
}

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	userId := c.Locals("userId")
	database.Database.Db.Find(&products, "user_id = ?", userId)
	if len(products) == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No products found"})
	}
	return c.Status(200).JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	var product models.Product
	id := c.Params("id")
	userId := c.Locals("userId")
	database.Database.Db.Find(&product, "id = ? AND user_id = ?", id, userId)
	if product.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No product found with given id"})
	}
	return c.Status(200).JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	var product models.Product
	id := c.Params("id")
	userId := c.Locals("userId")
	database.Database.Db.Find(&product, "id = ? AND user_id = ?", id, userId)
	if product.Name == "" {
		return c.Status(404).JSON(fiber.Map{"message": "No product found with given id"})
	}
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error parsing product"})
	}
	database.Database.Db.Model(&product).Updates(product)
	return c.Status(200).JSON(CreateResponseProduct(product))
}

func DeleteProduct(c *fiber.Ctx) error {
	var product models.Product
	id := c.Params("id")
	userId := c.Locals("userId")
	database.Database.Db.Find(&product, "id = ? AND user_id = ?", id, userId)
	if product.Name == "" {
		return c.Status(404).JSON(fiber.Map{"message": "No product found with given id"})
	}
	database.Database.Db.Delete(&product)
	return c.Status(200).JSON(fiber.Map{"message": "Product deleted"})
}

func ModifyStock(c *fiber.Ctx) error {
	var product models.Product
	id := c.Params("id")
	amount := c.Query("amount")
	userId := c.Locals("userId")
	database.Database.Db.Find(&product, "id = ? AND user_id = ?", id, userId)
	if product.Name == "" {
		return c.Status(404).JSON(fiber.Map{"message": "No product found with given id"})
	}

	a, _ := strconv.Atoi(amount)
	modifiedStock := product.Stock + uint(a)
	if modifiedStock < 0 {
		return c.Status(400).JSON(fiber.Map{"message": "Stock cannot be negative"})
	}

	database.Database.Db.Model(&product).Update("stock", modifiedStock)
	return c.Status(200).JSON(CreateResponseProduct(product))
}
