package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// TVState represents the current state of the TV
type TVState struct {
	PowerStatus string `json:"powerStatus"`
	Volume      int    `json:"volume"`
	Muted       bool   `json:"muted"`
	Timestamp   string `json:"timestamp"`
}

// SSEHandler streams TV state updates via Server-Sent Events
func (h *Handler) SSEHandler(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create a channel for flushing
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// Client disconnect detection
	ctx := r.Context()
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	// Send initial connection message
	fmt.Fprintf(w, "data: {\"type\":\"connected\"}\n\n")
	flusher.Flush()

	for {
		select {
		case <-ctx.Done():
			// Client disconnected
			return
		case <-ticker.C:
			// Poll TV state
			state := h.getTVState()
			if state != nil {
				data, err := json.Marshal(state)
				if err == nil {
					fmt.Fprintf(w, "data: %s\n\n", data)
					flusher.Flush()
				}
			}
		}
	}
}

// getTVState polls the TV for current state
func (h *Handler) getTVState() *TVState {
	state := &TVState{
		PowerStatus: "unknown",
		Volume:      0,
		Muted:       false,
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	// Get power status
	powerResult, _, err := h.Client.System.GetPowerStatus()
	if err == nil && powerResult.Result != nil && len(*powerResult.Result) > 0 {
		state.PowerStatus = (*powerResult.Result)[0].Status
	}

	// Get volume information
	volumeResult, _, err := h.Client.Audio.GetVolumeInformation()
	if err == nil && volumeResult.Result != nil && len(*volumeResult.Result) > 0 {
		volumes := (*volumeResult.Result)[0]
		for _, v := range volumes {
			if v.Target == "speaker" {
				state.Volume = v.Volume
				state.Muted = v.Mute
				break
			}
		}
	}

	return state
}
