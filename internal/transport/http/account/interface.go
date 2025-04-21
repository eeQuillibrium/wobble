package account

import "github.com/gofiber/fiber/v3"

type IAccountAPI interface {
	Healthcheck(c fiber.Ctx) error
}
