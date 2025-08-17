package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mateuse/desktop-builder-backend/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestWriteJSON(t *testing.T) {
	tests := []struct {
		name     string
		status   int
		payload  interface{}
		expected string
	}{
		{
			name:     "Write simple string",
			status:   http.StatusOK,
			payload:  "test message",
			expected: `"test message"`,
		},
		{
			name:   "Write struct",
			status: http.StatusCreated,
			payload: struct {
				Message string `json:"message"`
				Count   int    `json:"count"`
			}{
				Message: "test",
				Count:   5,
			},
			expected: `{"message":"test","count":5}`,
		},
		{
			name:     "Write nil",
			status:   http.StatusNoContent,
			payload:  nil,
			expected: `null`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			WriteJSON(w, tt.status, tt.payload)

			// Check status code
			assert.Equal(t, tt.status, w.Code)

			// Check content type
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

			// Check body content (trim newline that json.NewEncoder adds)
			body := w.Body.String()
			if len(body) > 0 && body[len(body)-1] == '\n' {
				body = body[:len(body)-1]
			}
			assert.JSONEq(t, tt.expected, body)
		})
	}
}

func TestWriteSuccess(t *testing.T) {
	tests := []struct {
		name    string
		status  int
		message string
		data    interface{}
	}{
		{
			name:    "Success with nil data",
			status:  http.StatusOK,
			message: "Operation successful",
			data:    nil,
		},
		{
			name:    "Success with string data",
			status:  http.StatusCreated,
			message: "Resource created",
			data:    "resource-id-123",
		},
		{
			name:    "Success with map data",
			status:  http.StatusOK,
			message: "Data retrieved",
			data:    map[string]interface{}{"id": float64(1), "name": "test"}, // JSON unmarshals numbers as float64
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			WriteSuccess(w, tt.status, tt.message, tt.data)

			// Check status code
			assert.Equal(t, tt.status, w.Code)

			// Check content type
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

			// Parse response
			var response models.SuccessResponse
			err := json.NewDecoder(w.Body).Decode(&response)
			assert.NoError(t, err)

			// Check response fields
			assert.Equal(t, tt.status, response.Code)
			assert.Equal(t, tt.message, response.Message)
			assert.Equal(t, tt.data, response.Data)
		})
	}
}

func TestWriteError(t *testing.T) {
	tests := []struct {
		name    string
		status  int
		message string
		data    interface{}
	}{
		{
			name:    "Error with nil data",
			status:  http.StatusBadRequest,
			message: "Invalid request",
			data:    nil,
		},
		{
			name:    "Error with string data",
			status:  http.StatusNotFound,
			message: "Resource not found",
			data:    "resource-id-123",
		},
		{
			name:    "Error with error details",
			status:  http.StatusInternalServerError,
			message: "Internal server error",
			data:    map[string]interface{}{"details": "Database connection failed"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			WriteError(w, tt.status, tt.message, tt.data)

			// Check status code
			assert.Equal(t, tt.status, w.Code)

			// Check content type
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

			// Parse response
			var response models.ErrorResponse
			err := json.NewDecoder(w.Body).Decode(&response)
			assert.NoError(t, err)

			// Check response fields
			assert.Equal(t, tt.status, response.Code)
			assert.Equal(t, tt.message, response.Message)
			assert.Equal(t, tt.data, response.Data)
		})
	}
}

// BenchmarkWriteJSON benchmarks the WriteJSON function
func BenchmarkWriteJSON(b *testing.B) {
	payload := map[string]interface{}{
		"message": "test message",
		"data":    []int{1, 2, 3, 4, 5},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		WriteJSON(w, http.StatusOK, payload)
	}
}

// BenchmarkWriteSuccess benchmarks the WriteSuccess function
func BenchmarkWriteSuccess(b *testing.B) {
	data := map[string]interface{}{
		"id":   123,
		"name": "test resource",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		WriteSuccess(w, http.StatusOK, "Success", data)
	}
}
