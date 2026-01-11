package handlers

import (
	"encoding/json"
	"net/http"
)

// IRCCSendRequest represents the request body for sending an IRCC command
type IRCCSendRequest struct {
	Command string `json:"command"`
}

// IRCCSendHandler sends an IRCC remote control command
func (h *Handler) IRCCSendHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req IRCCSendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Command == "" {
		respondWithError(w, http.StatusBadRequest, "command is required")
		return
	}

	_, err := h.Client.IRCC.SendIRCCCommand(req.Command)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithSuccess(w, map[string]string{"command": req.Command})
}
