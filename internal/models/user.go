package models

type User struct {
	ID       uint64 `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Login    string `db:"login"`
	Password string `db:"password"`
}
