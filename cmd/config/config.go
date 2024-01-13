package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser     string `mapstructure:"POSTGRES_USER"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName     string `mapstructure:"POSTGRES_DB"`
	DBPort     string `mapstructure:"POSTGRES_PORT"`
	DBHost     string `mapstructure:"POSTGRES_HOST"`
	SSLMode    string `mapstructure:"SSL_MODE"`
}

func LoadConfig() *Config {
	log.Println("Starting initialise config")

	viper.AddConfigPath(".")

	viper.SetConfigName(".env")

	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	log.Println("Config has been initialised successful")

	return &config
}
