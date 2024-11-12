package config

import "os"

// JWTKey — секретный ключ для подписи JWT
var JWTKey []byte

// LoadConfig — функция для загрузки конфигурации
func LoadConfig() {
	JWTKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(JWTKey) == 0 {
		JWTKey = []byte("my_secret_key")
	}
}
