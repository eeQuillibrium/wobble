package dto

type Product struct {
	ID          int     `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"img_url,omitempty"`
	Category    string  `json:"category,omitempty"`
	Amount      uint64  `json:"amount"`
}
