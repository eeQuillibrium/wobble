-- +goose Up
-- +goose StatementBegin
insert into product.list(name, price, img_url, description, amount, category)
VALUES (
           'Семга',
           1000,
           'img/store/products/salmon.png',
           'Жирнейшая красная рыба',
           800,
           'красная рыба'
       ),
       ('Вобла',
        100,
        'img/store/products/vobla.png',
        'Лучший снек. Подойдет для всего.',
        500,
        'снеки');


-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
