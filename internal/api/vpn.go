package api

import (
	"encoding/json"
	"net/http"
)

type MessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func StartVPN(w http.ResponseWriter, r *http.Request) {
	respondOK(w, "VPN start requested")
}

func StopVPN(w http.ResponseWriter, r *http.Request) {
	respondOK(w, "VPN stop requested")
}

func RestartVPN(w http.ResponseWriter, r *http.Request) {
	respondOK(w, "VPN restart requested")
}

func respondOK(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(MessageResponse{
		Success: true,
		Message: message,
	})
}
