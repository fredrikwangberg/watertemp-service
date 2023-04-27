package datasource

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func doGet(apiUrl string) (*http.Response, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	r, err := client.Get(apiUrl)

	if err != nil {
		return nil, err
	}
	return r, nil
}

func parseResponseBody(r []byte) ([]WaterTemperatureSchema, error) {
	var data []WaterTemperatureSchema
	err := json.Unmarshal([]byte(r), &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetDataFromDataSourceFullJson(apiUrl string) ([]WaterTemperatureSchema, error) {
	response, err := doGet(apiUrl)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	data, err := parseResponseBody(body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetMockedData() ([]WaterTemperatureSchema, error) {
	mockedJson := `[{"type":"Watertemp","temp_water":7.8,"formatted_time":"Nov 02 2021 20:10:14","alias":"Eklundsnäsbadet","ts":1635880214580,"latitude":59.16883,"longitude":17.59184,"gmap":"https://www.google.com/maps/search/?api=1&query=59.16883,17.59184"},{"type":"Watertemp","temp_water":8.9,"formatted_time":"Nov 02 2021 20:10:02","alias":"Åbynäsbadet","ts":1635880202745,"latitude":59.018397,"longitude":17.619576,"gmap":"https://www.google.com/maps/search/?api=1&query=59.018397,17.619576"},{"type":"Watertemp","temp_water":7.9,"formatted_time":"Nov 02 2021 20:08:02","alias":"Bergabadet","ts":1635880082291,"latitude":59.057008,"longitude":17.440774,"gmap":"https://www.google.com/maps/search/?api=1&query=59.057008,17.440774"},{"type":"Watertemp","temp_water":9.8,"formatted_time":"Nov 02 2021 19:54:07","alias":"Mälarbadet","ts":1635879247710,"latitude":59.222657,"longitude":17.611886,"gmap":"https://www.google.com/maps/search/?api=1&query=59.222657,17.611886"},{"type":"Watertemp","temp_water":9.8,"formatted_time":"Nov 02 2021 19:35:02","alias":"Bränningestrand","ts":1635878102976,"latitude":59.148617,"longitude":17.6674,"gmap":"https://www.google.com/maps/search/?api=1&query=59.148617,17.6674"},{"type":"Watertemp","temp_water":10,"formatted_time":"Nov 02 2021 19:27:34","alias":"Underåsbadet","ts":1635877654797,"latitude":59.26482,"longitude":17.536534,"gmap":"https://www.google.com/maps/search/?api=1&query=59.26482,17.536534"},{"type":"Watertemp","temp_water":8.4,"formatted_time":"Nov 02 2021 18:18:37","alias":"Nya Malmsjöbadet","ts":1635873517532,"latitude":59.234823,"longitude":17.536534,"gmap":"https://www.google.com/maps/search/?api=1&query=59.234823,17.536534"},{"type":"Watertemp","temp_water":10.2,"formatted_time":"Oct 22 2021 12:17:51","alias":"Farstanäsbadet","ts":1634897871476,"latitude":59.096884,"longitude":17.65387,"gmap":"https://www.google.com/maps/search/?api=1&query=59.096884,17.65387"},{"type":"Watertemp","temp_water":11.3,"formatted_time":"Oct 15 2021 13:40:31","alias":"Näsets udde(Glashyttan)","ts":1634298031991,"latitude":59.158419,"longitude":17.66072,"gmap":"https://www.google.com/maps/search/?api=1&query=59.158419,17.66072"},{"type":"Watertemp","temp_water":23.9,"formatted_time":"Jul 08 2021 12:27:23","alias":"Eklundsnäsbadet","ts":1625740043042,"latitude":59.16883,"longitude":17.59184,"gmap":"https://www.google.com/maps/search/?api=1&query=59.16883,17.59184"},{"type":"Watertemp","temp_water":18.1,"formatted_time":"Jun 19 2021 03:06:41","alias":"Åbynäsbadet","ts":1624064801077,"latitude":59.018397,"longitude":17.619576,"gmap":"https://www.google.com/maps/search/?api=1&query=59.018397,17.619576"}]`
	data, err := parseResponseBody([]byte(mockedJson))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetLocationWithLowestTemperature(locations []WaterTemperatureSchema) WaterTemperatureSchema {
	var lowestTemperatureCandidateIndex = 1
	for i := range locations {
		if hasLowerTemperature(locations[i], locations[lowestTemperatureCandidateIndex]) {
			lowestTemperatureCandidateIndex = i
		}
	}
	return locations[lowestTemperatureCandidateIndex]
}

func WaterTemperatureSchemaToString(wts WaterTemperatureSchema) string {
	return fmt.Sprintf("%+v\n", wts)
}

func hasLowerTemperature(location1 WaterTemperatureSchema, location2 WaterTemperatureSchema) bool {
	return location1.Temp_water < location2.Temp_water
}

/* TO DO: Need to discuss this part of the code!
The task was to "Create function that drops data from watertemp schema and converts the data
to a json object (name of location and temperature) - create own class"
Is it a reasonable solution?
Is it correctly located or should for example ColdSiteStruct be located elsewhere? */

// Define a struct type ColdSiteStruct
type ColdSiteStruct struct {
	TempWater float32 `json:"Temp_water"`
	Alias     string  `json:"Alias"`
}

// Define a function named GetJsonColdSite that takes an input parameter data of type []WaterTemperatureSchema, and returns a byte slice and an error
func GetJsonColdSite(data []WaterTemperatureSchema) ([]byte, error) {

	// Call GetLocationWithLowestTemperature function to retrieve the location with the lowest temperature from the input data
	coldSite := GetLocationWithLowestTemperature(data)

	// Create a new ColdSiteStruct instance with fields set to the values of the lowest temperature location
	coldSiteStruct := ColdSiteStruct{
		TempWater: coldSite.Temp_water,
		Alias:     coldSite.Alias,
	}

	// Convert the coldSiteStruct instance to a JSON byte array
	jsonColdSite, err := json.Marshal(coldSiteStruct) // If there is an error during the marshaling process, return nil and the error
	if err != nil {
		return nil, err

	}

	return jsonColdSite, nil // Otherwise, return the JSON byte array and nil
}

/* END TO DO */
