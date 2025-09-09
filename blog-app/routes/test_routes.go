package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/MohdMusaiyab/blog-app/controllers"
	"github.com/MohdMusaiyab/blog-app/middlewares"
)

// TestRoutes defines routes used for testing auth middleware
func TestRoutes(router *gin.Engine) {
	protected := router.Group("/test")
	protected.Use(middlewares.JWTAuthMiddleware())
	{
		protected.GET("/protected", controllers.ProtectedTest)
	}
}
