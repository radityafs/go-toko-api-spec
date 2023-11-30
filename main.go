package main

import (
	"go-toko/database"
	"go-toko/routes"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// INIT DATABASE
	database.Init()
	database.Migrate()

	// INIT FIBER
	app := fiber.New()

	// INIT ROUTES
	routes.Init(app)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}
