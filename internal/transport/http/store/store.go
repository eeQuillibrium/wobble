package store

import (
	"github.com/eeQuillibrium/wobble/pkg/logger"
	"github.com/gofiber/fiber/v3"
)

func (a *API) Store(c fiber.Ctx) error {
	products, err := a.ctrl.GetProducts(c.Context())
	if err != nil {
		logger.Ctx(c.Context()).Warn("error w/ products getting")
	}

	return c.Render("store", fiber.Map{
		"Products": products,
	})
}
