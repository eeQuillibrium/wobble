package store

import (
	"context"
	"github.com/eeQuillibrium/wobble/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

func (s Store) GetProducts(ctx context.Context) ([]models.Product, error) {
	rows, err := s.db.Query(ctx, queryGet)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s Store) GetProduct(ctx context.Context, id uint64) (models.Product, error) {
	var product models.Product
	if err := pgxscan.Get(ctx, s.db, &product, queryGetByID, id); err != nil {
		return models.Product{}, err
	}

	return product, nil
}
