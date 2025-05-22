package store

import "github.com/gofiber/fiber/v3"

type IStoreAPI interface {
	Store(c fiber.Ctx) error
	AddProduct(c fiber.Ctx) error
	GetProduct(c fiber.Ctx) error
}
