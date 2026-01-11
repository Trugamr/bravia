package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/trugamr/bravia/api"
	"github.com/trugamr/bravia/cmd/remote/config"
	"github.com/trugamr/bravia/cmd/remote/handlers"
)

func main() {
	// Load configuration
	cfg := config.New()
	if err := cfg.Load(); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %s\n", err)
		os.Exit(1)
	}

	// Parse base URL and create API client
	baseURL, err := url.Parse(cfg.BaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing base URL: %s\n", err)
		os.Exit(1)
	}

	client := api.NewClient(baseURL).WithAuthPSK(cfg.PSK)

	// Create handler with client
	h := handlers.NewHandler(client)

	// Set up HTTP routes
	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("/api/power/on", h.PowerOnHandler)
	mux.HandleFunc("/api/power/off", h.PowerOffHandler)
	mux.HandleFunc("/api/power/status", h.PowerStatusHandler)

	mux.HandleFunc("/api/volume/set", h.VolumeSetHandler)
	mux.HandleFunc("/api/volume/up", h.VolumeUpHandler)
	mux.HandleFunc("/api/volume/down", h.VolumeDownHandler)

	mux.HandleFunc("/api/apps", h.AppsListHandler)
	mux.HandleFunc("/api/apps/open", h.AppsOpenHandler)

	mux.HandleFunc("/api/inputs", h.InputsListHandler)
	mux.HandleFunc("/api/inputs/select", h.InputsSelectHandler)

	mux.HandleFunc("/api/ircc/send", h.IRCCSendHandler)

	mux.HandleFunc("/api/sse", h.SSEHandler)

	// Static file routes
	mux.Handle("/", http.FileServer(http.Dir("./cmd/remote/web")))

	// Add CORS middleware
	corsHandler := corsMiddleware(mux)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	fmt.Printf("Bravia TV Remote starting on http://localhost%s\n", addr)
	fmt.Printf("TV Base URL: %s\n", cfg.BaseURL)

	if err := http.ListenAndServe(addr, corsHandler); err != nil {
		log.Fatal(err)
	}
}

// corsMiddleware adds CORS headers to responses
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
