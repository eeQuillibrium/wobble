package account

import (
	"context"
	"fmt"
	"github.com/eeQuillibrium/wobble/internal/dto"
	"github.com/jackc/pgx/v5"
)

func (a *Account) CreateOrder(ctx context.Context, userID uint64, order dto.Order) error {
	var orderID uint64

	err := a.WrapWithTransaction(ctx, func(tx pgx.Tx) error {
		err := tx.QueryRow(ctx, queryCreateOrder, userID, order.TotalAmount, order.DeliveryAddress).Scan(&orderID)
		if err != nil {
			return fmt.Errorf("error with create order: %w", err)
		}

		query, args, err := a.buildInsertOrderProduct(orderID, order.Items)

		_, err = tx.Exec(ctx, query, args...)
		if err != nil {
			return fmt.Errorf("error with create order_products: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error with transaction: %w", err)
	}

	return nil
}

func (a *Account) buildInsertOrderProduct(orderID uint64, items []dto.OrderedProduct) (string, []interface{}, error) {
	in := a.qb.Insert("users.orders_products").
		Columns("order_id", "product_id", "quantity")

	for _, item := range items {
		in = in.Values(orderID, item.ID, item.Quantity)
	}

	return in.ToSql()
}
