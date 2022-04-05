package products

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := app.Group("/products")
	routes.Post("/", h.AddProduct)
	routes.Get("/", h.GetProducts)
	routes.Get("/:id", h.GetProduct)
	routes.Put("/:id", h.UpdateProduct)
	routes.Delete("/:id", h.DeleteProduct)
}
