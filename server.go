package main

import (
	"github.com/greenfield0000/go-food/back/database"
	"github.com/greenfield0000/go-food/back/handlers/auth"
	"log"
	http "net/http"
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
	http.HandleFunc("/", middleware(rootHandler))
	http.HandleFunc("/registry", middleware(auth.RegistryHandler))
	http.HandleFunc("/login", middleware(auth.LoginHandler))
	http.HandleFunc("/logout", middleware(auth.LogoutHandler))

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Run middleware start")
		next.ServeHTTP(w, r)
		log.Println("Run middleware finish")
	})
}
