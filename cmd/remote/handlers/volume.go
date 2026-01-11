package handlers

import (
	"encoding/json"
	"net/http"
)

// VolumeSetRequest represents the request body for setting volume
type VolumeSetRequest struct {
	Volume string `json:"volume"`
	Target string `json:"target,omitempty"`
}

// VolumeSetHandler sets the volume to a specific level
func (h *Handler) VolumeSetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req VolumeSetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Volume == "" {
		respondWithError(w, http.StatusBadRequest, "volume is required")
		return
	}

	// Default target to speaker if not provided
	target := req.Target
	if target == "" {
		target = "speaker"
	}

	result, _, err := h.Client.Audio.SetAudioVolume(req.Volume, target)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if result.Result == nil || len(*result.Result) == 0 {
		respondWithError(w, http.StatusInternalServerError, "Invalid response from TV")
		return
	}

	newVolume := (*result.Result)[0]

	respondWithSuccess(w, map[string]int{"volume": newVolume})
}

// VolumeUpHandler increases the volume
func (h *Handler) VolumeUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	result, _, err := h.Client.Audio.SetAudioVolume("+1", "speaker")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if result.Result == nil || len(*result.Result) == 0 {
		respondWithError(w, http.StatusInternalServerError, "Invalid response from TV")
		return
	}

	newVolume := (*result.Result)[0]

	respondWithSuccess(w, map[string]int{"volume": newVolume})
}

// VolumeDownHandler decreases the volume
func (h *Handler) VolumeDownHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	result, _, err := h.Client.Audio.SetAudioVolume("-1", "speaker")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if result.Result == nil || len(*result.Result) == 0 {
		respondWithError(w, http.StatusInternalServerError, "Invalid response from TV")
		return
	}

	newVolume := (*result.Result)[0]

	respondWithSuccess(w, map[string]int{"volume": newVolume})
}
