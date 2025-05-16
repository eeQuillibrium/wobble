package account

import (
	"github.com/eeQuillibrium/wobble/internal/dto"
	"github.com/gofiber/fiber/v3"
)

func (a *API) UpdateUser(c fiber.Ctx) error {
	userID := c.Locals("userID").(uint64)
	var upd dto.Update

	if err := c.Bind().JSON(&upd); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Не удалось расшифровать json",
		})
	}

	err := a.ctrl.UpdateUser(c.Context(), userID, upd)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Неуспешное обновление пользовательских данных",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"error":   nil,
	})
}
