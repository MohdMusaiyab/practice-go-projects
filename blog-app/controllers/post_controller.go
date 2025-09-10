package controllers

import (
	"net/http"
	"strconv"

	"github.com/MohdMusaiyab/blog-app/config"
	"github.com/MohdMusaiyab/blog-app/models"
	"github.com/MohdMusaiyab/blog-app/utils"
	"github.com/gin-gonic/gin"
)

// 🔹 Create a new post (protected)
func CreatePost(c *gin.Context) {
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
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  userID,
	}

	db := config.GetDB()
	if err := db.Create(&post).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to create post")
		return
	}

	utils.SuccessResponse(c, "post created successfully", post)
}

// 🔹 Get all posts (public)
func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	db := config.GetDB()
	if err := db.Preload("User").Find(&posts).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to fetch posts")
		return
	}

	utils.SuccessResponse(c, "posts fetched successfully", posts)
}

// 🔹 Get post by ID (public)
func GetPostByID(c *gin.Context) {
	id := c.Param("id")
	postID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid post id")
		return
	}

	var post models.Post
	db := config.GetDB()
	if err := db.Preload("User").First(&post, uint(postID)).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "post not found")
		return
	}

	utils.SuccessResponse(c, "post fetched successfully", post)
}

// 🔹 Update post (protected, only owner)
func UpdatePost(c *gin.Context) {
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

	id := c.Param("id")
	postID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid post id")
		return
	}

	var post models.Post
	db := config.GetDB()
	if err := db.First(&post, uint(postID)).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "post not found")
		return
	}

	// check ownership
	if post.UserID != userID {
		utils.ErrorResponse(c, http.StatusForbidden, "not authorized to update this post")
		return
	}

	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if input.Title != "" {
		post.Title = input.Title
	}
	if input.Content != "" {
		post.Content = input.Content
	}

	if err := db.Save(&post).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to update post")
		return
	}

	utils.SuccessResponse(c, "post updated successfully", post)
}

// 🔹 Delete post (protected, only owner)
func DeletePost(c *gin.Context) {
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

	id := c.Param("id")
	postID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid post id")
		return
	}

	var post models.Post
	db := config.GetDB()
	if err := db.First(&post, uint(postID)).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "post not found")
		return
	}

	// check ownership
	if post.UserID != userID {
		utils.ErrorResponse(c, http.StatusForbidden, "not authorized to delete this post")
		return
	}

	if err := db.Delete(&post).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to delete post")
		return
	}

	utils.SuccessResponse(c, "post deleted successfully", nil)
}
