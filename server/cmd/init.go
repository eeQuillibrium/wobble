package main

import (
	"context"
	"github.com/eeQuillibrium/wobble/internal/controller"
	"github.com/eeQuillibrium/wobble/internal/repository/psql"
	internal_http "github.com/eeQuillibrium/wobble/internal/transport/http"
	"github.com/eeQuillibrium/wobble/pkg/logger"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"os"
)

const cDSN = "postgres://postgres:secret@localhost:5433/wobble?sslmode=disable"

func initDB(ctx context.Context) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, cDSN)
	if err != nil {
		logger.Ctx(ctx).Log(zap.WarnLevel, "pool is down")
	}

	applyMigrations(pool)

	if err := pool.Ping(ctx); err != nil {
		logger.Ctx(ctx).Log(zap.WarnLevel, "ping not working")
	}

	return pool
}

func initServer(pool *pgxpool.Pool, app *fiber.App) internal_http.Server {
	repo := psql.New(pool)

	ctrl := controller.New(repo)

	api := internal_http.New(ctrl)

	srv := internal_http.NewServer(api, app)

	srv.InitHttp()

	return srv
}

func NewFiberApp(engine fiber.Views) *fiber.App {
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(static.New("static/*", static.Config{
		FS: os.DirFS("frontend"),
	}))

	return app
}

func NewFiberTemplates() fiber.Views {
	engine := html.New("./frontend/html", ".html")

	engine.Reload(true)

	engine.Debug(true)

	engine.Delims("{{", "}}")

	return engine
}
