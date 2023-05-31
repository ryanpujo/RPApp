package infrastructure

import "github.com/spf13/viper"

type config struct {
	PORT int    `mapstructure:"port"`
	DSN  string `mapstructure:"dsn"`
}

var Cfg config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/app/")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := viper.Unmarshal(&Cfg)
	if err != nil {
		panic(err)
	}
}
