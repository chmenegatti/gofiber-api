package main

import (
	"log"

	"github.com/chmenegatti/gofiber-api/database"
	"github.com/chmenegatti/gofiber-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
