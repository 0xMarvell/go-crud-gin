package main

import (
	"log"

	"github.com/0xMarvell/simple-blog-posts/pkg/config"
	"github.com/0xMarvell/simple-blog-posts/pkg/models"
)

func init() {
	config.LoadEnv()
	config.Connect()
}

func main() {
	config.DB.AutoMigrate(&models.Post{})
	log.Println("Migration successful 👍")
}
