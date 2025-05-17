package store

import (
	"context"
	"github.com/eeQuillibrium/wobble/internal/dto"
	"github.com/eeQuillibrium/wobble/internal/models"
)

type IStoreRepository interface {
	GetProducts(ctx context.Context) ([]models.Product, error)
	CreateProduct(ctx context.Context, product dto.Product) error
}
