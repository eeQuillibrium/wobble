package controller

import (
	"github.com/eeQuillibrium/wobble/internal/controller/account"
	"github.com/eeQuillibrium/wobble/internal/controller/cart"
	"github.com/eeQuillibrium/wobble/internal/controller/store"
	"github.com/eeQuillibrium/wobble/internal/repository/psql"
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
