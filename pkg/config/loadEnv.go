package config

import (
	"github.com/0xMarvell/go-crud-gin/pkg/utils"
	"github.com/spf13/viper"
)

// LoadEnv loads the provided environment variables.
func LoadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	utils.CheckErr("error loading environment variables: ", err)
}
