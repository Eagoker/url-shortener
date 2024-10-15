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

	shortURL, err := pkg.GenerateShortURL(requestBody.OriginalUrl)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error generating short URL")
	}

	// Сохраняем URL в базу данных
	_, err = h.db.Exec(context.Background(), `
		INSERT INTO urls (short_url, original_url) VALUES ($1, $2)
	`, shortURL, requestBody.OriginalUrl)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error saving URL")
	}

	fullURL := "http://" + h.serverAddress + "/" + shortURL
	return c.JSON(http.StatusCreated, Response{ShortUrl: fullURL})
}
