package http

import (
	"wobble/backend/internal/controller"
	"wobble/backend/internal/transport/http/account"
	"wobble/backend/internal/transport/http/cart"
	"wobble/backend/internal/transport/http/store"
)

type API struct {
	account account.IAccountAPI
	cart    cart.ICartAPI
	store   store.IStoreAPI
}

func New(ctrl controller.Controller) API {
	return API{
		account: account.New(ctrl.Account),
		cart:    cart.New(ctrl.Cart),
		store:   store.New(ctrl.Store),
	}
}
