package contact

import (
	"github.com/eeQuillibrium/wobble/internal/dto"
	"github.com/gofiber/fiber/v3"
)

func (a *API) CreateAppeal(c fiber.Ctx) error {
	var appeal dto.Appeal

	if err := c.Bind().JSON(&appeal); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Неверный формат данных",
		})
	}

	err := a.ctrl.CreateAppeal(c.Context(), appeal)

	return err
}
