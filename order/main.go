package main

import (
	"log"

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

	app.Get("/cart", middleware.Auth, modules.GetCarts)
	app.Get("/cart/:id", middleware.Auth, modules.GetCart)
	app.Post("/cart", middleware.Auth, modules.CreateCart)
	app.Put("/cart/:id", middleware.Auth, modules.UpdateCart)
	app.Delete("/cart/:id", middleware.Auth, modules.DeleteCart)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupModules(app)

	log.Fatal(app.Listen(":3003"))
}