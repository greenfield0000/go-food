package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/greenfield0000/go-food/microservices/go-food-auth/model"
)

// CreatePerson create user with request data
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error read body"))
		return
	}

	var user model.User
	if err = json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Unmarshal with error " + err.Error()))
		return
	}

}
