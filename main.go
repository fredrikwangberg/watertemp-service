package main

import (
	"fmt"
	"net/http"

	"github.com/fredrikwangberg/watertemp/datasource"
)

func main() {

	fmt.Println("Check the response from the web server with http://localhost:8080/coldest")
	http.HandleFunc("/coldest", func(w http.ResponseWriter, r *http.Request) {
		datasource.ColdestTemperatureLocationHandler(w, r, datasource.GetTemperatureLocationFromSource)
	})
	http.ListenAndServe(":80", nil)

}
