package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB_USER     string `mapstructure:"POSTGRES_USER"`
	DB_PASSWORD string `mapstructure:"POSTGRES_PASSWORD"`
	DB_HOST     string `mapstructure:"POSTGRES_HOST"`
	DB_PORT     string `mapstructure:"POSTGRES_PORT"`
	DB_NAME     string `mapstructure:"DB_NAME"`
	API_PORT    string `mapstructure:"API_PORT"`
	FileConfig
	TokenConfig
}

type FileConfig struct{}

type TokenConfig struct{}

func NewConfig() (Config, error) {
	cfg := Config{}
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("☠️ cannot read configuration")
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("☠️ Environment can't be loaded: ", err)
	}
	return cfg, nil
}
