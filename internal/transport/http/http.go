// Package http реализует HTTP-сервер на базе Fiber.
// Предоставляет:
//   - Конфигурацию сервера
//   - Запуск и graceful shutdown
//   - Интеграцию с API-обработчиками
package http

import (
	"context"
	"github.com/eeQuillibrium/wobble/pkg/logger"
	"github.com/gofiber/fiber/v3"
)

// Port - стандартный порт для работы сервера.
// Может быть переопределен через переменные окружения или конфигурацию.
const Port = "3000"

// Server представляет HTTP-сервер приложения.
// Содержит:
//   - app: инстанс Fiber для маршрутизации и middleware
//   - api: набор API-обработчиков (REST/gRPC)
//
// Сервер автоматически обрабатывает graceful shutdown при отмене контекста.
type Server struct {
	app *fiber.App
	api API
}

// NewServer создает новый HTTP-сервер.
//
// Параметры:
//   - api: инициализированные обработчики API (см. package api)
//   - app: предварительно сконфигурированный инстанс Fiber (с middleware)
//
// Пример:
//
//	fiberApp := fiber.New()
//	api := http.NewAPI(controllers)
//	srv := http.NewServer(api, fiberApp)
func NewServer(api API, app *fiber.App) Server {
	return Server{
		app: app,
		api: api,
	}
}

// Run запускает сервер с поддержкой graceful shutdown.
// Блокирующая операция, должна быть главной горутиной.
func (s *Server) Run(ctx context.Context) {
	err := s.app.Listen(":" + Port)
	if err != nil {
		logger.Ctx(ctx).Fatal("failed to Listen port")
	}
}
