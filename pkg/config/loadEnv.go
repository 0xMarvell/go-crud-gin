package config

import (
	"github.com/joho/godotenv"
)

// LoadEnv loads the provided environment variables.
func LoadEnv() {
	godotenv.Load()
	// utils.CheckErr("error loading environment variables: ", err)
}
