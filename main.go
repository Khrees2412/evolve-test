package main

import (
	"evolve/database"
	"evolve/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}
	log.Println("Environment variables successfully loaded. Starting application...")
}

func main() {
	app := fiber.New()

	//Connect Database
	database.Connect()

	//Setup routes
	router.Setup(app)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Evolve Credit Test")
	})
	//Activate CORS
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// 404 Handler
	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(404) // => 404 "Not Found"
	})
	// Get the PORT from hosting service env
	port := os.Getenv("PORT")

	// Verify if hosting service provided the port or not
	if port == "" {
		port = "8008"
	}

	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}

}
