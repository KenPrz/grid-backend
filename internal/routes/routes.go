package routes

import (
	"github.com/KenPrz/channel-grid-backend/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {

	r.Get("/ping", handlers.PingHandler)
}
