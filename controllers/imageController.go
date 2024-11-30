package controllers

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/theparthshira/go-image-search/utils"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": 200, "message": "Server is running 🚀🚀"})
}

func HandleFileupload(c *fiber.Ctx) error {
	fmt.Println("uploading...")

	file, err := c.FormFile("image")

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	token, err := utils.GetB2AuthToken()

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	fmt.Println("token: ", token)

	uploadURL, err := utils.GetB2UploadURL(token)

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	fmt.Println("uploadURL: ", uploadURL)

	fileURL, err := utils.UploadB2File(file, uploadURL)

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	fmt.Println("fileURL: ", fileURL)

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": fileURL})
}

func HandleDeleteImage(c *fiber.Ctx) error {
	imageName := c.Params("imageName")

	err := os.Remove(fmt.Sprintf("./images/%s", imageName))
	if err != nil {
		log.Println(err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server Error", "data": nil})
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image deleted successfully", "data": nil})
}

func GetImagesList(c *fiber.Ctx) error {
	path := "D:\\live\\go-image-search\\images"
	dir, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var files []string

	for _, file := range dir {
		files = append(files, file.Name())
	}

	fmt.Println(files)

	return c.JSON(fiber.Map{"status": 201, "message": "Images listed successfully", "data": files})
}
