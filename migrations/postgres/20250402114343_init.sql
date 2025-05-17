-- +goose Up
-- +goose StatementBegin
create schema users;
create schema product;
create schema cart;
create schema reference;
create schema appeals;

create type reference.user_role as enum ('admin', 'user');

create table users.list(
    id serial primary key,
    name varchar not null,
    email varchar,
    login varchar not null unique,
    password varchar not null,
    role reference.user_role default 'user' not null
);

create table product.list(
    id serial primary key,
    name varchar not null,
    price numeric,
    img_url varchar not null,
    description varchar,
    category varchar,
    amount int
);

create table cart.list(
    id serial primary key,
    user_id int references users.list(id),
    product_id int references product.list(id)
);

create table users.orders(
    id bigserial primary key,
    user_id int references users.list(id),
    created_at timestamptz default current_timestamp not null,
    status varchar,
    delivery_address varchar,
    total int
);

create table users.orders_products(
    id bigserial primary key,
    order_id bigint references users.orders(id),
    product_id int references product.list(id),
    quantity int not null
);

create table appeals.list(
    id bigserial primary key,
    name varchar not null,
    email varchar not null,
    phone varchar not null,
    subject varchar not null,
    message varchar not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop schema users cascade;
drop schema product cascade;
drop schema cart cascade;
drop schema appeals cascade;
-- +goose StatementEnd
