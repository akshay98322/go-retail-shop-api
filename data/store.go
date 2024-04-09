package data

import "go-retail-shop-api/models"

// Our in-memory 'database' of grocery items.
var GroceryDB = []models.GroceryItem{}

// A simple counter for our IDs, to simulate auto-increment since we're not using a real DB.
var IDCounter = 1
