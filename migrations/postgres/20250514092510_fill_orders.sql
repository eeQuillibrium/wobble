-- +goose Up
-- +goose StatementBegin
insert into users.list(
                       name,
                       email,
                       login,
                       password
) VALUES
('Nikita','nikita@mail.ru','login','$2a$10$3PZkZJSCod3ErT6KflvPKutSiTUfFMN6y9WcPb3xlahsKk4607dTy');

insert into users.orders(
    user_id,
    created_at,
    status,
    total
) VALUES
    (1, current_timestamp, 'Активный', 1200, 'г Москва');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
