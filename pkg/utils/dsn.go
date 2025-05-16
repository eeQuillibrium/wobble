package utils

import "os"

const DSN = "postgres://postgres:secret@localhost:5433/wobble?sslmode=disable"

func GetDSN() string {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = DSN
	}

	return dsn
}
