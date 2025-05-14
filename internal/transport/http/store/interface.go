package store

import "github.com/gofiber/fiber/v3"

type IStoreAPI interface {
	Healthcheck(c fiber.Ctx) error
	Store(c fiber.Ctx) error
}
