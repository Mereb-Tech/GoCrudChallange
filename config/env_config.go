package config

import (
	models "GoCrudChallenge/Domain/Models"
	"log"

	"github.com/spf13/viper"
)


func NewConfig() *models.Env {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv() 

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found.")
	}

	return &models.Env{
		BASE_URL: viper.GetString("BASE_URL"),
	}
}
