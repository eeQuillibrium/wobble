package models

import "time"

type Order struct {
	ID              uint64           `db:"id"`
	CreatedAt       time.Time        `db:"created_at"`
	Status          string           `db:"status"`
	Items           []OrderedProduct `db:"items"`
	TotalAmount     uint64           `db:"total_amount"`
	DeliveryAddress *string          `db:"delivery_address"`
}

type OrderedProduct struct {
	ID          int     `db:"id" json:"id"`
	Name        string  `db:"name" json:"name"`
	Description string  `db:"description" json:"description"`
	Price       float64 `db:"price" json:"price"`
	ImageURL    string  `db:"img_url" json:"imageUrl"`
	Quantity    int     `db:"quantity" json:"quantity"`
	Category    string  `db:"category"`
}
