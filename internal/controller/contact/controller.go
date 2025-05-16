package contact

import repository "github.com/eeQuillibrium/wobble/internal/repository/psql/contact"

type Controller struct {
	r repository.IContactRepository
}

func New(r repository.IContactRepository) *Controller {
	return &Controller{
		r: r,
	}
}
