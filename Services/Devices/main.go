package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	//db, _ := dbConnect()
	// checkError(err)
	// fmt.Println("Database successfully connected")
	// defer db.Close()

	a:=App{}
	a.Init(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
	
	a.Run(":8080")




	App.SetupRouter()
	r := addRouter(dataServer.Router)

	fmt.Println("Server listening on port 8080")
	e := http.ListenAndServe(":8080", r)
	checkError(e)
}
