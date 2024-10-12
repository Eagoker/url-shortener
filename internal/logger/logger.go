package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"github.com/labstack/echo/v4"
)

// InitLogger инициализирует zap-логгер и возвращает его
func InitLogger() (*zap.Logger, error) {
	// Открытие файла для записи логов
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	writeSyncer := zapcore.AddSync(logFile)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, writeSyncer, zap.InfoLevel)
	logger := zap.New(core)

	return logger, nil
}

// RequestLoggerMiddleware возвращает middleware для логгирования запросов и ответов
func RequestLoggerMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			// Вызов следующего обработчика
			err := next(c)

			// Время выполнения запроса
			latency := time.Since(start)

			// Размер ответа
			res := c.Response()
			contentLength := res.Size

			// Логгирование запроса и ответа
			logger.Info("HTTP request and response",
				zap.String("method", c.Request().Method),
				zap.String("uri", c.Request().RequestURI),
				zap.Int("status", res.Status),
				zap.Duration("latency", latency),
				zap.Int64("response_size", contentLength),
			)

			return err
		}
	}
}
