package cart

import "github.com/eeQuillibrium/wobble/internal/controller/cart"

type API struct {
	ctrl cart.ICartController
}

func New(ctrl cart.ICartController) API {
	return API{
		ctrl: ctrl,
	}
}
