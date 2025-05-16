package contact

import "github.com/gofiber/fiber/v3"

type IContactAPI interface {
	CreateAppeal(c fiber.Ctx) error
}
