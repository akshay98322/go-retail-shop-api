package handlers

import (
	"go-retail-shop-api/data"
	"go-retail-shop-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddGroceryItem(c *gin.Context) {
	var newItem models.GroceryItem

	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign a new ID and append to the database
	newItem.ID = data.IDCounter
	data.GroceryDB = append(data.GroceryDB, newItem)
	data.IDCounter++

	c.JSON(http.StatusCreated, newItem)
}

// Fetch a specific grocery item by ID
func GetGroceryItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, item := range data.GroceryDB {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

// List all grocery items
func ListGroceryItems(c *gin.Context) {
	c.JSON(http.StatusOK, data.GroceryDB)
}

// Update a specific grocery item by ID
func UpdateGroceryItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var updatedItem models.GroceryItem
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for index, item := range data.GroceryDB {
		if item.ID == id {
			data.GroceryDB[index].Name = updatedItem.Name
			data.GroceryDB[index].Count = updatedItem.Count
			data.GroceryDB[index].Category = updatedItem.Category
			c.JSON(http.StatusOK, data.GroceryDB[index])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

// Delete a specific grocery item by ID
func DeleteGroceryItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for index, item := range data.GroceryDB {
		if item.ID == id {
			// Remove item from our "database"
			data.GroceryDB = append(data.GroceryDB[:index], data.GroceryDB[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}
