package account

import repository "wobble/backend/internal/repository/psql/account"

type Controller struct {
	r repository.IAccountRepository
}

func New(r repository.IAccountRepository) Controller {
	return Controller{
		r: r,
	}
}
