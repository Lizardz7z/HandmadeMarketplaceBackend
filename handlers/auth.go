package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"diploma/config"
	"diploma/models"

	"github.com/golang-jwt/jwt/v4"
)

// Временная БД
var users = map[string]string{
	"testuser": "password123",
}

// LoginHandler — хэндлер для обработки логина
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[user.Username]
	if !ok || expectedPassword != user.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JWTKey)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	response := map[string]string{"token": tokenString}
	json.NewEncoder(w).Encode(response)
}
