package store

import (
	"context"
	"github.com/eeQuillibrium/wobble/internal/dto"
)

func (s Store) CreateProduct(ctx context.Context, product dto.Product) error {
	_, err := s.db.Exec(ctx, queryCreateProduct, product.Name, product.Price, product.ImageURL, product.Description, product.Category, product.Amount)
	if err != nil {
		return err
	}

	return nil
}
