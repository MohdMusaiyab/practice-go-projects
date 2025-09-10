package controllers

import (
	"net/http"
	"strconv"

	"github.com/MohdMusaiyab/blog-app/config"
	"github.com/MohdMusaiyab/blog-app/models"
	"github.com/MohdMusaiyab/blog-app/utils"
	"github.com/gin-gonic/gin"
)

// CreateComment creates a comment (protected)
func CreateComment(c *gin.Context) {
	var input struct {
		Content string `json:"content" binding:"required"`
		PostID  uint   `json:"post_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	val, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		return
	}
	userID, err := toUintFromInterface(val)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "invalid user id in context")
		return
	}

	comment := models.Comment{
		Content: input.Content,
		PostID:  input.PostID,
		UserID:  userID,
	}

	db := config.GetDB()
	if err := db.Create(&comment).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to create comment")
		return
	}

	utils.SuccessResponse(c, "comment created successfully", comment)
}

// GetCommentsByPost returns all comments for a given post (public)
func GetCommentsByPost(c *gin.Context) {
	postIDParam := c.Param("postID")
	if postIDParam == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "postID parameter is required")
		return
	}
	postID, err := strconv.ParseUint(postIDParam, 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid postID parameter")
		return
	}

	var comments []models.Comment
	db := config.GetDB()
	if err := db.Where("post_id = ?", uint(postID)).Preload("User").Find(&comments).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to fetch comments")
		return
	}

	utils.SuccessResponse(c, "comments fetched successfully", comments)
}

// UpdateComment updates a comment (protected, owner only)
func UpdateComment(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id parameter is required")
		return
	}
	commentID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	db := config.GetDB()
	var comment models.Comment
	if err := db.First(&comment, uint(commentID)).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "comment not found")
		return
	}

	val, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		return
	}
	userID, err := toUintFromInterface(val)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "invalid user id in context")
		return
	}
	if comment.UserID != userID {
		utils.ErrorResponse(c, http.StatusForbidden, "not allowed to update this comment")
		return
	}

	var input struct {
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	comment.Content = input.Content
	if err := db.Save(&comment).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to update comment")
		return
	}

	utils.SuccessResponse(c, "comment updated successfully", comment)
}

// DeleteComment deletes a comment (protected, owner only)
func DeleteComment(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id parameter is required")
		return
	}
	commentID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	db := config.GetDB()
	var comment models.Comment
	if err := db.First(&comment, uint(commentID)).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "comment not found")
		return
	}

	val, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		return
	}
	userID, err := toUintFromInterface(val)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "invalid user id in context")
		return
	}
	if comment.UserID != userID {
		utils.ErrorResponse(c, http.StatusForbidden, "not allowed to delete this comment")
		return
	}

	if err := db.Delete(&comment).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to delete comment")
		return
	}

	utils.SuccessResponse(c, "comment deleted successfully", nil)
}
