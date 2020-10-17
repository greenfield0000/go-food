package auth

import (
	"encoding/json"
	"fmt"
	"github.com/greenfield0000/go-food/back/model"
	"github.com/greenfield0000/go-food/back/repository"
	"io/ioutil"
	"log"
	"net/http"
)

var acr repository.AccountRepository

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Read body error %s", err)
		return
	}

	var accountModel model.AccountModel
	err = json.Unmarshal(body, &accountModel)
	if err != nil {
		log.Println("Error unmarshal login model")
		return
	}

	account, err := acr.FindAccount(accountModel)
	if err != nil {
		w.Write([]byte("Error" + err.Error()))
		return
	}

	if account == nil {
		w.Write([]byte("Account not found"))
		return
	}

	marshal, err := json.Marshal(account)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error ", err.Error())))
		return
	}
	w.Write(marshal)
}

func RegistryHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error read body"))
		return
	}

	var accountModel model.AccountModel
	if err = json.Unmarshal(body, &accountModel); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Unmarshal with error " + err.Error()))
		return
	}

	account, err := acr.FindAccount(accountModel)
	if err != nil {
		w.Write([]byte("Error find account"))
		return
	}
	// если такой акк зарегистрирован
	if account != nil {
		w.Write([]byte("Name '" + accountModel.Login + "' already exist "))
		return
	}

	err = acr.Create(&accountModel)
	if err != nil {
		w.Write([]byte("Not created"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Created"))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout action"))
}
