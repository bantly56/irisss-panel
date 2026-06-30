package api

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

type StatusResponse struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Status    string `json:"status"`
	Time      string `json:"time"`
	GoVersion string `json:"go_version"`
}

func Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := StatusResponse{
		Name:      "Irisss Panel",
		Version:   "0.1.0",
		Status:    "running",
		Time:      time.Now().Format(time.RFC3339),
		GoVersion: runtime.Version(),
	}

	_ = json.NewEncoder(w).Encode(resp)
}
