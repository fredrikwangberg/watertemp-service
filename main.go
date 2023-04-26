package main

import (
	"fmt"
	"github.com/fredrikwangberg/watertemp/datasource"
)

func main() {
	data, err := getData(false)
	if err == nil {
		coldSite := datasource.GetLocationWithLowestTemperature(data)
		fmt.Println("Location with lowest water temperature:")
		fmt.Printf(datasource.WaterTemperatureSchemaToString(coldSite))
	} else {
		fmt.Println("No data, err?")
		fmt.Println(err)
	}
}

func getData(useMocked bool) ([]datasource.WaterTemperatureSchema, error) {
	if useMocked {
		fmt.Println("Using mocked data")
		return datasource.GetMockedData()
	} else {
		apiUrl := "https://sodertaljeglue.eu-gb.mybluemix.net/getwatertemp"
		return datasource.GetDataFromDataSourceFullJson(apiUrl)
	}
}
