-- +goose Up
-- +goose StatementBegin
ALTER TABLE users.orders
    ADD COLUMN delivery_address varchar;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
