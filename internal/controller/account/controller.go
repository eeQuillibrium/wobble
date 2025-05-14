package account

import repository "github.com/eeQuillibrium/wobble/internal/repository/psql/account"

type Controller struct {
	r repository.IAccountRepository
}

func New(r repository.IAccountRepository) *Controller {
	return &Controller{
		r: r,
	}
}
