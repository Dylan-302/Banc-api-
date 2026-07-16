package common

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ConfigDB struct {
	Host,
	Password,
	Username,
	Charset,
	Dbname,
	Port string
}

type Config struct {
	DB  *ConfigDB
	App *ConfigApp
}
type ConfigApp struct {
	Port string
}

func NewConfig() *Config {
	if _, err := os.Stat(".env"); err == nil {
		if loadErr := godotenv.Load(".env"); loadErr != nil {
			fmt.Println("error al cargar el .env", loadErr)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return &Config{
		App: &ConfigApp{
			Port: port,
		},
		DB: &ConfigDB{
			Host:     os.Getenv("DB_HOST"),
			Password: os.Getenv("DB_PASSWORD"),
			Username: os.Getenv("DB_USER_NAME"),
			Dbname:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
		},
	}
}
