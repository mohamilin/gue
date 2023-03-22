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
	product.Get("/", middleware.Auth, modules.GetProducts)
	product.Post("/", middleware.Auth, modules.CreateProduct)
	product.Get("/:id", middleware.Auth, modules.GetProduct)
	product.Put("/:id", middleware.Auth, modules.UpdateProduct)
	product.Delete("/:id", middleware.Auth, modules.DeleteProduct)
	product.Put("/:id/modifystock", middleware.Auth, modules.ModifyStock)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupModules(app)

	port := os.Getenv("PORT")

	log.Fatal(app.Listen(":" + port))
}
