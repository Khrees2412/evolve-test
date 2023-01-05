package router

import (
	"evolve/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Get("/users", controller.GetUsers)
	route.Get("/users/:email", controller.GetUserByEmail)
}
