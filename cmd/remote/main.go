package main

import (
	"fmt"
	"net/http"
)

// TODO: Implement server to serve static page and API to control the TV

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", rootHandler)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusNotImplemented)

	fmt.Fprintf(w, "not implemented")
}
