package contact

import (
	"context"
	"fmt"
	"github.com/eeQuillibrium/wobble/internal/dto"
)

func (c *Contact) CreateAppeal(ctx context.Context, appeal dto.Appeal) error {
	_, err := c.db.Exec(ctx, queryCreate, appeal.Name, appeal.Email, appeal.Phone, appeal.Subject, appeal.Message)
	if err != nil {
		return fmt.Errorf("error with creating appeal: %w", err)
	}

	return nil
}
