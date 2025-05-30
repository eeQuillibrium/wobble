package psql

import (
	"github.com/eeQuillibrium/wobble/internal/repository/psql/account"
	"github.com/eeQuillibrium/wobble/internal/repository/psql/contact"
	"github.com/eeQuillibrium/wobble/internal/repository/psql/store"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	Account account.IAccountRepository
	Store   store.IStoreRepository
	Contact contact.IContactRepository
}

func New(db *pgxpool.Pool) Repository {
	return Repository{
		Account: account.New(db),
		Store:   store.New(db),
		Contact: contact.New(db),
	}
}
