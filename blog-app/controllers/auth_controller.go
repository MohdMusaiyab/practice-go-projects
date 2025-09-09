package controllers

import (
	"net/http"

	"github.com/MohdMusaiyab/blog-app/config"
	"github.com/MohdMusaiyab/blog-app/dto"
	"github.com/MohdMusaiyab/blog-app/models"
	"github.com/MohdMusaiyab/blog-app/utils"
	"github.com/gin-gonic/gin"
)

// Register handles user signup
func Register(c *gin.Context) {
	var input dto.RegisterDTO

	// Bind incoming JSON to DTO
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Validate DTO
	if err := utils.ValidateStruct(input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to hash password")
		return
	}

	// Create user object
	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
	}

	// Save user to DB
	db := config.GetDB()
	if err := db.Create(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "email or username already exists")
		return
	}

	utils.SuccessResponse(c, "user registered successfully", nil)
}

// Login handles user authentication
func Login(c *gin.Context) {
	var input dto.LoginDTO

	// Bind incoming JSON to DTO
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Validate DTO
	if err := utils.ValidateStruct(input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Find user by email
	var user models.User
	db := config.GetDB()
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "invalid credentials")
		return
	}

	// Check password
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		utils.ErrorResponse(c, http.StatusUnauthorized, "invalid credentials")
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to generate token")
		return
	}

	// Respond with token inside "data"
	utils.SuccessResponse(c, "login successful", gin.H{"token": token})
}
