package store

import (
	"github.com/eeQuillibrium/wobble/internal/dto"
	"github.com/gofiber/fiber/v3"
)

func (a *API) AddProduct(c fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Не удалось извлечь файл",
		})
	}

	var product dto.Product
	if err := c.Bind().Form(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Не удалось извлечь файл",
		})
	}

	err = a.ctrl.AddProduct(c.Context(), product, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Не удалось добавить продукт",
		})
	}

	return nil
}
