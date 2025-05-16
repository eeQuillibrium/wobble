package account

import (
	"context"
	"fmt"
	"github.com/eeQuillibrium/wobble/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) UpdateUser(ctx context.Context, userID uint64, upd dto.Update) error {
	user, err := c.r.GetUserByID(ctx, userID)
	if err != nil {
		fmt.Errorf("error with getting password: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(upd.Password)); err != nil {
		return fmt.Errorf("wrong password: %w", err)
	}

	err = c.r.UpdateUser(ctx, userID, upd)
	if err != nil {
		return fmt.Errorf("error with updating user")
	}

	return nil
}
