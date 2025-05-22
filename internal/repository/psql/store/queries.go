package store

const queryGet = `
	SELECT id, name, price, img_url, description, category FROM product.list
`

const queryCreateProduct = `
	INSERT INTO product.list(name, price, img_url, description, category, amount)
	VALUES ($1, $2, $3, $4, $5, $6)
`

const queryGetByID = `
	SELECT id, name, price, img_url, description, category
	FROM product.list
	WHERE id = $1
`
