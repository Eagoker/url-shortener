package handlers

import (
	"context"
	"net/http"

	"github.com/Eagoker/url-shortener/pkg"
	"github.com/labstack/echo/v4"
)

type Request struct {
	OriginalUrl string `json:"url"`
}

type Response struct {
	ShortUrl string `json:"result"`
}

func (h *Handler) ConvertToShort(c echo.Context) error {
	requestBody := new(Request)

	if err := c.Bind(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	// Проверяем, существует ли оригинальный URL в базе данных
	var existingShortURL string
	err := h.db.QueryRow(context.Background(), `
		SELECT short_url FROM urls WHERE original_url = $1
	`, requestBody.OriginalUrl).Scan(&existingShortURL)

	if err == nil {
		// Если URL уже существует, возвращаем его
		fullURL := "http://" + h.serverAddress + "/" + existingShortURL
		return c.JSON(http.StatusOK, Response{ShortUrl: fullURL})
	}

	// Генерируем новый короткий URL
	shortURL, err := pkg.GenerateShortURL(requestBody.OriginalUrl)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error generating short URL")
	}

	// Сохраняем новый URL в базу данных
	_, err = h.db.Exec(context.Background(), `
		INSERT INTO urls (short_url, original_url) VALUES ($1, $2)
	`, shortURL, requestBody.OriginalUrl)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error saving URL")
	}

	fullURL := "http://" + h.serverAddress + "/" + shortURL
	return c.JSON(http.StatusCreated, Response{ShortUrl: fullURL})
}
