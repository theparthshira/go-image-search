package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/theparthshira/go-image-search/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	routes.InitRoutes(app)

	log.Fatal(app.Listen(":4000"))
}
