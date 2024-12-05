package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/theparthshira/go-image-search/elasticsearch"
	"github.com/theparthshira/go-image-search/kafka"
	"github.com/theparthshira/go-image-search/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	routes.InitRoutes(app)

	go func() {
		client, err := elasticsearch.ConnectElasticSearch()

		if err != nil {
			fmt.Println("issue running elasticsearch")
			return
		}

		fmt.Println("Starting Kafka consumer...")
		kafka.ConsumeImageTagData(client)
	}()

	log.Fatal(app.Listen(":4000"))
}
