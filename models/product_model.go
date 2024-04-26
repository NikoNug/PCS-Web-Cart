package models

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
}

type Products struct {
	Products []Product `json:"products"`
}
