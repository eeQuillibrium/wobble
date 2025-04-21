package index

import "github.com/gofiber/fiber/v3"

func (i API) Index(c fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}
