-- +goose Up
-- +goose StatementBegin
insert into users.orders_products(
    order_id,
    product_id,
    quantity
) VALUES
      (1, 1, 1),
      (1, 2, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
