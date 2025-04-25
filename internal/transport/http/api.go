package http

import (
	"github.com/eeQuillibrium/wobble/internal/controller"
	"github.com/eeQuillibrium/wobble/internal/transport/http/account"
	"github.com/eeQuillibrium/wobble/internal/transport/http/cart"
	"github.com/eeQuillibrium/wobble/internal/transport/http/index"
	"github.com/eeQuillibrium/wobble/internal/transport/http/store"
)

type API struct {
	account account.IAccountAPI
	cart    cart.ICartAPI
	store   store.IStoreAPI
	index   index.IIndexAPI
}

func New(ctrl controller.Controller) API {
	return API{
		account: account.New(ctrl.Account),
		cart:    cart.New(ctrl.Cart),
		store:   store.New(ctrl.Store),
		index:   index.New(),
	}
}
