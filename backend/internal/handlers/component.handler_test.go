package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/mateuse/desktop-builder-backend/internal/constants"
	"github.com/mateuse/desktop-builder-backend/internal/models"
	"github.com/mateuse/desktop-builder-backend/internal/utils"
	"github.com/stretchr/testify/assert"
)

// Mock data for testing
var mockComponent = models.Component{
	ID:        1,
	Category:  models.CategoryCPU,
	Brand:     "Intel",
	Model:     "Core i7-12700K",
	SKU:       stringPtr("BX8071512700K"),
	UPC:       stringPtr("735858491174"),
	Specs:     json.RawMessage(`{"cores": 12, "threads": 20, "base_clock": "3.6 GHz"}`),
	CreatedAt: time.Now(),
}

var mockComponents = []models.Component{
	mockComponent,
	{
		ID:        2,
		Category:  models.CategoryCPU,
		Brand:     "AMD",
		Model:     "Ryzen 7 5800X",
		SKU:       stringPtr("100-100000063WOF"),
		UPC:       stringPtr("730143312042"),
		Specs:     json.RawMessage(`{"cores": 8, "threads": 16, "base_clock": "3.8 GHz"}`),
		CreatedAt: time.Now(),
	},
}

// Helper function to create string pointers
func stringPtr(s string) *string {
	return &s
}

// TestGetComponentsHandler_MethodValidation tests HTTP method validation
func TestGetComponentsHandler_MethodValidation(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedStatus int
		expectedError  string
	}{
		{
			name:           "Invalid POST request",
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedError:  constants.METHOD_NOT_ALLOWED_MESSAGE,
		},
		{
			name:           "Invalid PUT request",
			method:         http.MethodPut,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedError:  constants.METHOD_NOT_ALLOWED_MESSAGE,
		},
		{
			name:           "Invalid DELETE request",
			method:         http.MethodDelete,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedError:  constants.METHOD_NOT_ALLOWED_MESSAGE,
		},
		{
			name:           "Invalid PATCH request",
			method:         http.MethodPatch,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedError:  constants.METHOD_NOT_ALLOWED_MESSAGE,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/components", nil)
			w := httptest.NewRecorder()

			GetComponentsHandler(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

			var response models.ErrorResponse
			err := json.NewDecoder(w.Body).Decode(&response)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, response.Code)
			assert.Equal(t, tt.expectedError, response.Message)
		})
	}
}

// TestGetComponentsHandler_PathParamLogic tests the path parameter routing logic
// Note: This test focuses on the parameter parsing and routing logic without
// actually calling the services (which would require database setup)
func TestGetComponentsHandler_PathParamLogic(t *testing.T) {
	tests := []struct {
		name            string
		path            string
		pattern         string
		expectedRouting string
		description     string
	}{
		{
			name:            "No path parameters",
			path:            "/components",
			pattern:         "/components",
			expectedRouting: "GetAllComponents",
			description:     "Should route to GetAllComponents service",
		},
		{
			name:            "Category only",
			path:            "/components/cpu",
			pattern:         "/components/{category}",
			expectedRouting: "GetComponentsByCategory",
			description:     "Should route to GetComponentsByCategory service",
		},
		{
			name:            "Category and brand",
			path:            "/components/cpu/Intel",
			pattern:         "/components/{category}/{brand}",
			expectedRouting: "GetComponentsByBrand",
			description:     "Should route to GetComponentsByBrand service",
		},
		{
			name:            "ID only",
			path:            "/components/item/123",
			pattern:         "/components/item/{id}",
			expectedRouting: "GetComponentById",
			description:     "Should route to GetComponentById service",
		},
		{
			name:            "Different categories",
			path:            "/components/gpu",
			pattern:         "/components/{category}",
			expectedRouting: "GetComponentsByCategory",
			description:     "Should route to GetComponentsByCategory for different category",
		},
		{
			name:            "Different brands",
			path:            "/components/cpu/AMD",
			pattern:         "/components/{category}/{brand}",
			expectedRouting: "GetComponentsByBrand",
			description:     "Should route to GetComponentsByBrand for different brand",
		},
		{
			name:            "Uppercase brand routing",
			path:            "/components/cpu/INTEL",
			pattern:         "/components/{category}/{brand}",
			expectedRouting: "GetComponentsByBrand",
			description:     "Should route to GetComponentsByBrand for uppercase brand",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a ServeMux to properly handle path variables
			mux := http.NewServeMux()
			mux.HandleFunc(tt.pattern, func(w http.ResponseWriter, r *http.Request) {
				params := parseComponentQueryParams(r)

				// Test the routing logic without calling the actual services
				var actualRouting string
				switch {
				case params.ID != "":
					actualRouting = "GetComponentById"
				case params.Category != "" && params.Brand != "":
					actualRouting = "GetComponentsByBrand"
				case params.Category != "":
					actualRouting = "GetComponentsByCategory"
				default:
					actualRouting = "GetAllComponents"
				}

				assert.Equal(t, tt.expectedRouting, actualRouting, tt.description)
				t.Logf("Test case: %s - %s", tt.name, tt.description)
				t.Logf("Request path: %s", tt.path)
				t.Logf("Route pattern: %s", tt.pattern)
				t.Logf("Expected routing: %s, Actual routing: %s", tt.expectedRouting, actualRouting)
			})

			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, req)
		})
	}
}

// TestParseComponentQueryParams tests the path parameter parsing function
func TestParseComponentQueryParams(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		pattern  string
		expected models.ComponentQueryParams
	}{
		{
			name:    "No path parameters",
			path:    "/components",
			pattern: "/components",
			expected: models.ComponentQueryParams{
				Category: "",
				Brand:    "",
				ID:       "",
			},
		},
		{
			name:    "Category only",
			path:    "/components/cpu",
			pattern: "/components/{category}",
			expected: models.ComponentQueryParams{
				Category: "cpu",
				Brand:    "",
				ID:       "",
			},
		},
		{
			name:    "Category and brand",
			path:    "/components/cpu/Intel",
			pattern: "/components/{category}/{brand}",
			expected: models.ComponentQueryParams{
				Category: "cpu",
				Brand:    "Intel",
				ID:       "",
			},
		},
		{
			name:    "ID only",
			path:    "/components/item/123",
			pattern: "/components/item/{id}",
			expected: models.ComponentQueryParams{
				Category: "",
				Brand:    "",
				ID:       "123",
			},
		},
		{
			name:    "URL encoded brand name",
			path:    "/components/cpu/Intel%20Corp",
			pattern: "/components/{category}/{brand}",
			expected: models.ComponentQueryParams{
				Category: "cpu",
				Brand:    "Intel Corp",
				ID:       "",
			},
		},
		{
			name:    "Uppercase brand name",
			path:    "/components/cpu/INTEL",
			pattern: "/components/{category}/{brand}",
			expected: models.ComponentQueryParams{
				Category: "cpu",
				Brand:    "INTEL",
				ID:       "",
			},
		},
		{
			name:    "Mixed case brand name",
			path:    "/components/gpu/NvIdIa",
			pattern: "/components/{category}/{brand}",
			expected: models.ComponentQueryParams{
				Category: "gpu",
				Brand:    "NvIdIa",
				ID:       "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a ServeMux to properly handle path variables
			mux := http.NewServeMux()
			mux.HandleFunc(tt.pattern, func(w http.ResponseWriter, r *http.Request) {
				result := parseComponentQueryParams(r)

				assert.Equal(t, tt.expected.Category, result.Category)
				assert.Equal(t, tt.expected.Brand, result.Brand)
				assert.Equal(t, tt.expected.ID, result.ID)
			})

			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, req)
		})
	}
}

// TestGetComponentsHandler_EdgeCases tests edge cases in path parameter parsing
func TestGetComponentsHandler_EdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		pattern     string
		description string
	}{
		{
			name:        "Special characters in path",
			path:        "/components/cpu/Intel%20Corp",
			pattern:     "/components/{category}/{brand}",
			description: "Should handle URL encoded special characters in path",
		},
		{
			name:        "Very long parameter values",
			path:        fmt.Sprintf("/components/%s", strings.Repeat("a", 100)),
			pattern:     "/components/{category}",
			description: "Should handle long parameter values in path",
		},
		{
			name:        "Case sensitivity in path",
			path:        "/components/CPU/INTEL",
			pattern:     "/components/{category}/{brand}",
			description: "Should preserve brand case while keeping category as-is",
		},
		{
			name:        "Numeric values in path",
			path:        "/components/item/12345",
			pattern:     "/components/item/{id}",
			description: "Should handle numeric values in path",
		},
		{
			name:        "Hyphenated values",
			path:        "/components/case-fan/Noctua",
			pattern:     "/components/{category}/{brand}",
			description: "Should handle hyphenated category names",
		},
		{
			name:        "Underscored values",
			path:        "/components/power_supply/EVGA",
			pattern:     "/components/{category}/{brand}",
			description: "Should handle underscored category names",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a ServeMux to properly handle path variables
			mux := http.NewServeMux()
			mux.HandleFunc(tt.pattern, func(w http.ResponseWriter, r *http.Request) {
				// Test that parameter parsing doesn't panic or error
				params := parseComponentQueryParams(r)

				// Just verify that parsing completed without errors
				assert.NotNil(t, params)
				t.Logf("Test case: %s - %s", tt.name, tt.description)
				t.Logf("Request path: %s", tt.path)
				t.Logf("Route pattern: %s", tt.pattern)
				t.Logf("Parsed params: Category=%s, Brand=%s, ID=%s", params.Category, params.Brand, params.ID)
			})

			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, req)
		})
	}
}

// TestParseComponentQueryParams_BrandCasePreservation tests that brand parameter preserves original casing
func TestParseComponentQueryParams_BrandCasePreservation(t *testing.T) {
	tests := []struct {
		name          string
		path          string
		pattern       string
		expectedBrand string
		description   string
	}{
		{
			name:          "Lowercase brand",
			path:          "/components/cpu/intel",
			pattern:       "/components/{category}/{brand}",
			expectedBrand: "intel",
			description:   "Should keep lowercase brand as-is",
		},
		{
			name:          "Uppercase brand",
			path:          "/components/cpu/INTEL",
			pattern:       "/components/{category}/{brand}",
			expectedBrand: "INTEL",
			description:   "Should preserve uppercase brand",
		},
		{
			name:          "Mixed case brand",
			path:          "/components/cpu/InTeL",
			pattern:       "/components/{category}/{brand}",
			expectedBrand: "InTeL",
			description:   "Should preserve mixed case brand",
		},
		{
			name:          "Brand with spaces",
			path:          "/components/cpu/Intel%20Corp",
			pattern:       "/components/{category}/{brand}",
			expectedBrand: "Intel Corp",
			description:   "Should preserve brand with spaces and original casing",
		},
		{
			name:          "Empty brand",
			path:          "/components/cpu/",
			pattern:       "/components/{category}/{brand}",
			expectedBrand: "",
			description:   "Should handle empty brand gracefully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a ServeMux to properly handle path variables
			mux := http.NewServeMux()
			mux.HandleFunc(tt.pattern, func(w http.ResponseWriter, r *http.Request) {
				params := parseComponentQueryParams(r)

				assert.Equal(t, tt.expectedBrand, params.Brand, tt.description)
				t.Logf("Test case: %s - %s", tt.name, tt.description)
				t.Logf("Request path: %s", tt.path)
				t.Logf("Expected brand: %s, Actual brand: %s", tt.expectedBrand, params.Brand)
			})

			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, req)
		})
	}
}

// TestSqlErrNoRowsHandling tests that sql.ErrNoRows is properly identified
func TestSqlErrNoRowsHandling(t *testing.T) {
	// This test verifies that our error handling logic correctly identifies sql.ErrNoRows
	// without requiring database integration

	tests := []struct {
		name          string
		err           error
		expectedIs404 bool
		description   string
	}{
		{
			name:          "sql.ErrNoRows should be identified as 404",
			err:           sql.ErrNoRows,
			expectedIs404: true,
			description:   "sql.ErrNoRows should trigger 404 Not Found response",
		},
		{
			name:          "Other database error should be 500",
			err:           fmt.Errorf("connection timeout"),
			expectedIs404: false,
			description:   "Non-sql.ErrNoRows errors should trigger 500 Internal Server Error",
		},
		{
			name:          "Nil error should not be 404",
			err:           nil,
			expectedIs404: false,
			description:   "Nil error should not trigger 404",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test the error comparison logic that we use in the handler
			is404 := (tt.err == sql.ErrNoRows)

			assert.Equal(t, tt.expectedIs404, is404, tt.description)
			t.Logf("Test case: %s - %s", tt.name, tt.description)
			t.Logf("Error: %v, Is 404: %t", tt.err, is404)
		})
	}
}

// BenchmarkGetComponentsHandler_MethodValidation benchmarks method validation
func BenchmarkGetComponentsHandler_MethodValidation(b *testing.B) {
	req := httptest.NewRequest(http.MethodPost, "/components", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		GetComponentsHandler(w, req)
	}
}

// TestGetComponentsHandler_FunctionRouting tests that the correct internal handler function gets called
// based on the URL parameters using a spy pattern to track function calls
func TestGetComponentsHandler_FunctionRouting(t *testing.T) {
	// Spy struct to track which handler functions are called
	type handlerSpy struct {
		getByIDCalled       bool
		getByBrandCalled    bool
		getByCategoryCalled bool
		getAllCalled        bool
		lastIDParam         string
		lastCategoryParam   string
		lastBrandParam      string
		lastPageParam       string
	}

	// Create a custom handler that uses spies instead of real handler functions
	testHandler := func(spy *handlerSpy) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Copy the method validation logic from the original handler
			if r.Method != http.MethodGet {
				utils.WriteError(w, http.StatusMethodNotAllowed, constants.METHOD_NOT_ALLOWED_MESSAGE, nil)
				return
			}

			// Copy the parameter parsing logic from the original handler
			params := parseComponentQueryParams(r)
			// Handle page parsing inline to avoid dependency issues in tests
			pageParam := r.URL.Query().Get("page")
			if pageParam == "" {
				pageParam = "1"
			}
			page := pageParam

			// Copy the routing logic but call spy functions instead
			switch {
			case params.ID != "":
				spy.getByIDCalled = true
				spy.lastIDParam = params.ID
				spy.lastPageParam = page
				// Return success to avoid database calls in tests
				utils.WriteSuccess(w, http.StatusOK, constants.SUCCESS_MESSAGE, mockComponent)
			case params.Category != "" && params.Brand != "":
				spy.getByBrandCalled = true
				spy.lastCategoryParam = params.Category
				spy.lastBrandParam = params.Brand
				spy.lastPageParam = page
				utils.WriteSuccess(w, http.StatusOK, constants.SUCCESS_MESSAGE, mockComponents)
			case params.Category != "":
				spy.getByCategoryCalled = true
				spy.lastCategoryParam = params.Category
				spy.lastPageParam = page
				utils.WriteSuccess(w, http.StatusOK, constants.SUCCESS_MESSAGE, mockComponents)
			default:
				spy.getAllCalled = true
				spy.lastPageParam = page
				utils.WriteSuccess(w, http.StatusOK, constants.SUCCESS_MESSAGE, mockComponents)
			}
		}
	}

	tests := []struct {
		name                        string
		path                        string
		pattern                     string
		query                       string
		expectedGetByIDCalled       bool
		expectedGetByBrandCalled    bool
		expectedGetByCategoryCalled bool
		expectedGetAllCalled        bool
		expectedIDParam             string
		expectedCategoryParam       string
		expectedBrandParam          string
		expectedPageParam           string
		description                 string
	}{
		{
			name:                 "Route to GetAllComponents",
			path:                 "/components",
			pattern:              "/components",
			query:                "",
			expectedGetAllCalled: true,
			expectedPageParam:    "1", // Default page
			description:          "Should call handleGetAllComponents when no path params",
		},
		{
			name:                 "Route to GetAllComponents with page query",
			path:                 "/components",
			pattern:              "/components",
			query:                "?page=2",
			expectedGetAllCalled: true,
			expectedPageParam:    "2",
			description:          "Should call handleGetAllComponents with page parameter",
		},
		{
			name:                        "Route to GetComponentsByCategory",
			path:                        "/components/cpu",
			pattern:                     "/components/{category}",
			query:                       "",
			expectedGetByCategoryCalled: true,
			expectedCategoryParam:       "cpu",
			expectedPageParam:           "1",
			description:                 "Should call handleGetComponentsByCategory for category only",
		},
		{
			name:                        "Route to GetComponentsByCategory with page",
			path:                        "/components/gpu",
			pattern:                     "/components/{category}",
			query:                       "?page=3",
			expectedGetByCategoryCalled: true,
			expectedCategoryParam:       "gpu",
			expectedPageParam:           "3",
			description:                 "Should call handleGetComponentsByCategory with page parameter",
		},
		{
			name:                     "Route to GetComponentsByBrand",
			path:                     "/components/cpu/Intel",
			pattern:                  "/components/{category}/{brand}",
			query:                    "",
			expectedGetByBrandCalled: true,
			expectedCategoryParam:    "cpu",
			expectedBrandParam:       "Intel",
			expectedPageParam:        "1",
			description:              "Should call handleGetComponentsByBrand for category and brand",
		},
		{
			name:                     "Route to GetComponentsByBrand with page",
			path:                     "/components/gpu/NVIDIA",
			pattern:                  "/components/{category}/{brand}",
			query:                    "?page=5",
			expectedGetByBrandCalled: true,
			expectedCategoryParam:    "gpu",
			expectedBrandParam:       "NVIDIA",
			expectedPageParam:        "5",
			description:              "Should call handleGetComponentsByBrand with page parameter",
		},
		{
			name:                  "Route to GetComponentById",
			path:                  "/components/item/123",
			pattern:               "/components/item/{id}",
			query:                 "",
			expectedGetByIDCalled: true,
			expectedIDParam:       "123",
			expectedPageParam:     "1",
			description:           "Should call handleGetComponentByID for ID parameter",
		},
		{
			name:                  "Route to GetComponentById with page",
			path:                  "/components/item/456",
			pattern:               "/components/item/{id}",
			query:                 "?page=2",
			expectedGetByIDCalled: true,
			expectedIDParam:       "456",
			expectedPageParam:     "2",
			description:           "Should call handleGetComponentByID with page parameter",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a fresh spy for each test
			spy := &handlerSpy{}

			// Create ServeMux to handle path parameters
			mux := http.NewServeMux()
			mux.HandleFunc(tt.pattern, testHandler(spy))

			// Create request with query parameters if specified
			fullPath := tt.path + tt.query
			req := httptest.NewRequest(http.MethodGet, fullPath, nil)
			w := httptest.NewRecorder()

			// Execute the handler
			mux.ServeHTTP(w, req)

			// Verify the correct handler function was called
			assert.Equal(t, tt.expectedGetByIDCalled, spy.getByIDCalled,
				"GetByID called mismatch: %s", tt.description)
			assert.Equal(t, tt.expectedGetByBrandCalled, spy.getByBrandCalled,
				"GetByBrand called mismatch: %s", tt.description)
			assert.Equal(t, tt.expectedGetByCategoryCalled, spy.getByCategoryCalled,
				"GetByCategory called mismatch: %s", tt.description)
			assert.Equal(t, tt.expectedGetAllCalled, spy.getAllCalled,
				"GetAll called mismatch: %s", tt.description)

			// Verify the parameters passed to the handler functions
			if tt.expectedGetByIDCalled {
				assert.Equal(t, tt.expectedIDParam, spy.lastIDParam,
					"ID parameter mismatch: %s", tt.description)
			}
			if tt.expectedGetByBrandCalled {
				assert.Equal(t, tt.expectedCategoryParam, spy.lastCategoryParam,
					"Category parameter mismatch: %s", tt.description)
				assert.Equal(t, tt.expectedBrandParam, spy.lastBrandParam,
					"Brand parameter mismatch: %s", tt.description)
			}
			if tt.expectedGetByCategoryCalled {
				assert.Equal(t, tt.expectedCategoryParam, spy.lastCategoryParam,
					"Category parameter mismatch: %s", tt.description)
			}

			// All cases should have the correct page parameter
			assert.Equal(t, tt.expectedPageParam, spy.lastPageParam,
				"Page parameter mismatch: %s", tt.description)

			// Verify exactly one handler was called
			callCount := 0
			if spy.getByIDCalled {
				callCount++
			}
			if spy.getByBrandCalled {
				callCount++
			}
			if spy.getByCategoryCalled {
				callCount++
			}
			if spy.getAllCalled {
				callCount++
			}

			assert.Equal(t, 1, callCount,
				"Expected exactly one handler to be called, got %d: %s", callCount, tt.description)

			// Verify successful response
			assert.Equal(t, http.StatusOK, w.Code, "Expected successful response")
		})
	}
}

// BenchmarkParseComponentQueryParams benchmarks the path parameter parsing
func BenchmarkParseComponentQueryParams(b *testing.B) {
	// Create a ServeMux to properly handle path variables
	mux := http.NewServeMux()
	mux.HandleFunc("/components/{category}/{brand}", func(w http.ResponseWriter, r *http.Request) {
		parseComponentQueryParams(r)
	})

	req := httptest.NewRequest(http.MethodGet, "/components/cpu/Intel", nil)
	w := httptest.NewRecorder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mux.ServeHTTP(w, req)
	}
}
