package routes

import (
	"github.com/MohdMusaiyab/blog-app/controllers"
	"github.com/MohdMusaiyab/blog-app/middlewares"
	"github.com/gin-gonic/gin"
)

func CommentRoutes(r *gin.Engine) {
	// Public route: fetch all comments for a post
	r.GET("/posts/:postID/comments", controllers.GetCommentsByPost)

	// Protected routes
	comments := r.Group("/comments")
	comments.Use(middlewares.JWTAuthMiddleware())
	{
		comments.POST("/", controllers.CreateComment)
		comments.PUT("/:id", controllers.UpdateComment)
		comments.DELETE("/:id", controllers.DeleteComment)
	}
}
