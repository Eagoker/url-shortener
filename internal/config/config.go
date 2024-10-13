package config

import (
	"flag"
	"sync"

	"github.com/caarlos0/env"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	DatabaseDSN   string `env:"DATABASE_DSN"`
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
		databaseDSN := flag.String("db", "", "DB address")

		flag.Parse()

		// Разбираем переменные окружения
		err := env.Parse(config)
		if err != nil {
			panic("Failed to parse environment variables: " + err.Error())
		}

		// Устанавливаем значения из флагов, если они не заданы через окружение
		if config.ServerAddress == "" {
			config.ServerAddress = *serverAddress
		}
		if config.DatabaseDSN == "" {
			config.DatabaseDSN = *databaseDSN
		}

		// Проверяем, что обязательные параметры заданы
		if config.ServerAddress == "" {
			panic("Server address is not defined! Please provide it via flag or environment variable.")
		}
		if config.DatabaseDSN == "" {
			panic("Database DSN is not defined! Please provide it via flag or environment variable.")
		}
	})

	return config
}
