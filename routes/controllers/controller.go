package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/sodabitsdev/inventory-app/models"
	"github.com/sodabitsdev/inventory-app/utilities"
)

/*

// List all todos
func GetTodos(c *gin.Context) {
	var todo []Models.Todo
	err := Models.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

*/

// FindAllInventories retrieves all Inventories
func FindAllInventories(c *gin.Context) {
	log.Infoln("controller.FindAllInventories called")

	var inv []models.Inventory
	inv, err := models.FindAllInventories(utilities.DB())

	if err != nil {
		log.Errorln("Error from controller.FindAllInventories: ", err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, inv)
	}
}

// func FindAllInventoriesByDate(c *gin.Context) {

// }

// func GetInventoryByBarcode(c *gin.Context) {

// }

// func InsertInventory(c *gin.Context) {

// }

// func UpdateInventory(c *gin.Context) {

// }

// v1.GET("inventory", controllers.FindAllInventories)
// v1.GET("inventory/:date", controllers.FindAllInventoriesByDate)
// v1.GET("inventory/:barcode", controllers.GetInventoryByBarcode)
// v1.POST("inventory", controllers.InsertInventory)
// v1.PUT("inventory/:id", controllers.UpdateInventory)

// // List all todos
// func GetTodos(c *gin.Context) {
// 	var todo []Models.Todo
// 	err := Models.GetAllTodos(&todo)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, todo)
// 	}

// }

// // Create a Todo
// func CreateATodo(c *gin.Context) {
// 	var todo Models.Todo
// 	c.BindJSON(&todo)
// 	err := Models.CreateATodo(&todo)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, todo)
// 	}
// }

// // Get a particular Todo with id
// func GetATodo(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var todo Models.Todo
// 	err := Models.GetATodo(&todo, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, todo)
// 	}
// }

// // Update an existing Todo
// func UpdateATodo(c *gin.Context) {
// 	var todo Models.Todo
// 	id := c.Params.ByName("id")
// 	err := Models.GetATodo(&todo, id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, todo)
// 	}
// 	c.BindJSON(&todo)
// 	err = Models.UpdateATodo(&todo, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, todo)
// 	}
// }

// // Delete a Tdodo
// func DeleteATodo(c *gin.Context) {
// 	var todo Models.Todo
// 	id := c.Params.ByName("id")
// 	err := Models.DeleteATodo(&todo, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"id:" + id: "delete"})
// 	}
// }
