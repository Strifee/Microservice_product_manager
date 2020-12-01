package models

// Product type
type Product struct {
	ID        int     `jason:"id,omitempty"`
	Name      string  `json:"name"`
	UnitPrice float64 `json:"price"`
}
