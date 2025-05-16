package account

import (
	"context"
	"fmt"
	"github.com/eeQuillibrium/wobble/internal/dto"
)

func (a *Account) UpdateUser(ctx context.Context, userID uint64, upd dto.Update) error {
	_, err := a.db.Exec(ctx, queryUpdateUser, userID, upd.Name, upd.Email, upd.Login)
	if err != nil {
		return fmt.Errorf("error quering update user: %w", err)
	}

	return nil
}
