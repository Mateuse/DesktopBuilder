package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mateuse/desktop-builder-backend/internal/constants"
	"github.com/mateuse/desktop-builder-backend/internal/models"
	"github.com/mateuse/desktop-builder-backend/internal/routes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// HealthIntegrationTestSuite defines the test suite for health endpoint integration tests
type HealthIntegrationTestSuite struct {
	suite.Suite
	server *httptest.Server
	client *http.Client
}

// SetupSuite runs before all tests in the suite
func (suite *HealthIntegrationTestSuite) SetupSuite() {
	// Create a test server with the actual routes
	mux := http.NewServeMux()
	routes.RegisterHealthRoutes(mux)

	suite.server = httptest.NewServer(mux)
	suite.client = &http.Client{}
}

// TearDownSuite runs after all tests in the suite
func (suite *HealthIntegrationTestSuite) TearDownSuite() {
	suite.server.Close()
}

// TestHealthEndpoint_GET tests the health endpoint with GET method
func (suite *HealthIntegrationTestSuite) TestHealthEndpoint_GET() {
	// Make request to health endpoint
	resp, err := suite.client.Get(suite.server.URL + "/health")
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	// Check status code
	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode)

	// Check content type
	assert.Equal(suite.T(), "application/json", resp.Header.Get("Content-Type"))

	// Parse response body
	var response models.SuccessResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(suite.T(), err)

	// Verify response content
	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Equal(suite.T(), constants.HEALTH_MESSAGE, response.Message)
	assert.Nil(suite.T(), response.Data)
}

// TestHealthEndpoint_POST tests the health endpoint with POST method
func (suite *HealthIntegrationTestSuite) TestHealthEndpoint_POST() {
	// Make POST request to health endpoint
	resp, err := suite.client.Post(suite.server.URL+"/health", "application/json", nil)
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	// Check status code
	assert.Equal(suite.T(), http.StatusMethodNotAllowed, resp.StatusCode)

	// Check content type
	assert.Equal(suite.T(), "application/json", resp.Header.Get("Content-Type"))

	// Parse response body
	var response models.ErrorResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(suite.T(), err)

	// Verify response content
	assert.Equal(suite.T(), http.StatusMethodNotAllowed, response.Code)
	assert.Equal(suite.T(), constants.METHOD_NOT_ALLOWED_MESSAGE, response.Message)
	assert.Nil(suite.T(), response.Data)
}

// TestHealthEndpoint_PUT tests the health endpoint with PUT method
func (suite *HealthIntegrationTestSuite) TestHealthEndpoint_PUT() {
	// Create PUT request
	req, err := http.NewRequest(http.MethodPut, suite.server.URL+"/health", nil)
	assert.NoError(suite.T(), err)

	// Make request
	resp, err := suite.client.Do(req)
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	// Check status code
	assert.Equal(suite.T(), http.StatusMethodNotAllowed, resp.StatusCode)

	// Parse response body
	var response models.ErrorResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(suite.T(), err)

	// Verify response content
	assert.Equal(suite.T(), http.StatusMethodNotAllowed, response.Code)
	assert.Equal(suite.T(), constants.METHOD_NOT_ALLOWED_MESSAGE, response.Message)
}

// TestHealthEndpoint_DELETE tests the health endpoint with DELETE method
func (suite *HealthIntegrationTestSuite) TestHealthEndpoint_DELETE() {
	// Create DELETE request
	req, err := http.NewRequest(http.MethodDelete, suite.server.URL+"/health", nil)
	assert.NoError(suite.T(), err)

	// Make request
	resp, err := suite.client.Do(req)
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	// Check status code
	assert.Equal(suite.T(), http.StatusMethodNotAllowed, resp.StatusCode)

	// Parse response body
	var response models.ErrorResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(suite.T(), err)

	// Verify response content
	assert.Equal(suite.T(), http.StatusMethodNotAllowed, response.Code)
	assert.Equal(suite.T(), constants.METHOD_NOT_ALLOWED_MESSAGE, response.Message)
}

// TestHealthEndpoint_OPTIONS tests the health endpoint with OPTIONS method
func (suite *HealthIntegrationTestSuite) TestHealthEndpoint_OPTIONS() {
	// Create OPTIONS request
	req, err := http.NewRequest(http.MethodOptions, suite.server.URL+"/health", nil)
	assert.NoError(suite.T(), err)

	// Make request
	resp, err := suite.client.Do(req)
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	// Check status code
	assert.Equal(suite.T(), http.StatusMethodNotAllowed, resp.StatusCode)
}

// TestHealthEndpoint_InvalidPath tests invalid paths
func (suite *HealthIntegrationTestSuite) TestHealthEndpoint_InvalidPath() {
	// Make request to invalid path
	resp, err := suite.client.Get(suite.server.URL + "/invalid-path")
	assert.NoError(suite.T(), err)
	defer resp.Body.Close()

	// Should return 404
	assert.Equal(suite.T(), http.StatusNotFound, resp.StatusCode)
}

// TestHealthEndpoint_Concurrent tests concurrent requests
func (suite *HealthIntegrationTestSuite) TestHealthEndpoint_Concurrent() {
	const numRequests = 10
	results := make(chan error, numRequests)

	// Make concurrent requests
	for i := 0; i < numRequests; i++ {
		go func() {
			resp, err := suite.client.Get(suite.server.URL + "/health")
			if err != nil {
				results <- err
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				results <- assert.AnError
				return
			}

			results <- nil
		}()
	}

	// Collect results
	for i := 0; i < numRequests; i++ {
		err := <-results
		assert.NoError(suite.T(), err)
	}
}

// TestHealthIntegrationTestSuite runs the test suite
func TestHealthIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(HealthIntegrationTestSuite))
}
