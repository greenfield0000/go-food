package model

import "gorm.io/gorm"

type AccountModel struct {
	gorm.Model
	Login    string `json:"login"`
	Password string `json:"password"`
}
