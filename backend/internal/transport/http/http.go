package http

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"wobble/backend/pkg/logger"
)

const cPort = "8089"

type Server struct {
	app *fiber.App
	api API
}

func NewServer(api API, app *fiber.App) Server {
	return Server{
		app: app,
		api: api,
	}
}

func (s *Server) Run(ctx context.Context) {
	err := s.app.Listen(":" + cPort)
	if err != nil {
		logger.Ctx(ctx).Fatal("failed to Listen port")
	}
}
