package main

import (
	"log"

	"github.com/dodysat/gue-product/database"
	"github.com/dodysat/gue-product/modules"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the product service")
}

func setupModules(app *fiber.App) {
	app.Get("/", welcome)

	app.Post("/product", modules.CreateProduct)
	app.Get("/product", modules.GetProducts)
	app.Get("/product/:id", modules.GetProduct)
	app.Put("/product/:id", modules.UpdateProduct)
	app.Delete("/product/:id", modules.DeleteProduct)
	app.Put("/product/:id/modifystock", modules.ModifyStock)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupModules(app)

	log.Fatal(app.Listen(":3002"))
}
