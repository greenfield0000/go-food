package model

import (
	"time"

	"gorm.io/gorm"
)

type AccountModel struct {
	gorm.Model
	Login    string `db:"login"`
	Password string `db:"password"`
}

type BaseResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Account Account `json:"account"`
}

type Account struct {
	Uuid     string `json:"uuid"`
	NickName string `json:"nickname"`
	User     User   `json:"user"`
}

type addressItem struct {
	Cadnum      string `json:"cadnum"`
	ContentType string `json:"contentType"`
	Guid        string `json:"guid"`
	Id          string `json:"id"`
	Ifnsfl      string `json:"ifnsfl"`
	Ifnsul      string `json:"ifnsul"`
	Name        string `json:"name"`
	Okato       string `json:"okato"`
	Oktmo       string `json:"oktmo"`
	ParentGuid  string `json:"parentGuid"`
	TypeAddress string `json:"type"`
	TypeShort   string `json:"typeShort"`
	Zip         int32  `json:"zip"`
}

type address struct {
	Region, City, Street, Building addressItem
	Appartment                     string
}

type User struct {
	Name        string    `json:"name"`
	SurName     string    `json:"surName"`
	LastName    string    `json:"lastName"`
	BirthDay    time.Time `json:"birthDay"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	AddressList []address `json:"addressList"`
}
