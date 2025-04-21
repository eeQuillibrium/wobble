package store

import (
	"context"
	"wobble/backend/internal/models"
)

type IStoreRepository interface {
	GetProducts(ctx context.Context) ([]models.Product, error)
}
