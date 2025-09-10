package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MohdMusaiyab/blog-app/config"
	"github.com/MohdMusaiyab/blog-app/models"
	"github.com/MohdMusaiyab/blog-app/utils"
	"github.com/gin-gonic/gin"
)

func toUintFromInterface(v interface{}) (uint, error) {
	switch t := v.(type) {
	case uint:
		return t, nil
	case uint8:
		return uint(t), nil
	case uint16:
		return uint(t), nil
	case uint32:
		return uint(t), nil
	case uint64:
		return uint(t), nil
	case int:
		return uint(t), nil
	case int8:
		return uint(t), nil
	case int16:
		return uint(t), nil
	case int32:
		return uint(t), nil
	case int64:
		return uint(t), nil
	case float64:
		return uint(t), nil
	case string:
		u64, err := strconv.ParseUint(t, 10, 64)
		if err != nil {
			return 0, err
		}
		return uint(u64), nil
	default:
		return 0, fmt.Errorf("unsupported id type: %T", v)
	}
}

// GetCurrentUser returns profile of the logged-in user (uses user_id set by middleware)
func GetCurrentUser(c *gin.Context) {
	val, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "user not found in context")
		return
	}

	userID, err := toUintFromInterface(val)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "invalid user id in context")
		return
	}

	var user models.User
	db := config.GetDB()
	if err := db.First(&user, userID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}

	utils.SuccessResponse(c, "current user fetched successfully", user)
}

// GetUserByID returns a user's public profile by id (public route)
func GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id parameter is required")
		return
	}

	idUint, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var user models.User
	db := config.GetDB()
	if err := db.First(&user, uint(idUint)).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}

	utils.SuccessResponse(c, "user fetched successfully", user)
}

// UpdateCurrentUser updates fields of the logged-in user (username, email, password)
func UpdateCurrentUser(c *gin.Context) {
	val, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "user not found in context")
		return
	}

	userID, err := toUintFromInterface(val)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "invalid user id in context")
		return
	}

	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	db := config.GetDB()
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}

	// 🔹 Check uniqueness for username
	if input.Username != "" && input.Username != user.Username {
		var existing models.User
		if err := db.Where("username = ?", input.Username).First(&existing).Error; err == nil {
			utils.ErrorResponse(c, http.StatusBadRequest, "username is already taken")
			return
		}
		user.Username = input.Username
	}

	// 🔹 Check uniqueness for email
	if input.Email != "" && input.Email != user.Email {
		var existing models.User
		if err := db.Where("email = ?", input.Email).First(&existing).Error; err == nil {
			utils.ErrorResponse(c, http.StatusBadRequest, "email is already taken")
			return
		}
		user.Email = input.Email
	}

	// 🔹 Update password if provided
	if input.Password != "" {
		hashed, err := utils.HashPassword(input.Password)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "failed to hash password")
			return
		}
		user.Password = hashed
	}

	if err := db.Save(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to update user")
		return
	}

	utils.SuccessResponse(c, "user updated successfully", user)
}

// DeleteCurrentUser deletes the logged-in user's account
func DeleteCurrentUser(c *gin.Context) {
	val, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "user not found in context")
		return
	}

	userID, err := toUintFromInterface(val)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "invalid user id in context")
		return
	}

	db := config.GetDB()
	if err := db.Delete(&models.User{}, userID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to delete user")
		return
	}

	utils.SuccessResponse(c, "user deleted successfully", nil)
}
