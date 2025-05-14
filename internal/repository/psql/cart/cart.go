package cart

import "github.com/jackc/pgx/v5/pgxpool"

type Cart struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) ICartRepository {
	return Cart{
		db: db,
	}
}
