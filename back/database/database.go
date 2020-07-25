package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
)

var conn *pgx.Conn

type DataBaseSevice struct {
	dbUrl      string
	dbName     string
	dbUser     string
	dbPassword string
}

const (
	dbUrl      = "jdbc:postgresql://0.0.0.0:5432/central-db"
	dbName     = "central-db"
	dbUser     = "admim"
	dbPassword = "admim"
)

// Создание сервиса для работы с базой
func Init() *DataBaseSevice {
	log.Println("Start init db service")

	dbService := &DataBaseSevice{
		dbUrl,
		dbName,
		dbUser,
		dbPassword,
	}

	conn, err := dbService.openConnect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// проверим соединение
	if err := conn.Ping(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)

	}
	log.Println("Finish db service")
	return dbService
}

// Открытие соединения к бд
func (db *DataBaseSevice) openConnect() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), dbUrl)
}
