-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA appeals;

CREATE TABLE appeals.list(
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
DROP SCHEMA appeals;
-- +goose StatementEnd
