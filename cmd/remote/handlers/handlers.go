package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/trugamr/bravia/api"
)

// Handler holds the Bravia API client for all handler functions
type Handler struct {
	Client *api.Client
}

// NewHandler creates a new handler with the given Bravia API client
func NewHandler(client *api.Client) *Handler {
	return &Handler{
		Client: client,
	}
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse represents a generic success response
type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

// respondWithError sends an error response
func respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

// respondWithJSON sends a JSON response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

// respondWithSuccess sends a success response
func respondWithSuccess(w http.ResponseWriter, data interface{}) {
	respondWithJSON(w, http.StatusOK, SuccessResponse{
		Success: true,
		Data:    data,
	})
}
