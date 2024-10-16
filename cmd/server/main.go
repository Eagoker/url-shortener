package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Eagoker/url-shortener/internal/handlers"
	"github.com/Eagoker/url-shortener/internal/config"
	"github.com/Eagoker/url-shortener/internal/logger"

	"github.com/Eagoker/url-shortener/internal/middleware"

	"github.com/jackc/pgx/v4/pgxpool"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
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

	// Настройка подключения к базе данных
	dbURL := cfg.DatabaseURL
	dbpool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		zapLogger.Fatal("Unable to connect to database", zap.Error(err))
	}
	defer dbpool.Close()

	// Создание таблицы пользователей, если она не существует
	_, err = dbpool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) UNIQUE NOT NULL,
			password_hash TEXT NOT NULL
		)
	`)
	if err != nil {
		zapLogger.Fatal("Failed to create users table", zap.Error(err))
	}

	// Создание таблицы URL, если она не существует
	_, err = dbpool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS urls (
			id SERIAL PRIMARY KEY,
			short_url VARCHAR(255) UNIQUE NOT NULL,
			original_url TEXT NOT NULL,
			user_id INTEGER REFERENCES users(id) ON DELETE CASCADE
		)
	`)
	if err != nil {
		zapLogger.Fatal("Failed to create urls table", zap.Error(err))
	}

	// Настройка Echo
	e := echo.New()
	// БД

	// Подключение middleware для логгирования запросов
	e.Use(logger.RequestLoggerMiddleware(zapLogger))
	e.Use(echoMiddleware.Gzip())

	// Передаем адрес сервера и пул подключений к базе данных в хендлеры
	h := handlers.NewHandler(cfg.ServerAddress, dbpool, cfg)

	// Маршруты

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	// Ваша функция main остаётся прежней
	e.POST("/", h.ConvertToShort, middleware.JwtMiddleware(cfg))
	e.GET("/", h.GetUserUrls, middleware.JwtMiddleware(cfg))
	e.GET("/:id", h.GetOriginalUrl, middleware.JwtMiddleware(cfg))

	// Запуск сервера
	zapLogger.Info("Starting server", zap.String("address", cfg.ServerAddress))
	if err := e.Start(cfg.ServerAddress); err != http.ErrServerClosed {
		zapLogger.Fatal("Server failed to start", zap.Error(err))
	}
}
