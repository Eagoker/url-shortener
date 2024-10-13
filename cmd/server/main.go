package main

import (
	"log"
	"net/http"

	"github.com/Eagoker/url-shortener/internal/handlers"
	"github.com/Eagoker/url-shortener/internal/config"
	"github.com/Eagoker/url-shortener/internal/logger"
	"github.com/Eagoker/url-shortener/internal/database"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// Чтение конфигурации
	cfg := config.GetConfig()

	// Инициализация логгера
	zapLogger, err := logger.InitLogger()
	if err != nil {
		log.Fatalf("Ошибка при инициализации логгера: %v", err)
	}
	defer zapLogger.Sync()

	// Настройка Echo
	e := echo.New()
	// БД
	connStr := cfg.DatabaseDSN // Или используйте другую конфигурацию
    database.ConnectDB(connStr)

	// Подключение middleware для логгирования запросов
	e.Use(logger.RequestLoggerMiddleware(zapLogger))
	e.Use(middleware.Gzip())

	// Маршруты
	e.POST("/", func(c echo.Context) error {
		return handlers.ConvertToShort(c, cfg.ServerAddress) // Передаем адрес сервера
	})
	e.GET("/:id", handlers.GetOriginalUrl)

	// Запуск сервера
	zapLogger.Info("Starting server", zap.String("address", cfg.ServerAddress))
	if err := e.Start(cfg.ServerAddress); err != http.ErrServerClosed {
		zapLogger.Fatal("Server failed to start", zap.Error(err))
	}
}
