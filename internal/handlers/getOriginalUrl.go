package handlers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

)

func GetOriginalUrl(c echo.Context) error{
	if c.Request().Method != http.MethodGet{
		return echo.NewHTTPError(http.StatusMethodNotAllowed, "Method not allowed!")
	}
	
	path := c.Request().URL.Path

	// Убираем первый слеш и берем первую часть до следующего слеша
	trimmedPath := strings.TrimPrefix(path, "/")
	id := strings.Split(trimmedPath, "/")[0]

	//тут будет логика получения ориг юрла
	_ = id

	originalUrl := "https://practicum.yandex.ru/"
	
	// w.Header().Set("Location", originalUrl)
	// w.WriteHeader(http.StatusTemporaryRedirect)

	return c.Redirect(http.StatusMovedPermanently, originalUrl)
}