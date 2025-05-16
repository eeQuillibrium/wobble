package account

import (
	"context"
	"github.com/eeQuillibrium/wobble/internal/dto"
)

func (c *Controller) CreateOrder(ctx context.Context, userID uint64, order dto.Order) error {
	return c.r.CreateOrder(ctx, userID, order)
}
