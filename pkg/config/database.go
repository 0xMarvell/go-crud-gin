package config

import (
	"fmt"

	"github.com/0xMarvell/go-crud-gin/pkg/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes a database connection.
func Connect() {
	LoadEnv()

	var err error
	host := viper.GetString("HOST")
	user := viper.GetString("USER")
	password := viper.GetString("PASSWORD")
	dbname := viper.Get("DB_NAME")
	port := viper.GetString("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host,
		user,
		password,
		dbname,
		port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.CheckErr("could not connect to database: ", err)
}
