package config

import (
	"github.com/0xMarvell/simple-blog-posts/pkg/utils"
	"github.com/spf13/viper"
)

// LoadEnv loads the provided environment variables.
func LoadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	utils.CheckErr("error loading environment variables: ", err)
}
