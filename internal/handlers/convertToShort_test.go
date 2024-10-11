package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestConvertToShort(t *testing.T) {
	tests := []struct {
		name               string
		requestBody        []byte
		expectedStatus     int
		expectedBodyPrefix string
		method             string
	}{
		{
			name:               "Successful conversion",
			requestBody:        []byte("https://example.com/original/url"),
			expectedStatus:     http.StatusCreated,
			method:             http.MethodPost,
		},
		{
			name:               "Method not allowed",
			requestBody:        nil,
			expectedStatus:     http.StatusMethodNotAllowed,
			method:             http.MethodGet,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем Echo-инстанс
			e := echo.New()

			// Создаем новый запрос и записывающий ответ
			req := httptest.NewRequest(tt.method, "/convert", bytes.NewBuffer(tt.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Вызываем обработчик
			err := ConvertToShort(c)
			if err != nil {
				// Если произошла ошибка, проверяем ее статус
				he, ok := err.(*echo.HTTPError)
				if ok {
					assert.Equal(t, tt.expectedStatus, he.Code)
				} else {
					t.Fatalf("unexpected error: %v", err)
				}
			} else {
				// Проверяем статус ответа
				assert.Equal(t, tt.expectedStatus, rec.Code)
			}
		})
	}
}
