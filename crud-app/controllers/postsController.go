package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rynfkn/crud-app/initializers"
	"github.com/rynfkn/crud-app/models"
)

func GetAllPost(ctx *gin.Context) {
	var posts []models.Post

	// get all the posts
	if err := initializers.DB.Find(&posts).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"posts" : posts,
	})
}

func GetPostById(ctx *gin.Context) {
	// get the id
	id := ctx.Param("id")
	var post models.Post
	
	// find the post based on the id
	if err := initializers.DB.Find(&post, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"post" : post,
	})
}

func AddNewPost(ctx *gin.Context) {
	var post models.Post

	// bind the json data to the model
	if err := ctx.BindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	// save to database
	if err := initializers.DB.Create(&post).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {
			"error" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func UpdatePostById(ctx *gin.Context) {
	//get the id
	id := ctx.Param("id")
	var post models.Post

	// bind the json data to the model
	if err := ctx.BindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// update the data in database
	if err := initializers.DB.Model(&post).Where("id = ?", id).Updates(post).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func DeletePostById(ctx *gin.Context) {
	id := ctx.Param("id")
	var post models.Post

	if err := initializers.DB.Delete(&post, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}