package router

import (
	"evolve/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/users", controller.GetUsers)
	app.Get("/users/:email", controller.GetUserByEmail)
}
