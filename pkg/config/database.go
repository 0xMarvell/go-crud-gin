package config

import (
	"fmt"
	"log"
	"os"

	"github.com/0xMarvell/simple-blog-posts/pkg/models"
	"github.com/0xMarvell/simple-blog-posts/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes a database connection.
func Connect() {
	// LoadEnv()

	var DBConnectErr error
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host,
		user,
		password,
		dbname,
		port)

	DB, DBConnectErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.CheckErr("could not connect to database: ", DBConnectErr)
	log.Println("Database Connection Successful 👍")

}

// RunMigrations runs migrations for the database.
func RunMigrations() {
	migrationErr := DB.AutoMigrate(&models.Post{})
	utils.CheckErr("Migration failed: ", migrationErr)
	log.Println("Migration Successful 👍")
}
