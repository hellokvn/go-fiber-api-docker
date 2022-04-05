package books

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

	routes := app.Group("/books")
	routes.Post("/", h.AddBook)
	routes.Get("/", h.GetBooks)
	routes.Get("/:id", h.GetBook)
	routes.Put("/:id", h.UpdateBook)
	routes.Delete("/:id", h.DeleteBook)
}
