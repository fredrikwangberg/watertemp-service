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

	gotStatus := response.Result().StatusCode
	wantStatus := http.StatusOK
	assert.Equal(t, gotStatus, wantStatus)

	gotHeader := response.Header().Get("Content-Type")
	wantHeader := "application/json; charset=utf-8"
	assert.Equal(t, gotHeader, wantHeader)

	var jsonResponse map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &jsonResponse)
	if err != nil {
		t.Errorf("Response is not valid JSON: %v", err)
	}

	gotColdestTemperatureLocation := response.Body.String()
	wantColdestTemperatureLocation := `{"Temperature":7.8,"Location":"Eklundsnäsbadet"}`
	assert.Equal(t, gotColdestTemperatureLocation, wantColdestTemperatureLocation)

}
