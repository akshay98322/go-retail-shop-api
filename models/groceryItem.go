package models

type GroceryItem struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Count    int    `json:"count"`
	Category string `json:"category"`
}
