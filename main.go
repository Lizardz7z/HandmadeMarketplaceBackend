package main

import (
	"fmt"
	"log"
	"net/http"

	"diploma/config"
	"diploma/handlers"
)

func main() {
	config.LoadConfig()

	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/protected", handlers.ProtectedHandler)

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
