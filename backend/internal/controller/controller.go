package controller

import (
	"wobble/backend/internal/controller/account"
	"wobble/backend/internal/controller/cart"
	"wobble/backend/internal/controller/store"
	"wobble/backend/internal/repository/psql"
)

type Controller struct {
	Account account.IAccountController
	Cart    cart.ICartController
	Store   store.IStoreController
}

func New(r psql.Repository) Controller {
	return Controller{
		Account: account.New(r.Account),
		Cart:    cart.New(r.Cart),
		Store:   store.New(r.Store),
	}
}
