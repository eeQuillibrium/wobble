package store

import "wobble/backend/internal/controller/store"

type API struct {
	ctrl store.IStoreController
}

func New(ctrl store.IStoreController) API {
	return API{
		ctrl: ctrl,
	}
}
