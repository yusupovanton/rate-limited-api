package example

import (
	"log/slog"
	"net/http"
)

//go:generate ../../../bin/mockery --name rateLimiter

type rateLimiter interface {
	IsAllowed(ip string) bool
}

type Handler struct {
	rl rateLimiter

	log *slog.Logger
}

func NewHandler(rl rateLimiter, log *slog.Logger) *Handler {
	return &Handler{rl: rl, log: log}
}

func (api *Handler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "only get method is allowed", http.StatusMethodNotAllowed)

			return
		}

		ip := r.RemoteAddr
		log := api.log.With("ip", ip)

		if !api.rl.IsAllowed(ip) {
			log.Error("request denied: rate limit exceeded")
			http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)

			return
		}

		log.Info("request accepted")
	}
}
