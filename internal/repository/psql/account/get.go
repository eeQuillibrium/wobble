package account

import (
	"context"
	"fmt"
	"github.com/eeQuillibrium/wobble/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

func (a *Account) GetUserByID(ctx context.Context, userID uint64) (models.User, error) {
	var user models.User

	err := pgxscan.Get(ctx, a.db, &user, queryGetUserByID, userID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (a *Account) GetOrdersByUserID(ctx context.Context, userID uint64) ([]models.Order, error) {
	rows, err := a.db.Query(ctx, queryGetOrders, userID)
	if err != nil {
		return nil, err
	}

	results, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Order])
	if err != nil {
		return nil, fmt.Errorf("errorCollectRows for GetOrders: %w", err)
	}

	return results, nil
}
