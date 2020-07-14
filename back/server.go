package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type LoginModel struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

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

	var loginModel LoginModel
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
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/registry", registryHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
