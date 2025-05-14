package index

import "github.com/gofiber/fiber/v3"

type IIndexAPI interface {
	Index(c fiber.Ctx) error
}
