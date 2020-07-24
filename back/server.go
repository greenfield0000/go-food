package main

import (
	"encoding/json"
	"fmt"
	model "github.com/greenfield0000/go-food/back/model"
	"io/ioutil"
	"log"
	"net/http"
)

// started server function
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server started"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Read body error %s", err)
		return
	}

	var loginModel model.LoginModel
	err = json.Unmarshal(body, &loginModel)
	if err != nil {
		log.Println("Error unmarshal login model")
		return
	}

	w.Write([]byte(fmt.Sprintf("Was readed login = %s and password %s", loginModel.UserName, loginModel.Password)))
}

func registryHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("registry action"))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout action"))
}

func main() {
	http.HandleFunc("/", midleware(rootHandler))
	http.HandleFunc("/registry", midleware(registryHandler))
	http.HandleFunc("/login", midleware(loginHandler))
	http.HandleFunc("/logout", midleware(logoutHandler))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func midleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Run midleware start")
		next.ServeHTTP(w, r)
		log.Println("Run midleware finish")

	})
}
