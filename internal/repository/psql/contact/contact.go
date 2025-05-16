package contact

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Contact struct {
	db *pgxpool.Pool
	qb squirrel.StatementBuilderType
}

func New(db *pgxpool.Pool) IContactRepository {
	qb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	return &Contact{
		db: db,
		qb: qb,
	}
}
