package account

import "github.com/gofiber/fiber/v3"

func (a *API) AuthT(c fiber.Ctx) error {
	return c.Render("auth", fiber.Map{})
}

func (a *API) RegisterT(c fiber.Ctx) error {
	return c.Render("registration", fiber.Map{})
}

func (a *API) IndexT(c fiber.Ctx) error {
	return c.Render("account", fiber.Map{})
}

func (a *API) ChangeUserInfoT(c fiber.Ctx) error {
	return c.Render("changeuser", fiber.Map{})
}
