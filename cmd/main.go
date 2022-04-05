package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hellokvn/go-fiber-api-docker/pkg/books"
	"github.com/hellokvn/go-fiber-api-docker/pkg/common/config"
	"github.com/hellokvn/go-fiber-api-docker/pkg/common/db"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	books.RegisterRoutes(app, h)

	app.Listen(c.Port)
}
