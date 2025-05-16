package account

const queryCreate = `
	INSERT INTO 
		users.list(name, email, login, password) 
	VALUES($1, $2, $3, $4)
	RETURNING id`

const queryGetUser = `
	SELECT id, password FROM users.list
    WHERE login = $1
`

const queryGetUserByID = `
	SELECT id, name, email, login FROM users.list
	WHERE id = $1
`

const queryGetOrders = `
SELECT
    o.id,
    o.created_at,
    o.status,
    o.total AS total_amount,
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
INSERT INTO users.orders(user_id, status, total) 
VALUES ($1, 'Активный', $2)
RETURNING id
`

const queryCreateOrderProduct = `
INSERT INTO users.orders_products(order_id, product_id, quantity)
VALUES ($1, $2, $3)
`
