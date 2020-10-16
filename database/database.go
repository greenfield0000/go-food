package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/greenfield0000/go-food/back/model"
	"github.com/jackc/pgx/v4"
)

type DataBaseSevice struct {
	DBConn *pgx.Conn
}

const (
	dbName     = "central-db"
	dbUser     = "admin"
	dbPassword = "admin"
)

// Создание сервиса для работы с базой
func Init() *DataBaseSevice {
	log.Println("Start init db service")

	dbService := &DataBaseSevice{}

	DBConn, err := dbService.openConnect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	dbService.DBConn = DBConn

	// проверим соединение
	if err := DBConn.Ping(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)

	}
	log.Println("Finish db service")
	return dbService
}

// Открытие соединения к бд
func (db *DataBaseSevice) openConnect() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@localhost/%s", dbUser, dbPassword, dbName))
}

func (db *DataBaseSevice) SelectAll(tableName string) []model.DishModel {
	if db.DBConn != nil {
		rows, err := db.DBConn.Query(context.Background(), fmt.Sprintf("select * from \"%s\"", tableName))
		if err != nil {
			log.Fatalln("SelectAll with error " + err.Error())
		}
		defer rows.Close()

		var dishs []model.DishModel

		for rows.Next() {
			var id int32
			var name string
			var sysName string

			err := rows.Scan(&id, &name, &sysName)
			if err != nil {
				log.Fatalln("Row scan with error " + err.Error())
			}
			dish := model.DishModel{
				ID: id,
				Name: name,
				Sysname: sysName,
			}
			dishs = append(dishs, dish)
		}

		return dishs
	}

	return []model.DishModel{}
}

func (db *DataBaseSevice) CloseConn() {
	if db.DBConn != nil {
		db.DBConn.Close(context.Background())
	}
}
