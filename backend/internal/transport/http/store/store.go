package store

import "github.com/gofiber/fiber/v3"

func (a API) Store(c fiber.Ctx) error {
	return c.SendString("sdf")
}
