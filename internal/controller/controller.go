package controller

import (
	"github.com/eeQuillibrium/wobble/internal/controller/account"
	"github.com/eeQuillibrium/wobble/internal/controller/contact"
	"github.com/eeQuillibrium/wobble/internal/controller/store"
	"github.com/eeQuillibrium/wobble/internal/repository/psql"
)

type Controller struct {
	Account account.IAccountController
	Store   store.IStoreController
	Contact contact.IContactController
}

func New(r psql.Repository) Controller {
	return Controller{
		Account: account.New(r.Account),
		Store:   store.New(r.Store),
		Contact: contact.New(r.Contact),
	}
}
