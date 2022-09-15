package main

import (
	"log"

	"github.com/0xMarvell/simple-blog-posts/pkg/config"
	"github.com/0xMarvell/simple-blog-posts/pkg/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.Connect()
	config.RunMigrations()
}

func main() {
	// To run Gin server in debug mode, comment out line 19
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	routes.RegisterBlogRoutes(r)

	log.Println("Server is up and running...")
	log.Fatal(r.Run())
}
