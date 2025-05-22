package store

import (
	"context"
	"github.com/eeQuillibrium/wobble/internal/models"
	"github.com/eeQuillibrium/wobble/pkg/logger"
)

func (c Controller) GetProducts(ctx context.Context) ([]models.Product, error) {
	logger.Ctx(ctx).Info("GetProducts")

	return c.r.GetProducts(ctx)
}

func (c Controller) GetProduct(ctx context.Context, id uint64) (models.Product, error) {
	logger.Ctx(ctx).Info("GetProducts")

	return c.r.GetProduct(ctx, id)
}
