package infrastructure

import (
	"github.com/spf13/viper"
)

type service struct {
	Address     string `mapstructure:"address"`
	ServicePort int    `mapstructure:"servicePort"`
}

type Config struct {
	Services map[string]service `mapstructure:"services"`
	Port     int                `mapstructure:"port"`
}

var config Config

func LoadConfig() Config {
	if config.Services != nil && config.Port != 0 {
		return config
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return config
}
