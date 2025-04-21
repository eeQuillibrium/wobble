-- +goose Up
-- +goose StatementBegin
CREATE TYPE product_category AS ENUM ('красная рыба', 'снеки');

ALTER TABLE product.list ADD COLUMN category product_category;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE product.list DROP COLUMN category;
DROP TYPE product_category;
-- +goose StatementEnd
