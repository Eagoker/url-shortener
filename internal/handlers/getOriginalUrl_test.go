package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetOriginalUrl(t *testing.T) {
	// Инициализируем Echo
	e := echo.New()

	// Создаем тестовый HTTP GET-запрос
	req := httptest.NewRequest(http.MethodGet, "/short-url-id", nil)
	rec := httptest.NewRecorder()

	// Создаем контекст с запросом и ответом
	c := e.NewContext(req, rec)

	// Вызываем тестируемый хендлер
	if assert.NoError(t, GetOriginalUrl(c)) {
		// Проверяем код ответа
		assert.Equal(t, http.StatusMovedPermanently, rec.Code)

		// Проверяем заголовок Location
		expectedLocation := "https://practicum.yandex.ru/"
		assert.Equal(t, expectedLocation, rec.Header().Get("Location"))
	}
}

func TestGetOriginalUrl_InvalidMethod(t *testing.T) {
	// Инициализируем Echo
	e := echo.New()

	// Создаем тестовый HTTP POST-запрос (недопустимый метод)
	req := httptest.NewRequest(http.MethodPost, "/short-url-id", nil)
	rec := httptest.NewRecorder()

	// Создаем контекст с запросом и ответом
	c := e.NewContext(req, rec)

	// Вызываем тестируемый хендлер
	err := GetOriginalUrl(c)

	// Проверяем, что вернулась ошибка
	if assert.Error(t, err) {
		he, ok := err.(*echo.HTTPError)
		if ok {
			// Проверяем код ответа для неверного метода
			assert.Equal(t, http.StatusMethodNotAllowed, he.Code)

			// Проверяем сообщение об ошибке
			expected := "Method not allowed!"
			assert.Equal(t, expected, he.Message)
		}
	}
}
