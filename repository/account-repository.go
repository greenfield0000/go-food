package repository

import (
	"errors"
	"github.com/greenfield0000/go-food/back/database"
	"github.com/greenfield0000/go-food/back/model"
	"github.com/greenfield0000/go-food/back/secure"
)

type AccountRepository struct{}

// FindAccount - find account by login and password params
func (ar *AccountRepository) FindAccount(accountModel model.AccountModel) (*model.AccountModel, error) {
	db, err := database.OpenDB()
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	var account model.AccountModel

	if err := db.Where("login = ?", accountModel.Login).Find(&account).Error; err != nil {
		return nil, errors.New("Get account is error")
	}

	if account.Login == accountModel.Login && secure.ComparePassword(account.Password, accountModel.Password) {
		return &account, nil
	}

	return nil, nil
}

// Create - function with create account
func (ar *AccountRepository) Create(accountModel *model.AccountModel) error {
	if accountModel == nil {
		return errors.New("Account is empty")
	}
	db, err := database.OpenDB()
	if err != nil {
		return errors.New("Error open database")
	}
	hash, err := secure.CreateHash(accountModel.Password)
	accountModel.Password = hash
	return db.Create(&accountModel).Error
}