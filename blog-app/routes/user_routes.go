package routes

import (
	"github.com/MohdMusaiyab/blog-app/controllers"
	"github.com/MohdMusaiyab/blog-app/middlewares"
	"github.com/gin-gonic/gin"
)

// UserRoutes registers user-related endpoints.
//
// Public:
//
//	GET  /users/:id     -> controllers.GetUserByID
//
// Protected (require JWT):
//
//	GET   /users/me     -> controllers.GetCurrentUser
//	PUT   /users/me     -> controllers.UpdateCurrentUser
//	DELETE /users/me    -> controllers.DeleteCurrentUser
func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		// Public route: fetch any user by ID
		userGroup.GET("/:id", controllers.GetUserByID)

		// Protected routes: operate on the currently authenticated user
		protected := userGroup.Group("") // same base path /users
		protected.Use(middlewares.JWTAuthMiddleware())
		{
			protected.GET("/me", controllers.GetCurrentUser)
			protected.PUT("/me", controllers.UpdateCurrentUser)
			protected.DELETE("/me", controllers.DeleteCurrentUser)
		}
	}
}
