package main

import (
	"fmt"
	"log"
	"net/http"
)

// Try curl http://127.0.0.1:8080/hello in terminal, you will see the following output:
// Request Headers:
// Header["Accept"] = ["*/*"]
// Header["User-Agent"] = ["curl/8.5.0"]

// Try curl http://127.0.0.1:8080/, you will see the following output:
// URL.Path = "/"

func main() {
	// Register route handlers
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/hello", handleHello)

	// Start the server
	log.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// handleHome handles requests to the root path
func handleHome(w http.ResponseWriter, req *http.Request) {
	log.Printf("Received request at %s from %s\n", req.URL.Path, req.RemoteAddr)
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handleHello handles requests to /hello path and displays all request headers
func handleHello(w http.ResponseWriter, req *http.Request) {
	log.Printf("Received request at %s from %s\n", req.URL.Path, req.RemoteAddr)
	fmt.Fprintf(w, "Request Headers:\n")
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
