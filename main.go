package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "hello\n")
	if err != nil {
		log.Fatal("something went wrong")
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			_, err := fmt.Fprintf(w, "%v: %v", name, h)
			if err != nil {
				log.Fatal("something went wrong")
			}
		}
	}
}

func main() {
	port := 8842

	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", hello)
	mux.HandleFunc("GET /headers", headers)

	log.Printf("starting server on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		log.Fatal("failed to bind to server")
	}
}
