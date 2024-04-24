package models

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type Products struct {
	Products []Product `json:"products"`
}
