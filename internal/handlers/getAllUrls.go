package handlers

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type UrlResponse struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}

func (h *Handler) GetUserUrls(c echo.Context) error {
	// Извлекаем токен из куки
	cookie, err := c.Cookie("jwt")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Unauthorized, no token provided",
		})
	}

	// Парсим токен
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		// Проверка алгоритма
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
		}
		return h.cfg.SecretKey, nil
	})
	if err != nil || !token.Valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Unauthorized, invalid token",
		})
	}

	// Извлекаем username из токена
	var username string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username = claims["username"].(string)
	} else {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to extract username from token",
		})
	}

	// Ищем user_id по имени пользователя
	var userID int
	err = h.db.QueryRow(context.Background(), `
		SELECT id FROM users WHERE username = $1
	`, username).Scan(&userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to find user",
		})
	}

	// Запрашиваем все URL-адреса, связанные с user_id
	rows, err := h.db.Query(context.Background(), `
		SELECT original_url, short_url FROM urls WHERE user_id = $1
	`, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch URLs",
		})
	}
	defer rows.Close()

	var urls []UrlResponse
	for rows.Next() {
		var url UrlResponse
		if err := rows.Scan(&url.OriginalURL, &url.ShortURL); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to scan URL",
			})
		}
		urls = append(urls, url)
	}

	// Возвращаем список URL-адресов
	return c.JSON(http.StatusOK, urls)
}
