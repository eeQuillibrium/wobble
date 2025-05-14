package cart

import "github.com/gofiber/fiber/v3"

type ICartAPI interface {
	Healthcheck(c fiber.Ctx) error
}
