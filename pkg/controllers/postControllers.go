package controllers

import (
	"github.com/0xMarvell/simple-blog-posts/pkg/config"
	"github.com/0xMarvell/simple-blog-posts/pkg/models"
	"github.com/0xMarvell/simple-blog-posts/pkg/utils"
	"github.com/gin-gonic/gin"
)

// CreatePost creates a new blog post.
func CreatePost(c *gin.Context) {
	var newPost struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	err := c.Bind(&newPost)
	utils.CheckErr("c.Bind error: ", err)

	post := models.Post{Title: newPost.Title, Body: newPost.Body}
	result := config.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"post":    post,
	})
}
