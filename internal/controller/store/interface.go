package store

import (
	"context"
	"github.com/eeQuillibrium/wobble/internal/dto"
	"github.com/eeQuillibrium/wobble/internal/models"
	"mime/multipart"
)

type IStoreController interface {
	GetProducts(ctx context.Context) ([]models.Product, error)
	AddProduct(ctx context.Context, product dto.Product, file *multipart.FileHeader) error
}
