package main

import (
	"log"
	"net/http"

	"github.com/Eagoker/url-shortener/internal/handlers"
	"github.com/Eagoker/url-shortener/internal/config"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetConfig()

	e := echo.New()

	e.POST("/", handlers.ConvertToShort)
	e.GET("/:id", handlers.GetOriginalUrl)

	if err := e.Start(config.ServerAddress); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	
}