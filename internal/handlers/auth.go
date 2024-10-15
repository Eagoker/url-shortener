package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


func (h *Handler) Login(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	// Ищем пользователя в базе данных
	var passwordHash string
	err := h.db.QueryRow(context.Background(), `
		SELECT password_hash FROM users WHERE username = $1
	`, req.Username).Scan(&passwordHash)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid username or password",
		})
	}

	// Проверяем пароль
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid username or password",
		})
	}

	// Генерация JWT токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Username,
		"exp":      time.Now().Add(time.Minute * 10).Unix(), 
	})

	tokenString, err := token.SignedString(h.cfg.SecretKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to generate token",
		})
	}

	// Установка JWT в куку
	cookie := new(http.Cookie)
	cookie.Name = "jwt"                     // Имя куки
	cookie.Value = tokenString               // Значение куки
	cookie.Expires = time.Now().Add(72 * time.Hour) // Время жизни куки
	cookie.Path = "/"                        // Путь, на котором кука доступна
	cookie.HttpOnly = true                   // Ограничение доступа к куке с помощью JavaScript
	cookie.Secure = true                     // Установите true, если используете HTTPS
	cookie.SameSite = http.SameSiteLaxMode      // Политика SameSite

	http.SetCookie(c.Response().Writer, cookie) // Установка куки в ответ

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login successful, token set in cookie",
	})
}
