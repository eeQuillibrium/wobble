package store

import (
	"context"
	"github.com/eeQuillibrium/wobble/internal/models"
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
