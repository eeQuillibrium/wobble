package http

import "github.com/gofiber/fiber/v3"

func (s *Server) InitHttp() {
	s.app.Get("/healthcheck", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Index
	s.app.Get("/", s.api.index.Index)

	// Account
	account := s.app.Group("/account")
	account.Get("healthcheck", s.api.account.Healthcheck)

	// Cart
	cart := s.app.Group("/cart")
	cart.Get("healthcheck", s.api.cart.Healthcheck)

	// Store
	store := s.app.Group("/store")
	store.Get("healthcheck", s.api.store.Healthcheck)
	store.Get("/", s.api.store.Store)
}
