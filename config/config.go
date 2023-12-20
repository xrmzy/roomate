package config

import (
	"errors"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	Driver   string
}

type FileConfig struct {
	FilePath string
}

type TokenConfig struct {
	JwtPrivateKey   string
	JwtPublicKey    string
	JwtLifeTime     time.Duration
}

type SheetConfig struct {
	ServiceAccountKey string
	SpreadsheetId     string
}

type Config struct {
	ApiConfig
	DbConfig
	FileConfig
	TokenConfig
	SheetConfig
}

func (c *Config) readConfig() error {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	c.FileConfig = FileConfig{FilePath: os.Getenv("LOG_FILE")}

	tokenLifeTime, err := time.ParseDuration(os.Getenv("TOKEN_LIFE_TIME"))
	if err != nil {
		return err
	}

	c.TokenConfig = TokenConfig{
		JwtPrivateKey:   os.Getenv("JWT_PRIVATE_KEY"),
		JwtPublicKey:    os.Getenv("JWT_PUBLIC_KEY"),
		JwtLifeTime:     tokenLifeTime,
	}

	c.SheetConfig = SheetConfig{
		ServiceAccountKey: os.Getenv("SERVICE_ACCOUNT_KEY"),
		SpreadsheetId:     os.Getenv("SPREADSHEET_ID"),
	}

	if c.ApiConfig.ApiPort == "" || c.DbConfig.Driver == "" || c.DbConfig.Host == "" || c.DbConfig.DbName == "" || c.DbConfig.Port == "" || c.DbConfig.Username == "" || c.DbConfig.Password == "" {
		return errors.New("all environment variables required")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.readConfig(); err != nil {
		return nil, err
	}
	return cfg, nil
}
