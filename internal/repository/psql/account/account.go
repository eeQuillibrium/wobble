package account

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Account struct {
	db *pgxpool.Pool
	qb squirrel.StatementBuilderType
}

func New(db *pgxpool.Pool) IAccountRepository {
	qb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	return &Account{
		db: db,
		qb: qb,
	}
}

func (a *Account) WrapWithTransaction(ctx context.Context, f func(tx pgx.Tx) error) error {
	tx, err := a.db.Begin(ctx)
	if err != nil {
		return tx.Rollback(ctx)
	}

	if err := f(tx); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}
