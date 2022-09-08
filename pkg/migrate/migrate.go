package main

import (
	"log"

	"github.com/0xMarvell/simple-blog-posts/pkg/config"
	"github.com/0xMarvell/simple-blog-posts/pkg/models"
	"github.com/0xMarvell/simple-blog-posts/pkg/utils"
)

func init() {
	config.LoadEnv()
	config.Connect()
}

func main() {
	err := config.DB.AutoMigrate(&models.Post{})
	utils.CheckErr("Migration failed: ", err)
	log.Println("Migration successful 👍")
}
