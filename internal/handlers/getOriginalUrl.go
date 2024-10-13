package handlers

import (
	"net/http"
	"strings"

	"github.com/Eagoker/url-shortener/internal/database"
	"github.com/labstack/echo/v4"
)

func GetOriginalUrl(c echo.Context) error {
	db := database.GetDB()

	if c.Request().Method != http.MethodGet {
		return echo.NewHTTPError(http.StatusMethodNotAllowed, "Method not allowed!")
	}

	// Извлекаем короткий URL из пути
	shortURL := strings.TrimPrefix(c.Request().URL.Path, "/")

	// Запрос в базу данных для поиска оригинального URL
	var originalURL string
	err := db.QueryRow(c.Request().Context(), "SELECT original_url FROM urls WHERE short_url = $1", shortURL).Scan(&originalURL)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Short URL not found in the database")
	}

	// Перенаправление на оригинальный URL
	return c.Redirect(http.StatusTemporaryRedirect, originalURL)
}
