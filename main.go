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

	/* TO DO: Updated main() in order to test the function that I built. */
	if err == nil {
		jsonColdSite, err := datasource.GetJsonColdSite(data)

		if err == nil {
			fmt.Println("Here is the string of the json-object for the location with the lowest temperature:")
			fmt.Println(string(jsonColdSite))
		} else {
			fmt.Println("Error getting JSON object: ", err)
		}

	} else {
		fmt.Println("No data, err?")
		fmt.Println(err)
	}
	/* END TO DO*/
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
