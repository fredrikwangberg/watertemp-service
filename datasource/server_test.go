package datasource

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/skyportsystems/testify/assert"
)

// Test LinkHandler for status code, content-type and response body
func TestLinkHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/test", nil)
	response := httptest.NewRecorder()
	mockData := []byte(`{"Temperature": "100", "Location": "Test"}`)
	ColdestTemperatureLocationHandler(response, request, mockData)

	// test if the status code is 200
	assert.Equal(t, http.StatusOK, response.Result().StatusCode)

	// test if the content-type header is set to "application/json"
	got := response.Header().Get("Content-Type")
	want := "application/json; charset=utf-8"
	assert.Equal(t, got, want)

	// test whether response body is valid JSON
	var jsonResponse map[string]string
	err := json.Unmarshal(response.Body.Bytes(), &jsonResponse)
	if err != nil {
		t.Errorf("Response is not valid JSON: %v", err)
	}

}

// Test StartServer for status code
func TestStartServer(t *testing.T) {
	mockData := []byte(`{"Temperature": "100", "Location": "Test"}`)

	go StartServer("/test", mockData)

	response, err := http.Get("http://localhost:8080/test")
	if err != nil {
		t.Fatal(err)
	}

	// test if the response status code is 200
	assert.Equal(t, response.StatusCode, http.StatusOK)
}
