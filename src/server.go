package main

import (
	"github.com/greenfield0000/go-food/microservices/go-food-auth/database"
	"github.com/greenfield0000/go-food/microservices/go-food-auth/handlers"
	"log"
	"net/http"
	"os"
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
	http.HandleFunc("/registry", handlers.RegistryHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)

	log.Fatalln(http.ListenAndServe(getServicePort(), nil))
}

// getServicePort get port with service listen
func getServicePort() string {
	servicePort := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		servicePort = ":" + port
	}
	return servicePort
}
