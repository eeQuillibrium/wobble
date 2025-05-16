package dto

type Order struct {
	Items []OrderedProduct `json:"items"`
}

type OrderedProduct struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Price    uint64 `json:"price"`
	Quantity uint64 `json:"quantity"`
}
