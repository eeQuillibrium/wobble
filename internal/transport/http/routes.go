package http

import (
	"github.com/eeQuillibrium/wobble/internal/transport/http/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

// InitHttp регистрирует все HTTP-роуты приложения.
//
// Структура API:
// GET  /healthcheck       - общая проверка работы сервиса
// GET  /                  - главная страница
//
// Группа account:
// GET  /account/healthcheck - проверка подсистемы аккаунтов
//
// Группа cart:
// GET  /cart/healthcheck    - проверка подсистемы корзины
//
// Группа store:
// GET  /store/healthcheck   - проверка подсистемы магазина
// GET  /store/              - основной интерфейс магазина
//
// Middleware:
// Все middleware должны быть зарегистрированы до вызова этого метода.
// Рекомендуемые middleware:
//   - Логирование
//   - Recovery
//   - CORS
//   - Rate Limiting
func (s *Server) InitHttp() {
	s.app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	s.app.Get("/healthcheck", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	s.app.Get("/", func(c fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	// Account
	account := s.app.Group("/account")

	account.Get("/", s.api.account.IndexT, middleware.JWTAuthMiddleware())
	account.Get("/auth", s.api.account.AuthT)
	account.Get("/registration", s.api.account.RegisterT)
	account.Get("/changeuser", s.api.account.ChangeUserInfoT)

	account.Post("/register", s.api.account.Register)
	account.Post("/login", s.api.account.Auth)
	account.Post("/logout", s.api.account.Logout)
	account.Get("/GetUser", s.api.account.GetUser, middleware.JWTAuthMiddleware())
	account.Get("/GetOrders", s.api.account.GetOrders, middleware.JWTAuthMiddleware())
	account.Post("/CreateOrder", s.api.account.CreateOrder, middleware.JWTAuthMiddleware())
	account.Post("/UpdateUser", s.api.account.UpdateUser, middleware.JWTAuthMiddleware())

	// Blog
	blog := s.app.Group("/blog")

	blog.Get("/", func(c fiber.Ctx) error {
		return c.Render("blog", fiber.Map{})
	})

	// Contact
	contact := s.app.Group("/contact")

	contact.Get("/", s.api.contact.Index)
	contact.Post("/CreateAppeal", s.api.contact.CreateAppeal)

	// Store
	store := s.app.Group("/store")

	store.Get("/", s.api.store.Store)
	store.Post("/AddProduct", s.api.store.AddProduct)
	store.Get("/:id", s.api.store.GetProduct)
}
