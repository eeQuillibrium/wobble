package store

import (
	"context"
	"wobble/backend/internal/models"
)

type IStoreController interface {
	GetProducts(ctx context.Context) ([]models.Product, error)
}
