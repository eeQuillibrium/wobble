package store

import (
	"context"
	"wobble/backend/internal/models"
	"wobble/backend/pkg/logger"
)

func (c Controller) GetProducts(ctx context.Context) ([]models.Product, error) {
	logger.Ctx(ctx).Info("GetProducts")

	return c.r.GetProducts(ctx)
}
