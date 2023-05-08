package datasource

import (
	"fmt"
	"net/http"
)

func LinkHandler(w http.ResponseWriter, r *http.Request, data []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func StartServer(endpoint string, data []byte) {
	fmt.Println("Check the response from the web server with http://localhost:8080" + endpoint)
	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		LinkHandler(w, r, data)
	})
	http.ListenAndServe(":8080", nil)
}
