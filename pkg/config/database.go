package config

import (
	"fmt"
	"log"

	"github.com/0xMarvell/simple-blog-posts/pkg/models"
	"github.com/0xMarvell/simple-blog-posts/pkg/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes a database connection.
func Connect() {
	LoadEnv()

	var DBConnectErr error
	host := viper.GetString("DB_HOST")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	dbname := viper.Get("DB_NAME")
	port := viper.GetString("DB_PORT")

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
	fmt.Println()
}
