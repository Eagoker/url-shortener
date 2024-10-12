package main

import (
	"log"
	"net/http"

	"github.com/Eagoker/url-shortener/internal/handlers"
	"github.com/Eagoker/url-shortener/internal/config"
	"github.com/Eagoker/url-shortener/internal/logger"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	// Чтение конфигурации
	config := config.GetConfig()

	// Инициализация логгера
	zapLogger, err := logger.InitLogger()
	if err != nil {
		log.Fatalf("Ошибка при инициализации логгера: %v", err)
	}
	defer zapLogger.Sync()

	// Настройка Echo
	e := echo.New()

	// Подключение middleware для логгирования запросов
	e.Use(logger.RequestLoggerMiddleware(zapLogger))

	// Маршруты
	e.POST("/", handlers.ConvertToShort)
	e.GET("/:id", handlers.GetOriginalUrl)
	e.POST("/api/shorten/", handlers.ApiShorten)

	// Запуск сервера
	zapLogger.Info("Starting server", zap.String("address", config.ServerAddress))
	if err := e.Start(config.ServerAddress); err != http.ErrServerClosed {
		zapLogger.Fatal("Server failed to start", zap.Error(err))
	}
}
