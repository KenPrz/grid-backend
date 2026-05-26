package services

import (
	"database/sql"
	"errors"
	"time"

	"github.com/KenPrz/channel-grid-backend/internal/database"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("secret-key")

func Register(email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	_, err = database.DB.Exec(
		"INSERT INTO users(email, password) VALUES(?,?)",
		email,
		string(hashedPassword),
	)

	return err
}

func Login(email string, password string) (string, error) {
	var hashedPassword string

	err := database.DB.QueryRow(
		"SELECT password FROM users WHERE email = ?",
		email,
	).Scan(&hashedPassword)

	if err == sql.ErrNoRows {
		return "", errors.New("Invalid Credentials")
	}

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("Invalid Credentials")
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(24 * time.Hour).Unix(),
		},
	)

	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, err
}
