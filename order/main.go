package main

import (
	"log"
	"os"

	"github.com/dodysat/gue-order/database"
	"github.com/dodysat/gue-order/middleware"
	"github.com/dodysat/gue-order/modules"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the order service")
}

func setupModules(app *fiber.App) {
	app.Get("/", welcome)

	cart := app.Group("/cart", middleware.Auth)
	cart.Get("/", modules.GetCarts)
	cart.Get("/:id", modules.GetCart)
	cart.Post("/", modules.CreateCart)
	cart.Put("/:id", modules.UpdateCart)
	cart.Delete("/:id", modules.DeleteCart)

	checkout := app.Group("/checkout", middleware.Auth)
	checkout.Get("/", modules.GetCheckouts)
	checkout.Get("/:id", modules.GetCheckout)
	checkout.Post("/", modules.CreateCheckout)
	checkout.Delete("/:id", modules.DeleteCheckout)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupModules(app)

	port := os.Getenv("PORT")

	log.Fatal(app.Listen(":" + port))
}
