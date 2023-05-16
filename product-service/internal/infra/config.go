package infra

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port int    `mapstructure:"port"`
	DSN  string `mapstructure:"dsn"`
}

var config Config

func LoadConfig() Config {
	if config.DSN != "" && config.Port != 0 {
		return config
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	return config
}
