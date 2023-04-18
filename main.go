package main

import (
	"fmt"
	"io.github.fredrikwangberg/watertemp"
)


func main() {
	data, err := getData(false)

	if err == nil {
		coldSite := watertemp.GetLocationWithLowestTemperature(data)
	    fmt.Println("Location with lowest water temperature:")
		fmt.Printf(watertemp.WaterTemperatureSchemaToString(coldSite))
	} else {
		fmt.Println("No data, err?")
		fmt.Println(err)
	}
}

func getData(useMocked bool) ([]watertemp.WaterTemperatureSchema, error) {
	if useMocked {
	    fmt.Println("Using mocked data")
		return watertemp.GetMockedData()
	} else {
		apiUrl := "https://sodertaljeglue.eu-gb.mybluemix.net/getwatertemp"
		return watertemp.GetDataFromDataSourceFullJson(apiUrl)
	}
}
