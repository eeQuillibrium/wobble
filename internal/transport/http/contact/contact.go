package contact

import (
	"github.com/eeQuillibrium/wobble/internal/controller/contact"
)

type API struct {
	ctrl contact.IContactController
}

func New(ctrl contact.IContactController) *API {
	return &API{
		ctrl: ctrl,
	}
}
