package main

import (
	"log"

	"github.com/dodysat/gue-auth/database"
	"github.com/dodysat/gue-auth/middleware"
	"github.com/dodysat/gue-auth/modules"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the auth service")
}

func setupModules(app *fiber.App) {
	app.Get("/", welcome)

	app.Post("/register", modules.UserRegister)
	app.Post("/login", modules.UserLogin)
	app.Get("/logout", middleware.Authentication, middleware.Activity, modules.UserLogout)
	app.Get("/refresh", middleware.RefreshToken, middleware.Activity, modules.UserRefresh)

	app.Post("/verify", middleware.Authentication, modules.UserVerify)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupModules(app)

	log.Fatal(app.Listen(":3001"))
}
