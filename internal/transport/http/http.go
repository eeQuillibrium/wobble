package http

import (
	"context"
	"github.com/eeQuillibrium/wobble/pkg/logger"
	"github.com/gofiber/fiber/v3"
)

const cPort = "3000"

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
