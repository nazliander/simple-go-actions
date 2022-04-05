package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nazliander/simple-go-actions/predictor"
)

type Response struct {
	Result float64 `json:"result"`
}

func forecast(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("n")
	parsedStrs, err := predictor.SplitParser(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := Response{
		Result: predictor.MovingAverage(parsedStrs),
	}
	w.Header().Set("Content-Type", "application/json")
	encode_err := json.NewEncoder(w).Encode(resp)
	if encode_err != nil {
		http.Error(w, encode_err.Error(), http.StatusBadRequest)
		return
	}
}

func handleRequests() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", forecast)
	log.Fatal(http.ListenAndServe(":8080", router))

}

func main() {
	handleRequests()
}
