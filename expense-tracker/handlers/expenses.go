package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/MohdMusaiyab/expense-tracker/config"
	"github.com/MohdMusaiyab/expense-tracker/models"
)

// CreateExpense handler
func CreateExpense(c *gin.Context) {
	var expense models.Expense

	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	if err := config.DB.Create(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Message: "Failed to create expense: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, APIResponse{
		Success: true,
		Message: "Expense created successfully",
		Data:    expense,
	})
}

// GetAllExpenses handler
func GetAllExpenses(c *gin.Context) {
	var expenses []models.Expense

	if err := config.DB.Find(&expenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Message: "Failed to fetch expenses: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Expenses fetched successfully",
		Data:    expenses,
	})
}

// GetExpenseByID handler
func GetExpenseByID(c *gin.Context) {
	id := c.Param("id")
	var expense models.Expense

	if err := config.DB.First(&expense, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, APIResponse{
			Success: false,
			Message: "Expense not found",
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Expense fetched successfully",
		Data:    expense,
	})
}

// UpdateExpense handler
func UpdateExpense(c *gin.Context) {
	id := c.Param("id")
	var expense models.Expense

	if err := config.DB.First(&expense, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, APIResponse{
			Success: false,
			Message: "Expense not found",
		})
		return
	}

	var input models.Expense
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	expense.Amount = input.Amount
	expense.Category = input.Category
	expense.Description = input.Description
	expense.Date = input.Date

	if err := config.DB.Save(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Message: "Failed to update expense: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Expense updated successfully",
		Data:    expense,
	})
}

// DeleteExpense handler
func DeleteExpense(c *gin.Context) {
	id := c.Param("id")
	var expense models.Expense

	if err := config.DB.First(&expense, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, APIResponse{
			Success: false,
			Message: "Expense not found",
		})
		return
	}

	if err := config.DB.Delete(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Message: "Failed to delete expense: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Expense deleted successfully",
	})
}

// GetExpensesByUser handler
func GetExpensesByUser(c *gin.Context) {
	userID := c.Param("id")
	var expenses []models.Expense

	if err := config.DB.Where("user_id = ?", userID).Find(&expenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Message: "Failed to fetch expenses: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Expenses fetched successfully for user",
		Data:    expenses,
	})
}
