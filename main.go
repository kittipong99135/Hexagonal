package main

import (
	"hex/database"
	"hex/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.RD_Init()
	db := database.DB_Init()

	app := fiber.New()

	routes.Routes(app, db)

	app.Listen("127.0.0.1:3000")
}
