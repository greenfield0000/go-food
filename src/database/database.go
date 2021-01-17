package database

import (
	"log"

	"github.com/greenfield0000/go-food/microservices/go-food-auth/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbName     = "central-db"
	dbUser     = "admin"
	dbPassword = "admin"
	dbPort     = "5432"
	sslMode    = "disable"
	dbTimeZone = "Europe/Moscow"
)

// OpenDB - Открытие бд
func OpenDB() (db *gorm.DB, err error) {
	dsn := "host=127.0.0.1 user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + sslMode + " TimeZone=" + dbTimeZone
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// startAutoMigrate - Function with start automigrate by struct
func StartAutoMigrate() {
	db, err := OpenDB()
	if err != nil {
		log.Fatal("startAutoMigrate is error = ", err.Error())
		return
	}
	// Запуск миграции
	err = db.AutoMigrate(&model.AccountModel{})
	if err != nil {
		log.Fatal("startAutoMigrate is error = ", err.Error())
		return
	}
}
