-- +goose Up
-- +goose StatementBegin
create schema users;
create schema product;
create schema cart;

create table users.list(
    id serial primary key,
    name varchar not null,
    email varchar,
    login varchar not null,
    password varchar not null
);

create table product.list(
    id serial primary key,
    name varchar not null,
    price numeric,
    img_url varchar not null,
    description varchar,
    amount int
);

create table cart.list(
    id serial primary key,
    user_id int references users.list(id),
    product_id int references product.list(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop schema users cascade;
drop schema product cascade;
drop schema cart cascade;
-- +goose StatementEnd
