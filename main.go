package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/theparthshira/go-image-search/elasticsearch"
	"github.com/theparthshira/go-image-search/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	routes.InitRoutes(app)

	client, _ := elasticsearch.ConnectElasticSearch()

	data := elasticsearch.PhotoTag{
		Tag: "123",
		Id:  "999",
	}

	elasticsearch.IndexElasticData(client, data)

	fmt.Println("test", elasticsearch.QueryElasticData(client, "123"))

	log.Fatal(app.Listen(":4000"))
}
