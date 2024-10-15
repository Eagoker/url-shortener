package config

import (
	"flag"
	"sync"

	"github.com/caarlos0/env"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	DatabaseURL   string `env:"DATABASE_URL"`
	SecretKey     string `env:"SECRET_KEY"` 
}

var (
	config *Config
	once   sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		// Создаем новый экземпляр структуры Config
		config = &Config{}

		// Парсим флаги
		serverAddress := flag.String("address", "localhost:8080", "Web-server address")
		databaseURL := flag.String("db", "", "Database URL")
		secretKey := flag.String("secret", "", "JWT Secret Key") // Новый флаг для секретного ключа
		flag.Parse()

		// Разбираем переменные окружения
		err := env.Parse(config)
		if err != nil || config.ServerAddress == "" {
			// Если переменные окружения не заданы или ошибка парсинга, используем флаги
			config.ServerAddress = *serverAddress
		}

		if config.DatabaseURL == "" {
			// Если переменная окружения DATABASE_URL не задана, используем флаг
			config.DatabaseURL = *databaseURL
		}

		if config.SecretKey == "" {
			// Если переменная окружения SECRET_KEY не задана, используем флаг
			config.SecretKey = *secretKey
		}
	})

	return config
}
