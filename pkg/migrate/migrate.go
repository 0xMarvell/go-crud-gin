package main

import (
	"log"

	"github.com/0xMarvell/go-crud-gin/pkg/config"
	"github.com/0xMarvell/go-crud-gin/pkg/models"
)

func init() {
	config.LoadEnv()
	config.Connect()
}

func main() {
	config.DB.AutoMigrate(&models.Post{})
	log.Println("Migration successful 👍")
}
