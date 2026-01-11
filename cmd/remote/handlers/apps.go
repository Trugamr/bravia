package handlers

import (
	"encoding/json"
	"net/http"
)

// AppOpenRequest represents the request body for opening an app
type AppOpenRequest struct {
	URI string `json:"uri"`
}

// AppsListHandler lists all available apps
func (h *Handler) AppsListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	result, _, err := h.Client.AppControl.GetApplicationList()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if result.Result == nil || len(*result.Result) == 0 {
		respondWithError(w, http.StatusInternalServerError, "Invalid response from TV")
		return
	}

	apps := (*result.Result)[0]

	respondWithSuccess(w, apps)
}

// AppsOpenHandler opens an app by URI
func (h *Handler) AppsOpenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req AppOpenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.URI == "" {
		respondWithError(w, http.StatusBadRequest, "uri is required")
		return
	}

	_, _, err := h.Client.AppControl.SetActiveApp(req.URI, nil)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(w, map[string]string{"uri": req.URI})
}
