-- +goose Up
-- +goose StatementBegin
create table users.orders(
    id bigserial primary key,
    user_id int references users.list(id),
    created_at timestamptz default current_timestamp not null,
    status varchar,
    total int
);

create table users.orders_products(
    id bigserial primary key,
    order_id bigint references users.orders(id),
    product_id int references product.list(id),
    quantity int not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
