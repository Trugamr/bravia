package handlers

import (
	"encoding/json"
	"net/http"
)

// InputSelectRequest represents the request body for selecting an input
type InputSelectRequest struct {
	URI string `json:"uri"`
}

// InputsListHandler lists all available external inputs
func (h *Handler) InputsListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	result, _, err := h.Client.AVContent.GetCurrentExternalInputsStatus()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if result.Result == nil || len(*result.Result) == 0 {
		respondWithError(w, http.StatusInternalServerError, "Invalid response from TV")
		return
	}

	inputs := (*result.Result)[0]

	respondWithSuccess(w, inputs)
}

// InputsSelectHandler selects an input by URI
func (h *Handler) InputsSelectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req InputSelectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.URI == "" {
		respondWithError(w, http.StatusBadRequest, "uri is required")
		return
	}

	_, _, err := h.Client.AVContent.SetPlayContent(req.URI)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(w, map[string]string{"uri": req.URI})
}
