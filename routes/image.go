package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theparthshira/go-image-search/controllers"
)

func InitRoutes(app *fiber.App) error {
	app.Static("images", "./images")

	app.Post("/", controllers.HandleFileupload)

	app.Delete("/:imageName", controllers.HandleDeleteImage)

	return nil
}
