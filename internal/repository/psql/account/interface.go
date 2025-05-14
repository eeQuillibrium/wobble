package account

import (
	"context"
	"github.com/eeQuillibrium/wobble/internal/dto"
)

type IAccountRepository interface {
	Register(ctx context.Context, reg dto.Register, passHash []byte) (uint64, error)
	Login(ctx context.Context, login string) ([]byte, uint64, error)
}
