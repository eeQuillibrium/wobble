package account

import (
	"context"
	"github.com/eeQuillibrium/wobble/internal/dto"
)

type IAccountController interface {
	Register(ctx context.Context, reg dto.Register) (uint64, error)
	Auth(ctx context.Context, log dto.Login) (string, uint64, error)
}
