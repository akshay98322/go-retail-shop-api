package main

import (
	"go-retail-shop-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Create item
	r.POST("/groceries", handlers.AddGroceryItem)
	// List all items
	r.GET("/groceries", handlers.ListGroceryItems)
	// List an item
	r.GET("/groceries/:id", handlers.GetGroceryItem)
	// Update an item (all fields)
	r.PUT("/groceries/:id", handlers.UpdateGroceryItem)
	// Delete an item
	r.DELETE("/groceries/:id", handlers.DeleteGroceryItem)

	r.Run() // default listens on :8080
}
