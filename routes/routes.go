package routes

import (
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

// TODO: add unit test files for routes

func SetupRouter(router *gin.Engine) {

	v1 := router.Group("v1")
	SetupProductsRoutes(v1)

}