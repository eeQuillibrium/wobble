package cart

import "github.com/gofiber/fiber/v3"

func (a API) Healthcheck(c fiber.Ctx) error {
	return c.SendString("Hello from account cart!")
}
