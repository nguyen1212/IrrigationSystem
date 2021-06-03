package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"irrigation.com/Devices/Data/database"
)

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func dbConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/device_database")
	if err != nil {
		return db, err
	}
	return db, nil
}

func main() {
	// http.HandleFunc("/devices/data", handleDeviceData)

	db, err := dbConnect()
	checkError(err)
	fmt.Println("Database successfully connected")
	defer db.Close()

	dbServer := &database.DB{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: db,
	}

	dbServer.SetupRouter()

	fmt.Println("Server listening on port 8081-test")
	er := http.ListenAndServe(":8081", dbServer.Router)
	checkError(er)
}
