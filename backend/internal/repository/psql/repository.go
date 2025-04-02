package psql

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"wobble/backend/internal/repository/psql/account"
	"wobble/backend/internal/repository/psql/cart"
	"wobble/backend/internal/repository/psql/store"
)

type Repository struct {
	Account account.IAccountRepository
	Cart    cart.ICartRepository
	Store   store.IStoreRepository
}

func New(db *pgxpool.Pool) Repository {
	return Repository{
		Account: account.New(db),
		Cart:    cart.New(db),
		Store:   store.New(db),
	}
}
