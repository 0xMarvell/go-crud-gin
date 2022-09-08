package routes

import (
	"github.com/0xMarvell/simple-blog-posts/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterBlogRoutes(r *gin.Engine) {
	r.POST("/api/v1/blogposts", controllers.CreatePost)
	r.GET("/api/v1/blogposts", controllers.GetPosts)
}
