package models

import (
	"github.com/golang-jwt/jwt/v4"
)

// User — структура данных пользователя
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims — структура данных для JWT
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
