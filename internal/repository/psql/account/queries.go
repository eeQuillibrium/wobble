package account

const queryCreate = `
	INSERT INTO 
		users.list(name, email, login, password) 
	VALUES($1, $2, $3, $4)
	RETURNING id, role`

const queryGetUser = `
	SELECT id, password, role FROM users.list
    WHERE login = $1
`

const queryGetUserByID = `
	SELECT id, name, email, login, password FROM users.list
	WHERE id = $1
`

const queryGetOrders = `
SELECT
    o.id,
    o.created_at,
    o.status,
    o.total AS total_amount,
    o.delivery_address as delivery_address,
    jsonb_agg(
            jsonb_build_object(
                    'id', p.id,
                    'name', p.name,
                    'description', p.description,
                    'price', p.price,
                    'imageUrl', p.img_url,
                    'category', p.category,
                    'quantity', op.quantity
            )
    ) AS items
FROM
    users.orders o
        JOIN
    users.orders_products op ON o.id = op.order_id
        JOIN
    product.list p ON op.product_id = p.id
WHERE
    o.user_id = $1
GROUP BY
    o.id
ORDER BY
    o.created_at DESC;`

const queryCreateOrder = `
INSERT INTO users.orders(user_id, status, total, delivery_address) 
VALUES ($1, 'Активный', $2, $3)
RETURNING id
`

const queryUpdateUser = `
UPDATE users.list
SET name = $2, email = $3, login = $4
WHERE id = $1
`
