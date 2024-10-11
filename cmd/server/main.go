package main

import (
	"log"
	"net/http"

	"github.com/Eagoker/url-shortener/internal/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", handlers.ConvertToShort)
	// mux.HandleFunc("/{id}", handlers.GetOriginalUrl)
	e := echo.New()

	e.POST("/", handlers.ConvertToShort)
	e.GET("/:id", handlers.GetOriginalUrl)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	
}