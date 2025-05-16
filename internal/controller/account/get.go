package account

import (
	"context"
	"github.com/eeQuillibrium/wobble/internal/models"
)

func (c *Controller) GetUserByID(ctx context.Context, userID uint64) (models.User, error) {
	return c.r.GetUserByID(ctx, userID)
}

func (c *Controller) GetOrdersByUserID(ctx context.Context, userID uint64) ([]models.Order, error) {
	return c.r.GetOrdersByUserID(ctx, userID)
}
