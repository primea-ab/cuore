package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"lethe.se/vito/cuore/pkg/model"
	"lethe.se/vito/cuore/pkg/serv"
	"log"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	cfg := serv.FetchConfig()

	port := cfg.Combined.Port
	db := make(map[string]map[string]model.Datapoint)

	fmt.Println("Starting listener handlers")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadFile("../frontend/index.html")
    	if err != nil {
        	log.Printf("%v", err)
    	}
		w.Write(body)
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fetchData(w, db)
		case http.MethodPost:
			addData(w, r, db)
		default:
			response(w, "Invalid HTTP Method", http.StatusMethodNotAllowed)
			return
		}
	})

	fmt.Println("Listening to port " + port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatalf("Failed to listen to port %v", port)
	}
}

func fetchData(w http.ResponseWriter, db map[string]map[string]model.Datapoint) {
	jsonResponse(w, db)
}

func addData(w http.ResponseWriter, r *http.Request, db map[string]map[string]model.Datapoint) {
	if err := r.ParseForm(); err != nil {
		_, err = fmt.Fprintf(w, "ParseForm() err: %v", err)
		if err != nil {
			fmt.Println("Failed to return error message")
		}
		return
	}

	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		response(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e model.Datapoint
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			response(w, "Bad Request. Wrong Type provided for field " + unmarshalErr.Field, http.StatusBadRequest)
		} else {
			response(w, "Bad Request " + err.Error(), http.StatusBadRequest)
		}
		return
	}

	e.Timestamp = time.Now().Format(time.UnixDate)

	if db[e.Identifier] == nil {
		db[e.Identifier] = make(map[string]model.Datapoint)
	}
	idb := db[e.Identifier]
	idb[e.Name] = e

	response(w, "Success", http.StatusOK)
	return
}

func response(w http.ResponseWriter, message string, httpStatusCode int) {
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
