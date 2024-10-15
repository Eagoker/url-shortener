package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetOriginalUrl(c echo.Context) error {
	id := c.Param("id")

	// Ищем оригинальный URL в базе данных
	var originalUrl string
	err := h.db.QueryRow(context.Background(), `
		SELECT original_url FROM urls WHERE short_url = $1
	`, id).Scan(&originalUrl)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "URL not found")
	}

	return c.Redirect(http.StatusTemporaryRedirect, originalUrl)
}
