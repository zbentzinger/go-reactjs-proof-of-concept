package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	indexRouter := http.HandlerFunc(indexHandler)
	exampleRouter := http.HandlerFunc(exampleHandler)

	http.Handle("/", indexRouter)
	http.Handle("/example", exampleRouter)

	log.Printf("Starting the HTTP server on port 8080\n")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.UserAgent(), r.Method, r.RequestURI, r.Body)
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.UserAgent(), r.Method, r.RequestURI, r.Body)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["message"] = "200 OK"
	resp["body"] = "Test Example"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s\n", err)
	}
	w.Write(jsonResp)
	return
}
