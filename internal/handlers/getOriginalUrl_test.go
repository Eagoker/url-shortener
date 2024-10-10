package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOriginalUrl(t *testing.T) {
	// Создаем тестовый HTTP GET-запрос
	req, err := http.NewRequest(http.MethodGet, "/short-url-id", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Создаем ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()

	// Вызываем тестируемый хендлер
	handler := http.HandlerFunc(GetOriginalUrl)
	handler.ServeHTTP(rr, req)

	// Проверяем код ответа
	if status := rr.Code; status != http.StatusTemporaryRedirect {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusTemporaryRedirect)
	}

	// Проверяем заголовок Location
	expectedLocation := "https://practicum.yandex.ru/"
	if location := rr.Header().Get("Location"); location != expectedLocation {
		t.Errorf("handler returned wrong Location header: got %v want %v", location, expectedLocation)
	}
}

func TestGetOriginalUrl_InvalidMethod(t *testing.T) {
	// Создаем тестовый HTTP POST-запрос (недопустимый метод)
	req, err := http.NewRequest(http.MethodPost, "/short-url-id", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Создаем ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()

	// Вызываем тестируемый хендлер
	handler := http.HandlerFunc(GetOriginalUrl)
	handler.ServeHTTP(rr, req)

	// Проверяем код ответа для неверного метода
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}

	// Проверяем сообщение об ошибке
	expected := "Method not allowed!\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
