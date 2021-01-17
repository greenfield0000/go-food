package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/greenfield0000/go-food/microservices/go-food-auth/model"
	"github.com/greenfield0000/go-food/microservices/go-food-auth/repository"
	"github.com/greenfield0000/go-secure-microservice"
)

var acr repository.AccountRepository

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Read body error %s", err)
		return
	}

	var loginRequest model.LoginRequest
	err = json.Unmarshal(body, &loginRequest)
	if err != nil {
		log.Println("Error unmarshal login model")
		return
	}
	account, err := acr.Find(loginRequest)
	if err != nil {
		w.Write([]byte("Error" + err.Error()))
		return
	}
	if account == nil {
		w.Write([]byte("Account not found"))
		return
	}
	tokenDetail, err := secure.CreateToken(uint64(account.ID))

	baseResp := model.BaseResponse{
		Status: http.StatusText(http.StatusOK),
		Result: model.Account{},
	}

	resp, err := json.Marshal(baseResp)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error ", err.Error())))
		return
	}

	acTokenCook := &http.Cookie{
		Name:     "token",
		Value:    tokenDetail.AccessToken,
		HttpOnly: false,
	}
	http.SetCookie(w, acTokenCook)
	rfTokenCook := &http.Cookie{
		Name:     "rtoken",
		Value:    tokenDetail.RefreshToken,
		HttpOnly: false,
	}

	http.SetCookie(w, rfTokenCook)
	w.Write(resp)
}

func RegistryHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error read body"))
		return
	}

	var loginRequest model.LoginRequest
	if err = json.Unmarshal(body, &loginRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Unmarshal with error " + err.Error()))
		return
	}

	account, err := acr.Find(loginRequest)
	if err != nil {
		w.Write([]byte("Error find account"))
		return
	}
	// если такой акк зарегистрирован
	if account != nil {
		w.Write([]byte("Name '" + loginRequest.Login + "' already exist "))
		return
	}

	newAccount := model.AccountModel{
		Login:    loginRequest.Login,
		Password: loginRequest.Password,
	}

	err = acr.Create(&newAccount)
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
