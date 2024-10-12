package config

import (
	"flag"
	"sync"

	"github.com/caarlos0/env"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
}

var (
	config *Config
	once sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		// Создаем новый экземпляр структуры Config
		config = &Config{}

		// Парсим флаги
		serverAddress := flag.String("address", ":8080", "Web-server address")
		flag.Parse()

		// Разбираем переменные окружения
		err := env.Parse(config) 
		if err != nil || config.ServerAddress == "" {
			// Если переменные окружения не заданы или ошибка парсинга, используем флаги
			config.ServerAddress = *serverAddress
		}
	})

	return config
}

