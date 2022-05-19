package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	host := "localhost:8080"
	mux := http.NewServeMux()
	mux.Handle("/api/health", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(400)
	}))
	mux.Handle("/api/rawjson", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		object := SomeObject{Data: Data{A: "test", B: 2}}
		data, err := json.Marshal(object)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write(data)

	}))
	if err := http.ListenAndServe(host, mux); err != nil {
		log.Fatal(err)
	}
}

type SomeObject struct {
	Data Data
}

type Data struct {
	A string
	B int
}
