package main

import (
	"log"
	"os"

	"github.com/dodysat/gue-product/database"
	"github.com/dodysat/gue-product/middleware"
	"github.com/dodysat/gue-product/modules"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the product service")
}

func setupModules(app *fiber.App) {
	app.Get("/", welcome)

	product := app.Group("/product", middleware.Auth)
	product.Get("/", modules.GetProducts)
	product.Post("/", modules.CreateProduct)
	product.Get("/:id", modules.GetProduct)
	product.Put("/:id", modules.UpdateProduct)
	product.Delete("/:id", modules.DeleteProduct)
	product.Put("/:id/modifystock", modules.ModifyStock)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupModules(app)

	port := os.Getenv("PORT")

	log.Fatal(app.Listen(":" + port))
}
