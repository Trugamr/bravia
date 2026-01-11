package handlers

import (
	"net/http"
)

// PowerOnHandler turns the TV on
func (h *Handler) PowerOnHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	_, _, err := h.Client.System.SetPowerStatus(true)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(w, map[string]string{"status": "on"})
}

// PowerOffHandler turns the TV off
func (h *Handler) PowerOffHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	_, _, err := h.Client.System.SetPowerStatus(false)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(w, map[string]string{"status": "off"})
}

// PowerStatusHandler gets the current power status
func (h *Handler) PowerStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	result, _, err := h.Client.System.GetPowerStatus()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if result.Result == nil || len(*result.Result) == 0 {
		respondWithError(w, http.StatusInternalServerError, "Invalid response from TV")
		return
	}

	status := (*result.Result)[0].Status

	respondWithSuccess(w, map[string]string{"status": status})
}
