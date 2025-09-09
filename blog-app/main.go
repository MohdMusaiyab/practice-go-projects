package main

import (
	"github.com/gin-gonic/gin"
	"github.com/MohdMusaiyab/blog-app/config"
	"github.com/MohdMusaiyab/blog-app/routes"
)

func main() {
	// Initialize DB
	config.InitDB()

	// Create Gin router
	router := gin.Default()

	// Register auth routes
	routes.AuthRoutes(router)
	routes.TestRoutes(router) // ✅ add test route

	// Start server
	router.Run(":" + "8080")
}
