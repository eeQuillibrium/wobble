package store

import (
	"context"
	"github.com/jackc/pgx/v5"
	"wobble/backend/internal/models"
)

func (r Store) GetProducts(ctx context.Context) ([]models.Product, error) {
	rows, err := r.db.Query(ctx, queryGet)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return nil, err
	}

	return res, nil
}
