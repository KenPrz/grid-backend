package main

import (
	"fmt"
	"net/http"

	"github.com/KenPrz/channel-grid-backend/internal/database"
	"github.com/KenPrz/channel-grid-backend/internal/routes"
	"github.com/go-chi/chi/v5"
)

func main() {
	database.Connect()

	r := chi.NewRouter()

	routes.RegisterRoutes(r)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		panic(err)
	}
}
