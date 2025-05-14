package main

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

// ApplyMigrations применяет SQL-миграции из директории ./migrations.
// Использует библиотеку golang-migrate/migrate.
func ApplyMigrations(pool *pgxpool.Pool) {
	goose.SetBaseFS(nil)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	db := stdlib.OpenDBFromPool(pool)

	if err := goose.Up(db, "migrations/postgres"); err != nil {
		panic(err)
	}
}
