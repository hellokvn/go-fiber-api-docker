package main

import (
	"log"

	"go-fiber-api-docker/pkg/common/config"
	"go-fiber-api-docker/pkg/common/db"
	"go-fiber-api-docker/pkg/products"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)
	app := fiber.New()

	products.RegisterRoutes(app, h)

	app.Listen(c.Port)
}
