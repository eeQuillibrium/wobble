package account

import (
	"context"
	"fmt"
	"github.com/eeQuillibrium/wobble/internal/dto"
	"github.com/eeQuillibrium/wobble/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) Register(ctx context.Context, reg dto.Register) (uint64, error) {
	passHash, err := utils.HashPass(reg.Password)
	if err != nil {
		return 0, err
	}
	return c.r.Register(ctx, reg, passHash)
}

func (c *Controller) Auth(ctx context.Context, log dto.Login) (string, uint64, error) {
	dbPass, userID, err := c.r.Login(ctx, log.Login)
	if err != nil {
		return "", 0, fmt.Errorf("error w/ getting password: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword(dbPass, []byte(log.Password)); err != nil {
		return "", 0, err
	}

	token, err := utils.GenerateJWT(userID)
	if err != nil {
		return "", 0, err
	}

	return token, userID, nil
}
