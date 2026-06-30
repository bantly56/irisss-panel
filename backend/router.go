package main

import "net/http"

func NewRouter() *http.ServeMux {

	mux := http.NewServeMux()

	// API
	mux.HandleFunc("/api/status", StatusHandler)
	mux.HandleFunc("/api/start", StartHandler)
	mux.HandleFunc("/api/stop", StopHandler)
	mux.HandleFunc("/api/restart", RestartHandler)

	mux.HandleFunc("/api/logs", LogsHandler)

	mux.HandleFunc("/api/import", ImportHandler)

	mux.HandleFunc("/api/backup", BackupHandler)
	mux.HandleFunc("/api/restore", RestoreHandler)

	mux.HandleFunc("/api/routing/enable", EnableRoutingHandler)
	mux.HandleFunc("/api/routing/disable", DisableRoutingHandler)

	// Frontend
	mux.Handle("/", http.FileServer(http.Dir("./frontend")))

	return mux
}
