package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

var Appconfig *viper.Viper

func Init() {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("config")
	config.AddConfigPath("src/config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file", err.Error())
	}
}

func GetConfig() *viper.Viper {
	return config
}
