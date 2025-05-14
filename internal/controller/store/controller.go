package store

import repository "github.com/eeQuillibrium/wobble/internal/repository/psql/store"

type Controller struct {
	r repository.IStoreRepository
}

func New(r repository.IStoreRepository) Controller {
	return Controller{
		r: r,
	}
}
