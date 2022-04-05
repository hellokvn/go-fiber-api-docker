package books

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hellokvn/go-fiber-api-docker/pkg/common/models"
)

func (h handler) GetBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&book)
}
