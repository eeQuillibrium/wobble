package contact

const queryCreate = `
	INSERT INTO appeals.list(name, email, phone, subject, message) 
	VALUES ($1, $2, $3, $4, $5)
`
