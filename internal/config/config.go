package config

import (
	"github.com/spf13/viper"
	"github.com/joho/godotenv"
	"log"
)

var Cfg *viper.Viper

func LoadConfig() error {
	// Load from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	v := viper.New()
	v.AutomaticEnv()

	Cfg = v

	return nil
}
