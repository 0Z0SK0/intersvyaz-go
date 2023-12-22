package config

import (
	"github.com/spf13/viper"
)

func Load() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
