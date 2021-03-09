package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type datapoint struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func main() {
	port := "7676"
	db := make(map[string]datapoint)

	fmt.Println("Starting listener handlers")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fetchData(w, db)
		case http.MethodPost:
			addData(w, r, db)
		default:
			errorResponse(w, "Invalid HTTP Method", http.StatusMethodNotAllowed)
			return
		}
	})

	fmt.Println("Listening to port " + port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatalf("Failed to listen to port " + port)
	}
}

func fetchData(w http.ResponseWriter, db map[string]datapoint) {
	jsonResponse(w, db)
}

// Should I clean up http.HandleFunc call?
func addData(w http.ResponseWriter, r *http.Request, db map[string]datapoint) {
	if err := r.ParseForm(); err != nil {
		_, err = fmt.Fprintf(w, "ParseForm() err: %v", err)
		if err != nil {
			fmt.Println("Failed to return error message")
		}
		return
	}

	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e datapoint
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field " + unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request " + err.Error(), http.StatusBadRequest)
		}
		return
	}
	db[e.Name] = e
	errorResponse(w, "Success", http.StatusOK)
	return
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResponse(w, resp)
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(data)
	_, err := w.Write(jsonResp)
	if err != nil {
		fmt.Println("Error writing json jsonResponse")
	}
}
