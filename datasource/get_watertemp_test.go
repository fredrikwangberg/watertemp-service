package datasource

import (
	"github.com/skyportsystems/testify/assert"
	"testing"
)

// TODO, write tests on parseing of json
//   Malformed json
// 	 No response

func TestParsing(t *testing.T) {
	testJson := `[{"type":"Watertemp","temp_water":7.8,"formatted_time":"Nov 02 2021 20:10:14","alias":"Site1","ts":1635880214580,"latitude":59.16883,"longitude":17.59184,"gmap":"link"}]`
	data, _ := parseResponseBody([]byte(testJson))
	assert.Equal(t, data[0].Alias, "Site1")
}

func TestCompareLocationTemperatures(t *testing.T) {
	location1 := WaterTemperatureSchema{
		Temp_water: 6,
		Alias:      "Warmest",
	}
	location2 := WaterTemperatureSchema{
		Temp_water: 5,
		Alias:      "Warmest",
	}
	if hasLowerTemperature(location1, location2) {
		t.Fatalf("location2 has lowest temperature, not location1")
	}
}
