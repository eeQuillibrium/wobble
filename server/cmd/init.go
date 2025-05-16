// Package main - точка входа веб-приложения на Fiber.
// Реализует:
//   - HTTP-сервер с REST API
//   - Подключение к PostgreSQL
//   - Рендеринг HTML через шаблоны
package main

import (
	"context"
	"github.com/eeQuillibrium/wobble/internal/controller"
	"github.com/eeQuillibrium/wobble/internal/repository/psql"
	internal_http "github.com/eeQuillibrium/wobble/internal/transport/http"
	"github.com/eeQuillibrium/wobble/pkg/logger"
	"github.com/eeQuillibrium/wobble/pkg/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"os"
)

// DSN - строка подключения к PostgreSQL.
// Формат: postgres://username:password@host:port/database?params
// В production следует использовать переменные окружения или секреты.

// InitDB создает пул подключений к PostgreSQL и проверяет его работоспособность.
//
// Параметры:
//   - ctx: контекст для управления таймаутами
//
// Возвращает:
//   - *pgxpool.Pool: пул подключений к БД
//   - При ошибке логирует предупреждение и продолжает работу (в продакшене нужно завершать приложение)
//
// Дополнительно:
//   - Применяет миграции через applyMigrations()
//   - Проверяет подключение через Ping()
func InitDB(ctx context.Context) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, utils.GetDSN())
	if err != nil {
		logger.Ctx(ctx).Log(zap.WarnLevel, "pool is down")
	}

	ApplyMigrations(pool)

	if err := pool.Ping(ctx); err != nil {
		logger.Ctx(ctx).Log(zap.WarnLevel, "ping not working")
	}

	return pool
}

// InitServer настраивает все слои приложения (репозиторий, контроллер, API).
//
// Параметры:
//   - pool: пул подключений к БД
//   - app: инстанс Fiber
//
// Возвращает:
//   - Server: готовый HTTP-сервер с инициализированными роутами
//
// Flow:
//
//	Репозиторий -> Контроллер -> API -> Роуты
func InitServer(pool *pgxpool.Pool, app *fiber.App) internal_http.Server {
	repo := psql.New(pool)

	ctrl := controller.New(repo)

	api := internal_http.New(ctrl)

	srv := internal_http.NewServer(api, app)

	srv.InitHttp()

	return srv
}

// NewFiberApp создает инстанс Fiber с настройками для веб-приложения.
//
// Параметры:
//   - engine: шаблонизатор для HTML (используется html/template)
//
// Особенности:
//   - Подключает статические файлы из директории ./frontend
//   - Настраивает маршрутизацию для SPA (если требуется)
func NewFiberApp(engine fiber.Views) *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit:       16 * 1024 * 1024, // 16MB
		ReadBufferSize:  4096,             // 4KB буфер чтения
		WriteBufferSize: 4096,             // 4KB буфер записи
		Views:           engine,
	})

	app.Use(static.New("static/*", static.Config{
		FS: os.DirFS("frontend"),
	}))

	return app
}

// NewFiberTemplates настраивает HTML-шаблонизатор для Fiber.
//
// Возвращает:
//   - fiber.Views: движок для рендеринга HTML с поддержкой горячей перезагрузки в dev-режиме
//
// Настройки:
//   - Delims: использует синтаксис "{{ }}" для совместимости с html/template
//   - Reload: автоматическая перезагрузка шаблонов при изменениях (только для разработки)
func NewFiberTemplates() fiber.Views {
	engine := html.New("./frontend/html", ".html")

	engine.Reload(true)

	engine.Debug(true)

	engine.Delims("{{", "}}")

	return engine
}
