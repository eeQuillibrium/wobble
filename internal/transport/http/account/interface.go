package account

import "github.com/gofiber/fiber/v3"

type IAccountAPI interface {
	AuthT(c fiber.Ctx) error
	IndexT(c fiber.Ctx) error
	RegisterT(c fiber.Ctx) error

	Auth(c fiber.Ctx) error
	Register(c fiber.Ctx) error
	GetUser(c fiber.Ctx) error
	GetOrders(c fiber.Ctx) error
	CreateOrder(c fiber.Ctx) error
}
