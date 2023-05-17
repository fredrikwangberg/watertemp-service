package datasource

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/skyportsystems/testify/assert"
)

func TestColdestTemperatureLocationHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/test", nil)
	response := httptest.NewRecorder()
	ColdestTemperatureLocationHandler(response, request, GetMockedData)

	assert.Equal(t, http.StatusOK, response.Result().StatusCode)

	got := response.Header().Get("Content-Type")
	want := "application/json; charset=utf-8"
	assert.Equal(t, got, want)

	var jsonResponse map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &jsonResponse)
	if err != nil {
		t.Errorf("Response is not valid JSON: %v", err)
	}

}
