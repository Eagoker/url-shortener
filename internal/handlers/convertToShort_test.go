package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConvertToShort(t *testing.T) {
	tests := []struct {
		name               string
		requestBody       []byte
		expectedStatus    int
		expectedBodyPrefix string
	}{
		{
			name:               "Successful conversion",
			requestBody:       []byte("https://example.com/original/url"),
			expectedStatus:    http.StatusCreated,
			expectedBodyPrefix: "http://localhost:8080/", // Убедитесь, что это корректный префикс
		},
		{
			name:            "Method not allowed",
			requestBody:    nil,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBodyPrefix: "", // Нет тела для метода, не поддерживаемого
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			if tt.requestBody != nil {
				req = httptest.NewRequest(http.MethodPost, "/convert", bytes.NewBuffer(tt.requestBody))
			} else {
				req = httptest.NewRequest(http.MethodGet, "/convert", nil)
			}
			recorder := httptest.NewRecorder()

			ConvertToShort(recorder, req)

			if status := recorder.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			if tt.expectedStatus == http.StatusCreated {
				if !bytes.HasPrefix(recorder.Body.Bytes(), []byte(tt.expectedBodyPrefix)) {
					t.Errorf("handler returned unexpected body: got %v", recorder.Body.String())
				}
			}
		})
	}
}
