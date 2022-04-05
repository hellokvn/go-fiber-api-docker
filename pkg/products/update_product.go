package products

import (
	"go-fiber-api-docker/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

type UpdateProductRequestBody struct {
	Name  string `json:"name"`
	Stock int32  `json:"stock"`
	Price int32  `json:"price"`
}

func (h handler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateProductRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var product models.Product

	if result := h.DB.First(&product, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	product.Name = body.Name
	product.Stock = body.Stock
	product.Price = body.Price

	// save product
	h.DB.Save(&product)

	return c.Status(fiber.StatusOK).JSON(&product)
}
