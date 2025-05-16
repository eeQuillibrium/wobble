package contact

import (
	"context"
	"github.com/eeQuillibrium/wobble/internal/dto"
)

func (c *Controller) CreateAppeal(ctx context.Context, appeal dto.Appeal) error {
	return c.r.CreateAppeal(ctx, appeal)
}
