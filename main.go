package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"sandbox/api/routes"
)

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	routes.SetupRouter(r)
	// Listen and Server in 0.0.0.0:8080
	endless.ListenAndServe(":8080", r)
}