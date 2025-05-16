package account

import (
	"context"
	"github.com/eeQuillibrium/wobble/internal/dto"
	"github.com/eeQuillibrium/wobble/internal/models"
)

type IAccountController interface {
	Register(ctx context.Context, reg dto.Register) (uint64, error)
	Auth(ctx context.Context, log dto.Login) (string, uint64, error)
	GetUserByID(ctx context.Context, userID uint64) (models.User, error)
	GetOrdersByUserID(ctx context.Context, userID uint64) ([]models.Order, error)
	CreateOrder(ctx context.Context, userID uint64, order dto.Order) error
}
