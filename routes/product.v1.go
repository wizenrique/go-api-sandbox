package routes

import (
	"github.com/gin-gonic/gin"
	"sandbox/api/controllers"
)

var controller = new(controllers.ProductController)

func SetupProductsRoutes(router *gin.RouterGroup) {

	products := router.Group("products")

	products.GET("/", controller.List)
	products.POST("/", controller.Create)
	products.GET("/:productId", controller.Get)	
	products.DELETE("/:productId", controller.Delete)
	products.PATCH("/:productId", controller.Update)

}