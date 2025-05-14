package store

import "github.com/jackc/pgx/v5/pgxpool"

type Store struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) IStoreRepository {
	return Store{
		db: db,
	}
}
