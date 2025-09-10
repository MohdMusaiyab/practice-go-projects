package routes

import (
	"github.com/MohdMusaiyab/blog-app/controllers"
	"github.com/MohdMusaiyab/blog-app/middlewares"
	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {
	posts := r.Group("/posts")

	// Public routes
	posts.GET("/", controllers.GetAllPosts)
	posts.GET("/:id", controllers.GetPostByID)

	// Protected routes (require auth)
	posts.Use(middlewares.JWTAuthMiddleware())
	{
		posts.POST("/", controllers.CreatePost)
		posts.PUT("/:id", controllers.UpdatePost)
		posts.DELETE("/:id", controllers.DeletePost)
	}
}
