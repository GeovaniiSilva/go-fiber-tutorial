package main

import (
	"github.com/GeovaniiSilva/go-fiber-tutorial/user"
	"github.com/gofiber/fiber/v2"
)

func Routers(app *fiber.App) {
	app.Get("/users", user.GetUsers)
	app.Get("/user/:id", user.GetUser)
	app.Post("/users", user.SaveUser)
	app.Delete("/user/:id", user.DeleteUser)
	app.Put("/user/:id", user.UpdateUser)
}

func main() {
	user.InitialMigration()
	app := fiber.New()
	Routers(app)
	app.Listen(":3000")
}
