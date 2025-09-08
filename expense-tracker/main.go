package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/MohdMusaiyab/expense-tracker/config"
	"github.com/MohdMusaiyab/expense-tracker/models"
	"github.com/MohdMusaiyab/expense-tracker/routes"
)

func main() {
	// Connect to Database
	config.ConnectDB()

	// Auto-migrate models (creates tables if they don't exist)
	err := config.DB.AutoMigrate(&models.User{}, &models.Expense{})
	if err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}
	log.Println("✅ Database migrated successfully!")

	// Create a new Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r)

	// Start server
	log.Println("🚀 Server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
