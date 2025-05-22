package store

import (
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func (a *API) GetProduct(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}

	product, err := a.ctrl.GetProduct(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}

	return c.Render("product", fiber.Map{
		"Product": product,
	})
}
