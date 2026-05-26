package routes

import (
	"net/http"

	"github.com/KenPrz/channel-grid-backend/internal/handlers"
	"github.com/KenPrz/channel-grid-backend/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {

	// Exteral Routes
	r.Get("/ping", handlers.PingHandler)
	r.Post("/register", handlers.RegisterHandler)
	r.Post("/login", handlers.LoginHandler)

	r.Group(func(protected chi.Router) {
		protected.Use(middleware.AuthMiddleware)

		protected.Get("/profile", func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			w.Write([]byte("protected route"))
		})
	})
}
