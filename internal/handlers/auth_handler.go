package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/KenPrz/channel-grid-backend/internal/services"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req AuthRequest

	json.NewDecoder(r.Body).Decode(&req)

	err := services.Register(
		req.Email,
		req.Password,
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "registered",
	})
}

func LoginHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req AuthRequest

	json.NewDecoder(r.Body).Decode(&req)

	token, err := services.Login(
		req.Email,
		req.Password,
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusUnauthorized,
		)

		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
