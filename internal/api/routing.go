package api

import "net/http"

func EnableRouting(w http.ResponseWriter, r *http.Request) {
	respondOK(w, "Routing enabled")
}

func DisableRouting(w http.ResponseWriter, r *http.Request) {
	respondOK(w, "Routing disabled")
}
