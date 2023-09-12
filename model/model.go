package model

type Product struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
