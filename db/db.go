package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB

func ConnectDB(s string) (*sql.DB, error) {
	db, err := sql.Open("mysql", s)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func Init() {
	dbInfo := fmt.Sprint(os.Getenv("DB_USER"), ":", os.Getenv("DB_PASS"), "@tcp(127.0.0.1:3306)/",
		os.Getenv("DB_NAME"))
	fmt.Println(dbInfo)

	var err error

	db, err = ConnectDB(dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

}

func GetDb() *sql.DB {
	return db
}
