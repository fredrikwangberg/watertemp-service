package main

import (
	"fmt"

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

			datasource.StartServer("/coldest", coldestJson)

		} else {
			fmt.Println("Error getting JSON object: ", err)
		}

	} else {
		fmt.Println("No data, err?")
		fmt.Println(err)
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
