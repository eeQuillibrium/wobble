package account

import (
	"context"
	"github.com/eeQuillibrium/wobble/internal/dto"
	"github.com/eeQuillibrium/wobble/internal/models"
)

type IAccountRepository interface {
	Register(ctx context.Context, reg dto.Register, passHash []byte) (models.RegisteredUser, error)
	GetLoginInfo(ctx context.Context, login string) (models.RegisteredUser, error)
	GetUserByID(ctx context.Context, userID uint64) (models.User, error)
	GetOrdersByUserID(ctx context.Context, userID uint64) ([]models.Order, error)
	CreateOrder(ctx context.Context, userID uint64, order dto.Order) error
	UpdateUser(ctx context.Context, userID uint64, upd dto.Update) error
}
