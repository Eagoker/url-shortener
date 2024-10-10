package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConvertToShort(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name               string
		requestBody       []byte
		expectedStatus    int
		expectedBody      string
	}{
		{
			name:            "Successful conversion",
			requestBody:    []byte("https://example.com/original/url"),
			expectedStatus: http.StatusCreated,
			expectedBody:   "", // Тело ответа будет рандомным, поэтому можно оставить пустым или использовать регулярное выражение
		},
		{
			name:            "Method not allowed",
			requestBody:    nil,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем тестовый запрос
			var req *http.Request
			if tt.requestBody != nil {
				req = httptest.NewRequest(http.MethodPost, "/convert", bytes.NewBuffer(tt.requestBody))
			} else {
				req = httptest.NewRequest(http.MethodGet, "/convert", nil)
			}
			recorder := httptest.NewRecorder()

			ConvertToShort(recorder, req)

			// Проверяем код состояния
			if status := recorder.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			// Если это успешный случай, то проверяем, что тело ответа не пустое
			if tt.expectedStatus == http.StatusCreated && recorder.Body.String() == "" {
				t.Errorf("handler returned empty body for successful conversion")
			}

			// В случае, если хотите проверить, что тело ответа соответствует формату короткого URL
			if tt.expectedStatus == http.StatusCreated {
				if !bytes.HasPrefix(recorder.Body.Bytes(), []byte("http://localhost:8080/")) {
					t.Errorf("handler returned unexpected body: got %v", recorder.Body.String())
				}
			}
		})
	}
}
