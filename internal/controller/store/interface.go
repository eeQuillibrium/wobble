package store

import (
	"context"
	"github.com/eeQuillibrium/wobble/internal/dto"
	"github.com/eeQuillibrium/wobble/internal/models"
	"mime/multipart"
)

type IStoreController interface {
	GetProducts(ctx context.Context) ([]models.Product, error)
	GetProduct(ctx context.Context, id uint64) (models.Product, error)
	AddProduct(ctx context.Context, product dto.Product, file *multipart.FileHeader) error
}
