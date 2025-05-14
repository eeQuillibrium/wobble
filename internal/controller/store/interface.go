package store

import (
	"context"
	"github.com/eeQuillibrium/wobble/internal/models"
)

type IStoreController interface {
	GetProducts(ctx context.Context) ([]models.Product, error)
}
