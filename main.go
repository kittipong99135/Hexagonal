package main

import (
	"hex/database"
	"hex/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.RD_Init()
	db := database.DB_Init()

	app := fiber.New()
	app.Use(cors.New())

	routes.Routes(app, db)

	app.Listen("127.0.0.1:3000")
}
