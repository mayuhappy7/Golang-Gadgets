package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Log incoming request
	log.Printf("Received request at %s from %s\n", req.URL.Path, req.RemoteAddr)

	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		fmt.Fprintf(w, "Request Headers:\n")
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %q\n", req.URL)
	}
}

func main() {
	engine := &Engine{}
	// Log server startup
	log.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", engine))
}
