package models

type Product struct {
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Price       float64 `db:"price"`
	ImageURL    string  `db:"img_url"`
	Category    string  `db:"category"`
}
