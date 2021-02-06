package main

import (
	"log"
	"net/http"
	"os"

	"github.com/greenfield0000/go-food/microservices/go-food-auth/database"
	"github.com/greenfield0000/go-food/microservices/go-food-auth/handlers"
)

// started server function
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server started"))
}

func init() {
	database.StartAutoMigrate()
}

func main() {
	// auth
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/auth/registry", middleWare(handlers.RegistryHandler))
	http.HandleFunc("/auth/login", middleWare(handlers.LoginHandler))
	http.HandleFunc("/auth/logout", middleWare(handlers.LogoutHandler))

	// Создание пользователя
	http.HandleFunc("/auth/createNewPerson", middleWare(handlers.CreatePerson))

	log.Fatalln(http.ListenAndServe(getServicePort(), nil))
}

func middleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isOptions := isEnabledCors(w, r); isOptions {
			return
		}
		next.ServeHTTP(w, r)
	}
}

func isEnabledCors(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, origin, accept, content-type, access-control-allow-methods, authorization")

	return r.Method == "OPTIONS"
}

// getServicePort get port with service listen
func getServicePort() string {
	servicePort := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		servicePort = ":" + port
	}
	return servicePort
}
