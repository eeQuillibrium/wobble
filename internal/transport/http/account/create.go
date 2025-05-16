package account

import (
	"github.com/eeQuillibrium/wobble/internal/dto"
	"github.com/gofiber/fiber/v3"
)

func (a *API) CreateOrder(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint64)
	var order dto.Order

	if err := c.Bind().JSON(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Неверный формат данных заказа",
		})
	}

	err := a.ctrl.CreateOrder(c.Context(), userID, order)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Не удалось создать заказ: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Заказ успешно создан",
	})
}
