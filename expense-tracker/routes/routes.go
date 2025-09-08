package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/MohdMusaiyab/expense-tracker/handlers"
)

// RegisterRoutes sets up all API routes for the project
func RegisterRoutes(r *gin.Engine) {
	// Group API under /api/v1
	api := r.Group("/api/v1")

	// User routes
	api.POST("/users", handlers.CreateUser)
	api.GET("/users", handlers.GetAllUsers)
	api.GET("/users/:id", handlers.GetUserByID)
	api.PUT("/users/:id", handlers.UpdateUser)
	api.DELETE("/users/:id", handlers.DeleteUser)

	// Expense routes
	api.POST("/expenses", handlers.CreateExpense)
	api.GET("/expenses", handlers.GetAllExpenses)
	api.GET("/expenses/:id", handlers.GetExpenseByID)
	api.PUT("/expenses/:id", handlers.UpdateExpense)
	api.DELETE("/expenses/:id", handlers.DeleteExpense)

	// Extra: fetch all expenses by user
	api.GET("/users/:id/expenses", handlers.GetExpensesByUser)
}
