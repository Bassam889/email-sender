package main

import (
	"os"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/leocm889/server/initializers"
	"github.com/leocm889/server/routes"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("Server is running and processing route: %s\n", c.Path())
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hola AMIGOOOS!")
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	port := os.Getenv("PORT")
	if os.Getenv("PORT") == "" {
		port = ":3000"
	}
	routes.SetupRoutes(app)
	// app.Use(cors.New())
	app.Listen(port)
}
