package main

import (
	"log"
	"net/http"
)

func main() {
	host := "localhost:8080"
	path := "/api/health"
	mux := http.NewServeMux()
	mux.Handle(path, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(404)
	}))
	if err := http.ListenAndServe(host, mux); err != nil {
		log.Fatal(err)
	}
}
