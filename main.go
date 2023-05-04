package main

import (
	"fmt"
	"net/http"

	"github.com/fredrikwangberg/watertemp/datasource"
)

func main() {
	data, err := getData(false)
	if err == nil {
		coldest := datasource.GetLocationWithLowestTemperature(data)
		coldestJson, err := datasource.GetTemperatureLocationJson(coldest)

		if err == nil {
			fmt.Println("Location with the lowest temperature:")
			fmt.Println(datasource.WaterTemperatureToString(coldest))
			fmt.Println(string(coldestJson))

			http.HandleFunc("/coldest", func(w http.ResponseWriter, r *http.Request) {
				datasource.LinkHandler(w, r, coldestJson)
			})

			fmt.Println("Check the response from the web server with http://localhost:8080/coldest")
			http.ListenAndServe(":8080", nil)
		} else {
			fmt.Println("Error getting JSON object: ", err)
		}
	}
}

func getData(useMocked bool) ([]datasource.WaterTemperature, error) {
	if useMocked {
		fmt.Println("Using mocked data")
		return datasource.GetMockedData()
	} else {
		apiUrl := "https://sodertaljeglue.eu-gb.mybluemix.net/getwatertemp"
		return datasource.GetDataFromDataSourceFullJson(apiUrl)
	}
}
