package cart

import "wobble/backend/internal/controller/cart"

type API struct {
	ctrl cart.ICartController
}

func New(ctrl cart.ICartController) API {
	return API{
		ctrl: ctrl,
	}
}
