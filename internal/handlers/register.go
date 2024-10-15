package handlers

import (
	"context"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Register(c echo.Context) error {
	req := new(RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to hash password",
		})
	}

	// Сохранение пользователя в базе данных
	_, err = h.db.Exec(context.Background(), `
		INSERT INTO users (username, password_hash) VALUES ($1, $2)
	`, req.Username, string(hashedPassword))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to register user",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "User registered successfully",
	})
}
