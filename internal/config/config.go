package config

import (
	"flag"
	"sync"
)

type Config struct {
	ServerAddress string
}

var (
	config *Config
	once sync.Once
)

func GetConfig() *Config{
	once.Do(func() {
		serverAddress := flag.String("address", "localhost:8080", "Web-server address")
		flag.Parse()

		config = &Config{
			ServerAddress: *serverAddress,
		}

	})

	return config
}