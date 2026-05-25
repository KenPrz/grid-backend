package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/KenPrz/channel-grid-backend/internal/services"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	message := services.Ping()

	respose := map[string]string{
		"message": message,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(respose)
}
