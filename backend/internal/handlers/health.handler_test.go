package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mateuse/desktop-builder-backend/internal/constants"
	"github.com/mateuse/desktop-builder-backend/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandler_GET(t *testing.T) {
	// Create a request to the health endpoint
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	// Call the handler
	HealthHandler(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the content type
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	// Parse the response body
	var response models.SuccessResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)

	// Check the response content
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, constants.HEALTH_MESSAGE, response.Message)
	assert.Nil(t, response.Data)
}

func TestHealthHandler_POST(t *testing.T) {
	// Create a POST request to the health endpoint
	req := httptest.NewRequest(http.MethodPost, "/health", nil)
	w := httptest.NewRecorder()

	// Call the handler
	HealthHandler(w, req)

	// Check the status code
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)

	// Check the content type
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	// Parse the response body
	var response models.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)

	// Check the response content
	assert.Equal(t, http.StatusMethodNotAllowed, response.Code)
	assert.Equal(t, constants.METHOD_NOT_ALLOWED_MESSAGE, response.Message)
	assert.Nil(t, response.Data)
}

func TestHealthHandler_PUT(t *testing.T) {
	// Create a PUT request to the health endpoint
	req := httptest.NewRequest(http.MethodPut, "/health", nil)
	w := httptest.NewRecorder()

	// Call the handler
	HealthHandler(w, req)

	// Check the status code
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)

	// Check the content type
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	// Parse the response body
	var response models.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)

	// Check the response content
	assert.Equal(t, http.StatusMethodNotAllowed, response.Code)
	assert.Equal(t, constants.METHOD_NOT_ALLOWED_MESSAGE, response.Message)
	assert.Nil(t, response.Data)
}

func TestHealthHandler_DELETE(t *testing.T) {
	// Create a DELETE request to the health endpoint
	req := httptest.NewRequest(http.MethodDelete, "/health", nil)
	w := httptest.NewRecorder()

	// Call the handler
	HealthHandler(w, req)

	// Check the status code
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)

	// Check the content type
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	// Parse the response body
	var response models.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)

	// Check the response content
	assert.Equal(t, http.StatusMethodNotAllowed, response.Code)
	assert.Equal(t, constants.METHOD_NOT_ALLOWED_MESSAGE, response.Message)
	assert.Nil(t, response.Data)
}

// BenchmarkHealthHandler benchmarks the health handler performance
func BenchmarkHealthHandler(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		HealthHandler(w, req)
	}
}
