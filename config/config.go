package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ConnectDB() *sql.DB {

	db, err := sql.Open("mysql", "admin:admin@tcp(godockerDB)/cake_store") //TODO: uncomment if want run with docker
	//db, err := sql.Open("mysql", "admin:admin@/cake_store") //TODO: run with "go run main.go"

	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}

	// make sure connection is available
	err = db.Ping()
	if err != nil {
		log.Printf("Test Connection Failed", err.Error())
	}
	return db
}
