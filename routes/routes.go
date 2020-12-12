package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sodabitsdev/inventory-app/routes/controllers"
)

// SetupRouter sets up default gin Engine and also REST verb mappings. It then
// returns a point to gin engine
func SetupRouter() *gin.Engine {
	// returns a gin Engine instance with the Logger and Recovery middleware
	r := gin.Default()

	// Group() function will make the APIs accessible on /v1
	v1 := r.Group("/v1")
	{
		// v1.GET("todo", Controllers.GetTodos)
		// v1.POST("todo", Controllers.CreateATodo)
		// v1.GET("todo/:id", Controllers.GetATodo)
		// v1.PUT("todo/:id", Controllers.UpdateATodo)
		// v1.DELETE("todo/:id", Controllers.DeleteATodo)

		v1.GET("inventory", controllers.FindAllInventories)
		// v1.GET("inventory/:date", controller.FindAllInventoriesByDate)
		// v1.GET("inventory/:barcode", controller.GetInventoryByBarcode)
		// v1.POST("inventory", controller.InsertInventory)
		// v1.PUT("inventory/:id", controller.UpdateInventory)

	}

	return r
}
