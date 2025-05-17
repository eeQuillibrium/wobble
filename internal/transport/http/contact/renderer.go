package contact

import "github.com/gofiber/fiber/v3"

func (a *API) Index(c fiber.Ctx) error {
	return c.Render("contact", fiber.Map{})
}
