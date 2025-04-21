package account

import "wobble/backend/internal/controller/account"

type API struct {
	ctrl account.IAccountController
}

func New(ctrl account.IAccountController) API {
	return API{
		ctrl: ctrl,
	}
}
