package config

import (
	"common/exception"
	"github.com/spf13/viper"
)

var config = new(Config)

func init() {
	viper.SetConfigFile("config/server.yml")
	if err := viper.ReadInConfig(); err != nil {
		panic(exception.GetConfigFailed(err))
		return
	}
	if err := viper.Unmarshal(config); err != nil {
		panic(exception.ParseConfigFailed(err))
		return
	}
}

func GetConfig() *Config {
	return config
}
