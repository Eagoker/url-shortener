package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Eagoker/url-shortener/pkg"
	"github.com/Eagoker/url-shortener/internal/database"
	"github.com/labstack/echo/v4"
)

type Request struct {
	OriginalUrl string `json:"url"`
}

type Response struct {
	ShortUrl string `json:"result"`
}

func ConvertToShort(c echo.Context, serverAddress string) error {
	db := database.GetDB()

	if c.Request().Method != http.MethodPost {
		return echo.NewHTTPError(http.StatusMethodNotAllowed, "Method not allowed!")
	}

	var req Request
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body!")
	}

	originalURL := req.OriginalUrl

	// Проверка, существует ли URL в базе данных
	var existingShortURL string
	err := db.QueryRow(c.Request().Context(), "SELECT short_url FROM urls WHERE original_url = $1", originalURL).Scan(&existingShortURL)
	if err == nil {
		// Если короткий URL уже существует в базе данных
		fullShortURL := serverAddressWithSlash(serverAddress) + existingShortURL // Формируем полный URL
		response := Response{
			ShortUrl: fullShortURL, // Возвращаем полный URL
		}
		return c.JSON(http.StatusOK, response)
	}

	if err.Error() != "no rows in result set" {
		// Возникла ошибка при запросе к базе данных, кроме "нет записей"
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to query the database")
	}

	// Генерируем короткий URL
	fullShortURL, err := pkg.GenerateShortURL(originalURL)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error with generating short URL")
	}

	// Вычленяем короткую часть после "/"
	shortURL := fullShortURL[strings.LastIndex(fullShortURL, "/")+1:]

	// Сохраняем короткий URL в базу данных
	_, err = db.Exec(c.Request().Context(), "INSERT INTO urls (original_url, short_url) VALUES ($1, $2)", originalURL, shortURL)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save URL to the database")
	}

	// Формируем полный короткий URL для ответа
	fullShortURL = serverAddressWithSlash(serverAddress) + shortURL

	// Отправляем ответ с полным коротким URL
	response := Response{
		ShortUrl: fullShortURL, // Возвращаем полный URL
	}

	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(http.StatusCreated, response)
}

// Функция для добавления слеша к адресу сервера, если его нет
func serverAddressWithSlash(serverAddress string) string {
	if !strings.HasSuffix(serverAddress, "/") {
		return serverAddress + "/"
	}
	return serverAddress
}
