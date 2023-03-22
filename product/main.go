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

	app.Get("/product", middleware.Auth, modules.GetProducts)
	app.Post("/product", middleware.Auth, modules.CreateProduct)
	app.Get("/product/:id", middleware.Auth, modules.GetProduct)
	app.Put("/product/:id", middleware.Auth, modules.UpdateProduct)
	app.Delete("/product/:id", middleware.Auth, modules.DeleteProduct)
	app.Put("/product/:id/modifystock", middleware.Auth, modules.ModifyStock)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupModules(app)

	port := os.Getenv("PORT")

	log.Fatal(app.Listen(":" + port))
}
