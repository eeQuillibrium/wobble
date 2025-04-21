package cart

import repository "wobble/backend/internal/repository/psql/cart"

type Controller struct {
	r repository.ICartRepository
}

func New(r repository.ICartRepository) Controller {
	return Controller{
		r: r,
	}
}
