package account

import "github.com/jackc/pgx/v5/pgxpool"

type Account struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) IAccountRepository {
	return &Account{
		db: db,
	}
}
