package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type AuthConfig struct {
	Secret string
}

type DbConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	// загружаем в систему данные из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}
	// создаем конфиг
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("SECRET"),
		},
	}
}
