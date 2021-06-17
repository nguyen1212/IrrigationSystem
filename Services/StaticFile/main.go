package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("../../Frontend/dist"))
	http.Handle("/", fs)
	fmt.Println("Server listening on port 10000")
	err	:=	http.ListenAndServe(":10000", nil)
	if err	!= nil{
		log.Fatal(err)
	}
}