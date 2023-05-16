package datasource

import (
	"net/http"
)

func ColdestTemperatureLocationHandler(w http.ResponseWriter, r *http.Request) {
	data, err := getData(false)
	if err == nil {
		coldest := GetLocationWithLowestTemperature(data)
		coldestJson, err := GetTemperatureLocationJson(coldest)

		if err == nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			_, err := w.Write(coldestJson)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
