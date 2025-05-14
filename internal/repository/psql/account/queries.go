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
