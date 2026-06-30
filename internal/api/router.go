package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {

	r := chi.NewRouter()

	r.Get("/api/status", Status)
	r.Post("/api/vless/import", ImportVLESS)

	r.Post("/api/vpn/start", StartVPN)
	r.Post("/api/vpn/stop", StopVPN)
	r.Post("/api/vpn/restart", RestartVPN)

	r.Post("/api/routing/enable", EnableRouting)
	r.Post("/api/routing/disable", DisableRouting)

	r.Get("/api/logs", Logs)

	return r
}
