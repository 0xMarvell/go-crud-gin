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
	r := gin.Default()
	routes.RegisterBlogRoutes(r)

	log.Fatal(r.Run())
}
