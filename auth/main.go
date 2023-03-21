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

	app.Post("/user/register", modules.UserRegister)
	app.Post("/user/login", modules.UserLogin)
	app.Get("/user/logout", middleware.Authentication, middleware.Activity, modules.UserLogout)
	app.Get("/user/refresh", middleware.RefreshToken, middleware.Activity, modules.UserRefresh)

	app.Get("/user/verify", middleware.Authentication, welcome)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupModules(app)

	log.Fatal(app.Listen(":3001"))
}
