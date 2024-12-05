package controllers

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/theparthshira/go-image-search/elasticsearch"
	"github.com/theparthshira/go-image-search/kafka"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": 200, "message": "Server is running 🚀🚀"})
}

func HandleFileupload(c *fiber.Ctx) error {
	// fmt.Println("uploading...")

	// file, err := c.FormFile("image")

	// if err != nil {
	// 	log.Println("image upload error --> ", err)
	// 	return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	// }

	// token, err := utils.GetB2AuthToken()

	// if err != nil {
	// 	log.Println("image upload error --> ", err)
	// 	return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	// }

	// fmt.Println("token: ", token)

	// uploadURL, err := utils.GetB2UploadURL(token)

	// if err != nil {
	// 	log.Println("image upload error --> ", err)
	// 	return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	// }

	// fmt.Println("uploadURL: ", uploadURL)

	// fileURL, err := utils.UploadB2File(file, uploadURL)

	// if err != nil {
	// 	log.Println("image upload error --> ", err)
	// 	return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	// }

	// fmt.Println("fileURL: ", fileURL)

	// kafka.PushCommentToQueue("comments", []byte(fileURL))

	// return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": fileURL})

	file, err := c.FormFile("image")

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	uniqueId := uuid.New()

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	fileExt := strings.Split(file.Filename, ".")[1]

	image := fmt.Sprintf("%s.%s", filename, fileExt)

	err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))

	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	imageUrl := fmt.Sprintf("http://localhost:4000/images/%s", image)

	data := map[string]interface{}{

		"imageName": image,
		"imageUrl":  imageUrl,
		"header":    file.Header,
		"size":      file.Size,
	}

	kafka.PushCommentToQueue("comments", []byte(image))

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})

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

func GetSearchQuery(c *fiber.Ctx) error {
	queries := c.Queries()

	client, _ := elasticsearch.ConnectElasticSearch()
	searchResult := elasticsearch.QueryElasticData(client, queries["search"])

	searchUrl := []string{}

	for _, id := range searchResult {
		searchUrl = append(searchUrl, fmt.Sprintf("http://localhost:4000/images/%s", id.Id))
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Images listed successfully", "data": searchUrl})
}
