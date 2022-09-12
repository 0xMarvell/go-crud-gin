package config

import (
	"github.com/0xMarvell/simple-blog-posts/pkg/utils"
	"github.com/joho/godotenv"
)

// LoadEnv loads the provided environment variables.
func LoadEnv() {
	err := godotenv.Load()
	utils.CheckErr("error loading environment variables: ", err)
}
