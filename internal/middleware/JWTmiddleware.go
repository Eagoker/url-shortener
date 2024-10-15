package middleware

import (
	"github.com/Eagoker/url-shortener/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JwtMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(cfg.SecretKey),
	})
}


